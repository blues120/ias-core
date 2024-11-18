package biz

import (
	"context"
	"errors"
	"time"

	"github.com/blues120/ias-core/data/iam"
)

var (
	ErrNoOrgRoles = errors.New("用户尚未分配组织架构")
)

type Organization struct {
	ID        uint32    // id
	Name      string    // 名称
	ParentID  uint32    // 父级Id
	IamRoleID string    // 当前不用，若 IAM 开放接口，保存 iam 对应角色 id
	CreatedAt time.Time // 创建时间
	UpdatedAt time.Time // 更新时间
}

// OrganizationListFilter 批量查询过滤条件
type OrganizationListFilter struct {
	// 批量查询
	Ids   []uint32 // 层级ID列表
	Names []string // 名称列表

	// 排序
	IsIdDesc bool
}

type OrganizationRepo interface {
	// Save 创建组织架构层级
	Save(ctx context.Context, or *Organization) (uint32, error)

	// Update 更新组织架构层级
	Update(ctx context.Context, id uint32, or *Organization) error

	// Delete 删除组织架构层级
	Delete(ctx context.Context, id uint32) error

	// FindById 根据id查询组织架构层级
	FindById(ctx context.Context, id uint32) (*Organization, error)

	// List 查询组织架构层级
	List(ctx context.Context, filter *OrganizationListFilter) ([]*Organization, error)

	// FindByParentId 根据父级id查询
	FindByParentId(ctx context.Context, parentID uint32) (*Organization, error)

	// SetRedisUserOrg redis设置用户当前层级
	SetRedisUserOrg(ctx context.Context, userId string, orgId uint32) error

	// GetRedisCurrentUserOrg redis查询用户当前设置的层级
	GetRedisCurrentUserOrg(ctx context.Context, userId string) (uint32, error)

	FindAccessOrgListById(ctx context.Context, currentUserOrg uint32) ([]uint32, error)

	// 删除某用户的组织架构缓存
	DeleteRedisCurrentUserOrg(ctx context.Context, userId string) error
}

type OrganizationUsecase struct {
	organizationRepo OrganizationRepo

	iamCli *iam.Client
}

func NewOrganizationUsecase(repo OrganizationRepo, iamCli *iam.Client) *OrganizationUsecase {
	return &OrganizationUsecase{organizationRepo: repo, iamCli: iamCli}
}

// Create 创建组织架构层级
func (uc *OrganizationUsecase) Create(ctx context.Context, or *Organization) (uint32, error) {
	return uc.organizationRepo.Save(ctx, or)
}

// Update 更新组织架构层级
func (uc *OrganizationUsecase) Update(ctx context.Context, id uint32, or *Organization) error {
	return uc.organizationRepo.Update(ctx, id, or)
}

// Delete 删除组织架构层级
func (uc *OrganizationUsecase) Delete(ctx context.Context, id uint32) error {
	return uc.organizationRepo.Delete(ctx, id)
}

// FindById 根据id查询组织架构层级
func (uc *OrganizationUsecase) FindById(ctx context.Context, id uint32) (*Organization, error) {
	return uc.organizationRepo.FindById(ctx, id)
}

// List 查询组织架构层级
func (uc *OrganizationUsecase) List(ctx context.Context, filter *OrganizationListFilter) ([]*Organization, error) {
	return uc.organizationRepo.List(ctx, filter)
}

// FindByParentId 根据父级id查询
func (uc *OrganizationUsecase) FindByParentId(ctx context.Context, parentID uint32) (*Organization, error) {
	return uc.organizationRepo.FindByParentId(ctx, parentID)
}

// SetUserOrg redis设置用户当前层级
func (uc *OrganizationUsecase) SetUserOrg(ctx context.Context, userId string, orgId uint32) error {
	// 存入缓存 有效期1个月
	return uc.organizationRepo.SetRedisUserOrg(ctx, userId, orgId)
}

// GetCurrentUserOrg redis查询用户当前设置的层级
func (uc *OrganizationUsecase) GetCurrentUserOrg(ctx context.Context, userId string) (uint32, error) {
	return uc.organizationRepo.GetRedisCurrentUserOrg(ctx, userId)
}

// FindAccessOrgListById 查询用户资源可被哪些层级访问，向上查到根节点
func (uc *OrganizationUsecase) FindAccessOrgListById(ctx context.Context, currentUserOrg uint32) ([]uint32, error) {
	return uc.organizationRepo.FindAccessOrgListById(ctx, currentUserOrg)
}

// DeleteUserOrg 删除某个用户的 org 缓存
func (uc *OrganizationUsecase) DeleteUserOrgCache(ctx context.Context, userId string) error {
	return uc.organizationRepo.DeleteRedisCurrentUserOrg(ctx, userId)
}

func (uc *OrganizationUsecase) findOrgListByNames(names []string) ([]uint32, error) {
	orgs, err := uc.List(context.Background(), &OrganizationListFilter{
		Names: names,
	})
	if err != nil {
		return nil, err
	}

	var ids []uint32
	for i := range orgs {
		ids = append(ids, orgs[i].ID)
	}

	return ids, nil
}

// GetUserOrgListById 通过用户 id 获取组织架构列表，使用 iam_client 查询，不依赖中间件设置的 ctx 上下文信息
func (uc *OrganizationUsecase) GetUserOrgListById(ctx context.Context, uinfo *iam.UserInfo) ([]uint32, error) {
	// 查询用户的组织架构角色名
	roles, err := uc.iamCli.GetOrgRoles(uinfo)
	if err != nil {
		return nil, err
	}

	if len(roles) == 0 {
		return nil, ErrNoOrgRoles
	}

	// 查询用户所属的组织架构
	orgs, err := uc.findOrgListByNames(roles)
	if err != nil {
		return nil, err
	}

	return orgs, nil
}
