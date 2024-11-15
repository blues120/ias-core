package file

import (
	"errors"
	"os"
)

func PathExists(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err == nil {
		if fi.IsDir() {
			return true, nil
		}
		return false, errors.New("存在同名文件")
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// EnsureDir 创建文件夹,已存在返回true,新建返回false
func EnsureDir(path string) (bool, error) {
	exist, err := PathExists(path)
	if err != nil {
		return exist, err
	}
	if !exist {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return exist, err
		}
	}
	return exist, nil
}
