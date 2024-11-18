//go:build skip
// +build skip

package iam

import (
	"crypto/tls"
	"fmt"
	"net/http"
	goHttp "net/http"
	"testing"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/stretchr/testify/assert"
	"github.com/blues120/ias-core/conf"
	"github.com/blues120/ias-core/pkg/openapi"
)

func initIamClient() *Client {
	return NewClient(&conf.IAM{
		IamHost:         "IamHost",
		AppId:           "AppId",
		AppSecret:       "AppSecret",
		Ac:              "app",
		PrivilegeAction: "c_view",
	}, log.GetLogger())
}

func TestIamMenu(t *testing.T) {
	client := initIamClient()

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	menus, err := client.GetMenus(&UserInfo{
		TenantID: "10012188",
		UserID:   "21784111",
	}, "", 0, openapi.WithTransport(tr))
	if err != nil {
		panic(err)
	}

	menusBytes, _ := json.Marshal(menus)
	fmt.Println(string(menusBytes))
}

func TestUserListPrivileges(t *testing.T) {
	client := initIamClient()

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	privileges, err := client.GetPrivileges(&UserInfo{
		TenantID: "10012188",
		UserID:   "21784111",
	}, openapi.WithTransport(tr))
	if err != nil {
		panic(err)
	}

	privilegesBytes, _ := json.Marshal(privileges)
	fmt.Println(string(privilegesBytes))
}

func TestGetUserEnabledMenus(t *testing.T) {
	client := initIamClient()

	allMenus := []*MenuItem{
		&MenuItem{
			Plist: "system_summary_view@enterprise|user|subuser|iam|osp@0",
			Name:  "系统总览",
			State: "online",
		},
		&MenuItem{
			Plist: "system_manage@enterprise|user|subuser|iam|osp@0",
			Name:  "系统设置",
			State: "online",
			Items: []*MenuItem{
				&MenuItem{
					Plist: "system_manage_organization@enterprise|user|subuser|iam|osp@0",
					Name:  "组织架构",
					State: "offline",
				},
				&MenuItem{
					Plist: "system_manage_site@enterprise|user|subuser|iam|osp@0",
					Name:  "场所管理",
					State: "online",
				},
			},
		},
		&MenuItem{
			Plist: "application_center_view@enterprise|user|subuser|iam|osp@0",
			Name:  "应用中心",
			State: "offline",
			Items: []*MenuItem{
				&MenuItem{
					Plist: "application_center_petrol_dashboard@enterprise|user|subuser|iam|osp@0",
					Name:  "加油站大屏看板",
					State: "online",
				},
				&MenuItem{
					Plist: "application_center_petrol_details@enterprise|user|subuser|iam|osp@0",
					Name:  "加油站作业明细",
					State: "online",
				},
			},
		},
	}

	actionMap1 := map[string]struct{}{
		"system_summary_view":               struct{}{},
		"system_manage_organization":        struct{}{},
		"system_manage_site":                struct{}{},
		"application_center_petrol_details": struct{}{},
	}
	expectedMenus1 := []*MenuItem{
		&MenuItem{
			Name: "系统总览",
		},
		&MenuItem{
			Name: "系统设置",
			Items: []*MenuItem{
				&MenuItem{
					Name: "场所管理",
				},
			},
		},
		&MenuItem{
			Name: "应用中心",
			Items: []*MenuItem{
				&MenuItem{
					Name: "加油站作业明细",
				},
			},
		},
	}

	// 1. 测试普通用户权限
	menus1 := client.getEnableMenus(allMenus, actionMap1, nil)
	assert.Equal(t, expectedMenus1, menus1, "")

	// 2. 测试全部权限
	adminActionMap := map[string]struct{}{
		highestPrivilegeAction: struct{}{},
	}
	expectedMenus2 := []*MenuItem{
		&MenuItem{
			Name: "系统总览",
		},
		&MenuItem{
			Name: "系统设置",
			Items: []*MenuItem{
				&MenuItem{
					Name: "场所管理",
				},
			},
		},
		&MenuItem{
			Name: "应用中心",
			Items: []*MenuItem{
				&MenuItem{
					Name: "加油站大屏看板",
				},
				&MenuItem{
					Name: "加油站作业明细",
				},
			},
		},
	}
	menus2 := client.getEnableMenus(allMenus, nil, adminActionMap)
	assert.Equal(t, expectedMenus2, menus2, "")

	// 3. 测试父菜单配置，而子菜单未配置
	actionMap3 := map[string]struct{}{
		"application_center_view": struct{}{},
		"system_manage":           struct{}{},
	}
	expectedMenus3 := []*MenuItem{
		&MenuItem{
			Name: "系统设置",
			Items: []*MenuItem{
				&MenuItem{
					Name: "场所管理",
				},
			},
		},
		&MenuItem{
			Name: "应用中心",
			Items: []*MenuItem{
				&MenuItem{
					Name: "加油站大屏看板",
				},
				&MenuItem{
					Name: "加油站作业明细",
				},
			},
		},
	}
	menus3 := client.getEnableMenus(allMenus, nil, actionMap3)
	assert.Equal(t, expectedMenus3, menus3, "")
}

func TestGetUserAuthInfo(t *testing.T) {
	client := initIamClient()

	req, _ := goHttp.NewRequest("GET", "", nil)
	req.AddCookie(&goHttp.Cookie{
		Name:  iamSessionCookiekey,
		Value: "e4258e75-7e15-42a9-a029-6d39e5f5fb90", // change this to real session when run this test
	})
	authInfo, err := client.GetUserCurrentAuthInfo(req)
	assert.Equal(t, nil, err, "")
	fmt.Printf("%+v\n", *authInfo)
}

func TestGetUserSecretKey(t *testing.T) {
	client := initIamClient()

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	secret, err := client.GetUserSecretKey("f6de54ef70bff3389e1963165b14314f", openapi.WithTransport(tr)) // change this to real ak when run this test
	assert.Equal(t, nil, err, "")
	fmt.Printf("%+v\n", *secret)
}
