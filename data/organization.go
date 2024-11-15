package data

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"gitlab.ctyuncdn.cn/ias/ias-core/biz"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/organization"
	"gitlab.ctyuncdn.cn/ias/ias-core/errors"
	"gitlab.ctyuncdn.cn/ias/ias-core/pkg/convert"
)

type organizationRepo struct {
	data *Data

	log *log.Helper
}

func NewOrganizationRepo(data *Data, logger log.Logger) biz.OrganizationRepo {
	return &organizationRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *organizationRepo) Save(ctx context.Context, or *biz.Organization) (uint32, error) {
	data, err := r.data.db.Organization(ctx).Create().SetOrganization(r.bizToEnt(or)).Save(ctx)
	if err != nil {
		return 0, err
	}
	return data.ID, nil
}

func (r *organizationRepo) Update(ctx context.Context, id uint32, or *biz.Organization) error {
	query := r.data.db.Organization(ctx).
		UpdateOneID(id).
		SetName(or.Name).
		SetIamRoleID(or.IamRoleID)
	return query.Exec(ctx)
}

func (r *organizationRepo) Delete(ctx context.Context, id uint32) error {
	arr, err := r.data.db.Organization(ctx).Query().Clone().Order(
		ent.Asc(organization.FieldID),
	).All(ctx)
	if err != nil {
		return err
	}
	// 删除当前层级及其子层级   子层级的id一定比当前层级的id小
	deleteIds := make([]uint32, 0)
	deleteIds = append(deleteIds, id)
	idMap := make(map[uint32]struct{})
	idMap[id] = struct{}{}
	for _, org := range arr {
		if _, ok := idMap[org.ParentID]; ok {
			deleteIds = append(deleteIds, org.ID)
			idMap[org.ID] = struct{}{}
		}
	}
	_, err = r.data.db.Organization(ctx).Delete().Where(organization.IDIn(deleteIds...)).Exec(ctx)
	return err
}

func (r *organizationRepo) FindById(ctx context.Context, id uint32) (*biz.Organization, error) {
	Org, err := r.data.db.Organization(ctx).Query().
		Where(organization.IDEQ(id)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.ErrorOrganizationNotFound("组织架构层级(id=%d)不存在", id)
		}
		return nil, err
	}
	return OrganizationEntToBiz(Org), nil
}

func (r *organizationRepo) List(ctx context.Context, filter *biz.OrganizationListFilter) ([]*biz.Organization, error) {
	query := r.data.db.Organization(ctx).Query()
	if filter != nil {
		query = r.buildQueryByFilter(query, filter)
	}
	arr, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	return OrganizationEntArrToBiz(arr), nil
}

// FindByParentId 根据父级id查询
func (r *organizationRepo) FindByParentId(ctx context.Context, parentID uint32) (*biz.Organization, error) {
	Org, err := r.data.db.Organization(ctx).Query().
		Where(organization.ParentID(parentID)).
		First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}
	return OrganizationEntToBiz(Org), nil
}

// SetRedisUserOrg redis设置用户当前层级
func (r *organizationRepo) SetRedisUserOrg(ctx context.Context, userId string, orgId uint32) error {
	// 存入缓存 有效期1个月
	return r.data.rdb.Set(ctx, GetUserOrganizationCacheKey(userId), orgId, 30*24*time.Hour).Err()
}

// GetRedisCurrentUserOrg redis查询用户当前设置的层级
func (r *organizationRepo) GetRedisCurrentUserOrg(ctx context.Context, userId string) (uint32, error) {
	payload, err := r.data.rdb.Get(ctx, GetUserOrganizationCacheKey(userId)).Result()
	if err != nil {
		return 0, err
	}
	orgId, err := strconv.ParseUint(payload, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(orgId), nil
}

func (r *organizationRepo) FindAccessOrgListById(ctx context.Context, currentUserOrg uint32) ([]uint32, error) {
	orgIds := make([]uint32, 0)
	orgIds = append(orgIds, currentUserOrg)
	// 降序查询
	arr, err := r.data.db.Organization(ctx).Query().Clone().Order(
		ent.Desc(organization.FieldID),
	).All(ctx)
	if err != nil {
		return nil, err
	}

	var parentId uint32
	// 先找出当前用户的父级id
	for _, org := range arr {
		if org.ID == currentUserOrg {
			parentId = org.ParentID
			break
		}
	}
	for _, org := range arr {
		if parentId == org.ID {
			orgIds = append(orgIds, org.ID)
			parentId = org.ParentID
			if parentId == 0 {
				break
			}
		}
	}
	// 从小到大排序
	sort.Slice(orgIds, func(i, j int) bool {
		return orgIds[i] < orgIds[j]
	})
	return orgIds, err
}

func (r *organizationRepo) DeleteRedisCurrentUserOrg(ctx context.Context, userId string) error {
	return r.data.rdb.Del(ctx, GetUserOrganizationCacheKey(userId)).Err()
}

// 查询条件
func (r *organizationRepo) buildQueryByFilter(query *ent.OrganizationQuery, filter *biz.OrganizationListFilter) *ent.OrganizationQuery {
	// 范围查询条件
	if len(filter.Ids) > 0 {
		query = query.Where(organization.IDIn(filter.Ids...))
	}
	if len(filter.Names) > 0 {
		query = query.Where(organization.NameIn(filter.Names...))
	}
	if filter.IsIdDesc {
		query = query.Order(ent.Desc(organization.FieldID))
	} else {
		query = query.Order(ent.Asc(organization.FieldID))
	}
	return query
}

func (r *organizationRepo) bizToEnt(or *biz.Organization) *ent.Organization {
	entOrganization := &ent.Organization{
		ID:        or.ID,
		CreatedAt: or.CreatedAt,
		UpdatedAt: or.UpdatedAt,
		Name:      or.Name,
		ParentID:  or.ParentID,
		IamRoleID: or.IamRoleID,
	}

	return entOrganization
}

func OrganizationEntToBiz(or *ent.Organization) *biz.Organization {
	bizOrganization := &biz.Organization{
		ID:        or.ID,
		CreatedAt: or.CreatedAt,
		UpdatedAt: or.UpdatedAt,
		Name:      or.Name,
		ParentID:  or.ParentID,
		IamRoleID: or.IamRoleID,
	}
	return bizOrganization
}

func OrganizationEntArrToBiz(arr []*ent.Organization) []*biz.Organization {
	return convert.ToArr[ent.Organization, biz.Organization](OrganizationEntToBiz, arr)
}

// 获取用户层级缓存的key
func GetUserOrganizationCacheKey(userId string) string {
	return fmt.Sprintf("organization:%s", userId)
}
