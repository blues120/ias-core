package biz

import (
	"archive/zip"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/blues120/ias-core/pkg/pagination"
)

var (
	ErrFileUploadingInitFailed        = errors.New("文件初始化上传失败")
	ErrDuplicateFiles                 = errors.New("多个重复文件")
	ErrUnknownFileStatus              = errors.New("文件上传状态未知")
	ErrFileIsNotUploading             = errors.New("文件不在上传")
	ErrFileNameExists                 = errors.New("文件名重复")
	ErrFileNameChangesDuringUploading = errors.New("文件上传中不能改名")
)

type AlgoManagementUsecase struct {
	algoUc       *AlgoUsecase
	ossUc        *OssUsecase
	fileUploadUc *FileUploadUsecase
	log          *log.Helper
}

func NewAlgoManagementUsecase(algoUc *AlgoUsecase, ossUc *OssUsecase, fileUploadUc *FileUploadUsecase, logger log.Logger) *AlgoManagementUsecase {
	return &AlgoManagementUsecase{algoUc: algoUc, ossUc: ossUc, fileUploadUc: fileUploadUc, log: log.NewHelper(logger)}
}

func (uc *AlgoManagementUsecase) createNewUpload(ctx context.Context, provider, fileName, md5 string, totalBytes uint64) error {
	key := uuid.New().String()
	uploadId, err := uc.ossUc.CreateMultipartUpload(ctx, key)
	if err != nil {
		return err
	}
	_, err = uc.fileUploadUc.Create(ctx, &FileUpload{
		Provider:   provider,
		FileName:   fileName,
		Md5:        md5,
		TotalBytes: totalBytes,
		Key:        key,
		UploadID:   uploadId,
		Status:     FileUploadStatusUploading,
	})
	return err
}

func (uc *AlgoManagementUsecase) queryUploads(ctx context.Context, fileName string, upload *FileUpload) (string, []int64, error) {
	switch upload.Status {
	case FileUploadStatusSuccess: // 已经上传完成
		filePath, err := uc.ossUc.GenerateTemporaryUrl(ctx, upload.Key, 12*time.Hour)
		if err != nil {
			return "", nil, err
		}

		if fileName != upload.FileName { // 文件上传完成后允许改名
			newUpload := *upload
			newUpload.FileName = fileName
			err := uc.fileUploadUc.Update(ctx, upload.ID, &newUpload)
			if err != nil {
				return "", nil, err
			}
		}

		return filePath, nil, nil
	case FileUploadStatusUploading: // 部分上传

		if fileName != upload.FileName { // 上传过程中不允许改名
			return "", nil, ErrFileNameChangesDuringUploading
		}

		// 将字符串转换为 time.Time
		dateTimeFormat := "2006-01-02 15:04:05"
		createTime, err := time.Parse(dateTimeFormat, upload.CreateTime)
		if err != nil {
			return "", nil, ErrFileNameChangesDuringUploading
		}
		// 计算时间差
		duration := time.Since(createTime)

		// 判断时间差是否小于一天
		if duration.Hours() < 24 {
			parts, err := uc.ossUc.ListParts(ctx, upload.Key, upload.UploadID, 0)
			if err != nil {
				return "", nil, err
			}
			var partNums []int64
			for i := range parts {
				partNums = append(partNums, parts[i].PartNumber)
			}
			return "", partNums, nil
		} else {
			// 此时uploadId已经失效
			key := uuid.New().String()
			uploadId, err := uc.ossUc.CreateMultipartUpload(ctx, key)
			if err != nil {
				return "", nil, ErrFileNameChangesDuringUploading
			}
			newUpload := *upload
			newUpload.Key = key
			newUpload.UploadID = uploadId
			err = uc.fileUploadUc.Update(ctx, upload.ID, &newUpload)
			if err != nil {
				return "", nil, err
			}

			return "", []int64{}, nil
		}

	case FileUploadStatusFailed: // 失败时状态改成 uploading 重新上传
		newUpload := *upload
		newUpload.Status = FileUploadStatusUploading
		err := uc.fileUploadUc.Update(ctx, upload.ID, &newUpload)
		return "", nil, err
	default:
		return "", nil, ErrUnknownFileStatus
	}
}

// 检查文件名是否重复，单个供应商内不允许重名
func (uc *AlgoManagementUsecase) isFileNameExist(ctx context.Context, provider, fileName string) (bool, error) {
	_, count, err := uc.fileUploadUc.List(ctx, &FileUploadFilter{
		ProviderEq: provider,
		FileNameEq: fileName,
	})
	if err != nil {
		return false, err
	}
	return count >= 1, nil
}

// 检查文件上传情况，适用于单机和云边版
func (uc *AlgoManagementUsecase) CheckUploads(ctx context.Context, provider, fileName, md5 string, totalBytes uint64) (string, []int64, error) {
	// 检查是否已经上传
	uploads, count, err := uc.fileUploadUc.List(ctx, &FileUploadFilter{
		ProviderEq: provider,
		Md5Eq:      md5,
	})
	if err != nil {
		return "", nil, err
	}

	switch count {
	case 0: // 未上传，返回空
		if ok, err := uc.isFileNameExist(ctx, provider, fileName); ok && err == nil {
			return "", nil, ErrFileNameExists
		}
		err := uc.createNewUpload(ctx, provider, fileName, md5, totalBytes)
		if err != nil {
			return "", nil, err
		}
		return "", nil, nil
	case 1: // 已经上传，返回链接或分片
		filePath, parts, err := uc.queryUploads(ctx, fileName, uploads[0])
		if err != nil {
			return "", nil, err
		}
		return filePath, parts, nil
	default:
		return "", nil, ErrDuplicateFiles
	}
}

// 上传分片
func (uc *AlgoManagementUsecase) UploadPart(ctx context.Context, provider, md5 string, partNumber int64, reader io.ReadSeeker) (etag string, err error) {
	// 检查是否已经上传
	uploads, count, err := uc.fileUploadUc.List(ctx, &FileUploadFilter{
		ProviderEq: provider,
		Md5Eq:      md5,
	})
	if err != nil {
		return "", err
	}

	// 文件未成功初始化上传，即 CheckUploads 失败
	if count == 0 {
		return "", ErrFileUploadingInitFailed
	}

	// 检查是否有重复文件
	if count != 1 {
		return "", ErrDuplicateFiles
	}

	// 检查文件状态
	if uploads[0].Status != FileUploadStatusUploading {
		return "", ErrFileIsNotUploading
	}

	return uc.ossUc.UploadPart(ctx, uploads[0].Key, uploads[0].UploadID, partNumber, reader)
}

// 完成上传，各个模式下metaData均为完整meta信息
func (uc *AlgoManagementUsecase) CompleteMultipartUpload(ctx context.Context, provider, md5 string, partsNum int64, metaData string) (uint64, string, error) {
	// 检查是否已经上传
	uploads, count, err := uc.fileUploadUc.List(ctx, &FileUploadFilter{
		ProviderEq: provider,
		Md5Eq:      md5,
	})
	if err != nil {
		return 0, "", err
	}

	// 检查是否有重复文件
	if count != 1 {
		return 0, "", ErrDuplicateFiles
	}

	// 检查文件状态
	if uploads[0].Status != FileUploadStatusUploading && uploads[0].Status != FileUploadStatusSuccess {
		return 0, "", ErrFileIsNotUploading
	}

	// 完成上传
	var etag string
	if uploads[0].Status == FileUploadStatusUploading {
		etag, err = uc.ossUc.CompleteMultipartUpload(ctx, uploads[0].Key, uploads[0].UploadID, partsNum)
		if err != nil {
			return 0, "", err
		}
	}

	// 更新文件表
	newUpload := *uploads[0]
	newUpload.Status = FileUploadStatusSuccess
	newUpload.Etag = etag
	newUpload.Meta = metaData

	var meta Meta
	if err := json.Unmarshal([]byte(metaData), &meta); err != nil {
		return 0, "", err
	}
	newUpload.Type = meta.Type
	newUpload.AlgoGroupID = uint64(meta.AlgoGroupID)

	if err := uc.fileUploadUc.Update(context.Background(), uploads[0].ID, &newUpload); err != nil {
		return 0, "", err
	}

	return newUpload.ID, etag, nil
}

// ListUploadFiles 前端接口加上分页参数
func (uc *AlgoManagementUsecase) ListUploadFiles(ctx context.Context, provider string, pageNum, pageSize int) ([]*FileUpload, int, error) {
	if pageSize == 0 {
		pageSize = 100
	}
	return uc.fileUploadUc.List(ctx, &FileUploadFilter{
		ProviderEq: provider,
		Pagination: &pagination.Pagination{
			PageNum:  pageNum,
			PageSize: pageSize,
		},
	})
}

// DeleteFile 删除文件
// 删除文件时是否要删除算法？
func (uc *AlgoManagementUsecase) DeleteFile(ctx context.Context, id uint64) error {
	return uc.fileUploadUc.Delete(ctx, id)
}

func (uc *AlgoManagementUsecase) FindFile(ctx context.Context, id uint64) (*FileUpload, error) {
	file, err := uc.fileUploadUc.Find(ctx, id)
	if err != nil {
		return nil, err
	}
	u, err := uc.ossUc.GenerateTemporaryUrl(ctx, file.Key, 12*time.Hour)
	if err != nil {
		return nil, err
	}
	file.Url = u
	return file, nil
}

type Meta struct {
	Version       string       `json:"version"`
	Algos         []*Algorithm `json:"algorithms"`
	DeleteFiles   []string     `json:"to_delete_files"`
	Provider      string       `json:"provider"` // sophgo_city/sophgo_park/ctyun_ias
	Fields        []string     `json:"fields"`
	AlgoGroupID   uint         `json:"algo_group_id"`
	AlgoGroupName string       `json:"algo_group_name"`
	BaseDir       string       `json:"base_dir"`
	Type          string       `json:"type"` // docker/file
	Platform      string       `json:"platform"`
	DeviceModel   string       `json:"device_model"`
	IsGroupType   uint         `json:"is_group_type"`
	NeedMetaFile  uint         `json:"need_meta_file"`
}

// Install 安装算法
func (uc *AlgoManagementUsecase) Install(ctx context.Context, filePath string) error {

	//读取meta.json
	zipFile, err := zip.OpenReader(filePath)
	if err != nil {
		return err
	}
	defer zipFile.Close()
	if zipFile.Comment == "" {
		return errors.New("zip文件中comment为空")
	}

	//解析meta.json
	meta := &Meta{}
	if err := json.Unmarshal([]byte(zipFile.Comment), meta); err != nil {
		return errors.New("解析meta.json失败:" + err.Error())
	}

	// 需要额外读取meta文件的场景
	if meta.NeedMetaFile == 1 {
		meta, err = uc.ReadMetaFile(zipFile)
		if err != nil {
			return err
		}
	}

	//执行安装
	uc.log.Debug("开始安装算法")

	switch meta.Provider {
	case "sophgo_city":
		err = uc.DoSophgoCityInstall(meta, filePath)
	case "sophgo_park":
		err = uc.DoSophgoParkInstall(meta, filePath)
	case "ctyun_ias":
		err = uc.DoCtyunIasInstall(meta, filePath)
	case "ctyun_telestream":
		err = uc.DoCtyunTelestreamInstall(meta, filePath)
	default:
		return errors.New("未知算法包,Provider无法识别")
	}
	if err != nil {
		uc.log.Errorf("安装包安装算法错误: %v", err.Error())
		return err
	}

	//更新算法表
	uc.log.Debug("开始更新算法表")
	if err = uc.UpdateAlgoTableByMeta(context.Background(), meta); err != nil {
		return err
	}
	uc.log.Debug("算法包安装成功")

	return nil
}

func (uc *AlgoManagementUsecase) ReadMetaFile(zipReader *zip.ReadCloser) (*Meta, error) {
	for _, file := range zipReader.File {
		if file.Name == "meta.json" {
			rc, err := file.Open()
			if err != nil {
				return nil, err
			}
			defer rc.Close()

			// 读取 meta.json 文件内容
			metaData, err := io.ReadAll(rc)
			if err != nil {
				return nil, err
			}

			meta := &Meta{}
			if err := json.Unmarshal(metaData, meta); err != nil {
				return nil, errors.New("解析meta.json失败:" + err.Error())
			}
			return meta, nil
		}
	}
	return nil, errors.New("meta.json file not found in the zip")

}

func (uc *AlgoManagementUsecase) DockerPullInstall(ctx context.Context, image string, meta string) error {
	cmd := exec.Command("docker", "pull", image)

	err := cmd.Run()
	if err != nil {
		uc.log.Errorf("docker pull image err:%v", err.Error())
		return err
	}
	metaItem := &Meta{}
	//解析meta.json
	if err := json.Unmarshal([]byte(meta), metaItem); err != nil {
		return errors.New("解析meta.json失败:" + err.Error())
	}

	//更新算法表
	uc.log.Debug("开始更新算法表")
	if err = uc.UpdateAlgoTableByMeta(context.Background(), metaItem); err != nil {
		return err
	}
	uc.log.Debug("算法包安装成功")
	return nil
}

// 云边安装算法和单机安装算法共用
func (uc *AlgoManagementUsecase) UpdateAlgoTableByMeta(ctx context.Context, meta *Meta) error {

	ctx = context.WithValue(ctx, UpdateAlgoKey{}, "install")
	// 更新算法表
	for _, algo := range meta.Algos {
		oldAlgos, _, err := uc.algoUc.List(ctx, &AlgoFilter{NameEq: algo.Name, ProviderEq: meta.Provider})
		if err != nil {
			uc.log.Errorf("查询算法失败:%s", err.Error())
			continue
		}

		// 填充算法组信息
		algo.Provider = meta.Provider
		algo.Platform = meta.Platform
		algo.DeviceModel = meta.DeviceModel
		algo.IsGroupType = meta.IsGroupType
		if meta.AlgoGroupName != "" && meta.Version != "" {
			algo.AlgoGroupID = meta.AlgoGroupID
			algo.AlgoGroupName = meta.AlgoGroupName
			algo.AlgoGroupVersion = meta.Version
		}
		if len(oldAlgos) >= 1 { // 有重名的时候，只替换
			newAlgo := oldAlgos[0]
			if len(meta.Fields) == 0 { //全部替换
				newAlgo = algo
			} else { // 部分替换
				for _, field := range meta.Fields {
					UpdateField(algo, newAlgo, field)
				}
			}

			if err = uc.algoUc.Update(ctx, oldAlgos[0].ID, newAlgo, nil); err != nil {
				uc.log.Errorf("更新算法失败:%s", err.Error())
			}
		} else { // 新增
			newAlgo := algo
			if _, err := uc.algoUc.Create(ctx, newAlgo, nil); err != nil {
				uc.log.Errorf("创建算法失败:%s", err.Error())
			}
		}
	}

	// 更新算法组信息，部分算法更新时整体组版本号也需要更新
	if meta.AlgoGroupName != "" && meta.Version != "" {
		if err := uc.algoUc.UpdateAlgoGroupVersion(ctx, meta.AlgoGroupName, meta.Version); err != nil {
			uc.log.Errorf("更新算法组版本失败: %s", err.Error())
		}
	}

	return nil
}

/*
使用 source 中的字段更新 target 中的同名字段
fieldName 使用 struct 中的公开字段名，首字母大写，如 Algorithm 中的 'AlgoGroupName'，非 json 字段名
*/
func UpdateField(source, target interface{}, fieldName string) error {
	sourceValue := reflect.ValueOf(source).Elem()
	targetValue := reflect.ValueOf(target).Elem()

	sourceField := sourceValue.FieldByName(fieldName)
	targetField := targetValue.FieldByName(fieldName)

	if !sourceField.IsValid() {
		return fmt.Errorf("source field %s not found", fieldName)
	}

	if !targetField.IsValid() {
		return fmt.Errorf("target field %s not found", fieldName)
	}

	if targetField.Kind() != sourceField.Kind() {
		return fmt.Errorf("field %s types do not match", fieldName)
	}

	targetField.Set(sourceField)
	return nil
}

type CityResponse struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Success bool   `json:"success"`
}

// 算能城管版安装算法
func (uc *AlgoManagementUsecase) DoSophgoCityInstall(meta *Meta, filePath string) error {
	// 获取解压目录
	baseDir := "/data/app/imageai"
	if meta.BaseDir != "" {
		baseDir = meta.BaseDir
		uc.log.Infof("使用至newBaseDir:%s", baseDir)
	}

	// 解压
	unzipCmd := exec.Command("unzip", "-o", filePath, "-d", baseDir)
	if err := unzipCmd.Run(); err != nil {
		uc.log.Errorf("解压失败:%s", err.Error())
		return err
	}

	if err := exec.Command("chmod", "-R", "755", baseDir).Run(); err != nil {
		uc.log.Errorf("更改 %s 文件权限失败:%s", baseDir, err.Error())
		return err
	}
	uc.log.Info("城管算法包解压成功")

	// 重启算法服务
	responseBody, err := getURLContent("http://127.0.0.1:8879/api/v1/aiplatform/event/reload")
	if err != nil {
		return err
	}

	// 解析响应
	var cityResponse CityResponse
	if err = json.Unmarshal([]byte(responseBody), &cityResponse); err != nil {
		return err
	}
	if cityResponse.Code != 0 {
		return errors.New("算能城管算法服务重启失败:" + cityResponse.Msg)
	}
	uc.log.Info("城管算法服务重启成功")

	return nil
}

// 算能园区版全量安装算法
func (uc *AlgoManagementUsecase) DoSophgoParkInstall(meta *Meta, filePath string) error {
	parkBaseDir := "/data/yuanqu"
	if meta.BaseDir != "" {
		parkBaseDir = meta.BaseDir
		uc.log.Infof("使用至newBaseDir:%s", parkBaseDir)
	}

	// 获取解压目录
	var unzipTopDir string
	var err error
	if unzipTopDir, err = GetZipTopDir(filePath); err != nil {
		uc.log.Errorf("获取解压目录失败:%s", err.Error())
		return err
	}
	destinationDir := filepath.Join(parkBaseDir, unzipTopDir) + "/ais" // 新的目录地址,/data/yuanqu/301/ais
	uc.log.Infof("新的目录地址:%s", destinationDir)

	// 解压
	unzipCmd := exec.Command("unzip", "-o", filePath, "-d", parkBaseDir)
	if err := unzipCmd.Run(); err != nil {
		uc.log.Errorf("解压失败:%s", err.Error())
		return err
	}
	uc.log.Info("解压成功")

	// 获取pid，偶发会出现两个服务进程，获取启动时间最早的
	pidOutput, err := exec.Command("sh", "-c", "ps aux | grep 'ais_app -flagfile' | grep -v grep | sort -k22 | head -n 1 | awk '{print $2}'").Output()
	if err != nil {
		uc.log.Errorf("Error getting PID: %s", err.Error())
		return err
	}
	pid := strings.TrimSpace(string(pidOutput))

	// 从pid获取工作目录
	pwdxOutput, err := exec.Command("pwdx", pid).Output()
	if err != nil {
		uc.log.Errorf("Error getting working directory from PID %s, err: %s", pid, err.Error())
		return err
	}
	pwdxOutputParts := strings.Fields(string(pwdxOutput))
	serviceDir := pwdxOutputParts[len(pwdxOutputParts)-1] + "/.." // 原目录地址，/data/yuanqu/3.0.0_051b6c_20230207_release/aibox
	uc.log.Infof("原目录地址:%s", serviceDir)

	// 停止服务
	stopCmd := exec.Command("sh", "-c", "cd "+serviceDir+" && python start.py stop")
	if err := stopCmd.Run(); err != nil {
		uc.log.Errorf("Error stopping service:%s", err.Error())
		return err
	}
	uc.log.Info("停止服务成功")

	// 复制license
	if serviceDir != destinationDir {
		authLicenseSrc := filepath.Join(serviceDir, "aibox/auth/licence")
		authLicenseDest := filepath.Join(destinationDir, "aibox/auth/licence")
		cpCmd := exec.Command("cp", authLicenseSrc, authLicenseDest)
		if err := cpCmd.Run(); err != nil {
			uc.log.Errorf("Error cp license:%s", err.Error())
			return err
		}
		uc.log.Info("复制license成功")
	}
	// 修改flagfile.conf
	flagConfPath := filepath.Join(destinationDir, "conf/flagfile.conf")
	sedCmd := exec.Command("sed", "-i", "s/auth_type=2/auth_type=1/g", flagConfPath)
	if err := sedCmd.Run(); err != nil {
		uc.log.Errorf("Error sed flag file:%s", err.Error())
		return err
	}
	uc.log.Info("修改flagfile.conf成功")

	// 启动服务
	startCmd := exec.Command("sh", "-c", "cd "+destinationDir+" && python start.py start")
	if err := startCmd.Run(); err != nil {
		uc.log.Errorf("Error starting service:%s", err.Error())
		return err
	}
	uc.log.Info("启动服务成功")
	return nil
}

// GetZipTopDir 获取zip文件中的顶层目录
func GetZipTopDir(zipFilePath string) (string, error) {
	var topLevelDir string
	// 打开 zip 文件
	zipFile, err := zip.OpenReader(zipFilePath)
	if err != nil {
		return "", err
	}
	defer zipFile.Close()

	// 遍历 zip 文件中的每个文件/目录
	for _, file := range zipFile.File {

		// 返回安装包中的顶层目录
		if topLevelDir == "" {
			topLevelDir = file.Name
			if strings.Contains(topLevelDir, "/") {
				topLevelDir = strings.Split(topLevelDir, "/")[0]
			}
		}
	}
	return topLevelDir, nil
}

// getURLContent 获取URL内容
func getURLContent(url string) (string, error) {
	// 发送 GET 请求
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	// 创建一个缓冲区来存储响应内容
	var contentBuilder strings.Builder

	// 将响应内容流式地复制到缓冲区中
	_, err = io.Copy(&contentBuilder, response.Body)
	if err != nil {
		return "", err
	}

	return contentBuilder.String(), nil
}

// IAS 自有算法安装
func (uc *AlgoManagementUsecase) DoCtyunIasInstall(meta *Meta, filePath string) error {
	// 解压
	tempDir := "./temp"
	unzipCmd := exec.Command("unzip", "-o", filePath, "-d", tempDir)
	if err := unzipCmd.Run(); err != nil {
		uc.log.Errorf("解压失败:%s", err.Error())
		return err
	}
	defer func() {
		os.RemoveAll(tempDir)
		uc.log.Info("删除临时解压目录成功")
	}()
	uc.log.Info("算法包解压成功")

	// 读取算法镜像文件，支持多个
	files, err := os.ReadDir(tempDir)
	if err != nil {
		uc.log.Errorf("读取镜像文件失败:%s", err.Error())
		return err
	}

	if len(files) == 0 {
		uc.log.Error("算法镜像压缩包为空")
		return errors.New("算法镜像压缩包为空")
	}

	uc.log.Infof("开始加载 %d 个算法镜像", len(files))

	var filesCount int
	for i := range files {
		if files[i].IsDir() {
			continue
		}
		fileName := files[i].Name()
		// 默认后缀是 .tar
		if strings.HasSuffix(fileName, ".tar") {
			filesCount += 1
			loadCmd := exec.Command("docker", "load", "-i", filepath.Join(tempDir, fileName))
			if err := loadCmd.Run(); err != nil {
				uc.log.Errorf("加载算法镜像 %s 失败, err: %s", fileName, err.Error())
				return err
			}
			uc.log.Infof("加载镜像 %s 成功", fileName)
		}
	}

	if filesCount == 0 {
		return errors.New("没有可安装的镜像")
	}

	return nil
}

// DoCtyunTelestreamInstall IAS telestream自有算法安装
func (uc *AlgoManagementUsecase) DoCtyunTelestreamInstall(meta *Meta, filePath string) error {
	// 解压
	tempDir := "/data/telestream"
	unzipCmd := exec.Command("unzip", "-o", filePath, "-d", tempDir)
	if err := unzipCmd.Run(); err != nil {
		uc.log.Errorf("解压失败:%s", err.Error())
		return err
	}
	uc.log.Info("算法包解压成功")

	return RestartCtyunTelestream(uc.log)
}

// RestartCtyunTelestream IAS telestream重启
func RestartCtyunTelestream(log *log.Helper) error {
	// 启动telestream
	cmdExec := exec.Command("/bin/bash", "-c", "systemctl restart telestream.service")
	_, err := cmdExec.Output()
	if err != nil {
		log.Errorf("telestream算法启动失败:%s", err.Error())
		return errors.New("telestream算法启动失败")
	}
	log.Info("telestream算法服务重启成功")

	return nil
}

// findFirstTarFile 查找第一个 .tar 文件，如果没有找到则返回空字符串。
func (uc *AlgoManagementUsecase) FindFirstTarFile(folderPath string) string {
	files, err := os.ReadDir(folderPath)
	if err != nil {
		uc.log.Warnf("Error reading directory:%v", err)
		return ""
	}

	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".tar" {
			return file.Name()
		}
	}

	return ""
}

// extractImageID 从输出中提取镜像 ID，如果找不到则返回空字符串。
func (uc *AlgoManagementUsecase) ExtractImageID(output string) string {
	// 在输出中查找 "Loaded image: " 来获取镜像 ID
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "Loaded image: ") {
			parts := strings.Split(line, " ")
			if len(parts) > 2 {
				return parts[2]
			}
		}
	}
	return ""
}
