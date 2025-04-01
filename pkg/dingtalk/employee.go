package dingtalk

import (
	"net/http"

	"github.com/zhaoyunxing92/dingtalk-mcp/pkg/dingtalk/constant"
	"github.com/zhaoyunxing92/dingtalk-mcp/pkg/dingtalk/response"
)

// GetEmployeesCount 获取员工数量
func (ds *DingTalk) GetEmployeesCount(onlyActive bool) (int, error) {
	var (
		body = map[string]bool{"only_active": onlyActive}
		data = &response.CountUserResponse{}
		err  error
	)
	if err = ds.Request(http.MethodPost, constant.GetUserCountKey, nil, body, data); err != nil {
		return 0, err
	}
	return data.Result.Count, nil
}

// GetSimpleEmployees 获取部门用户基础信息
func (ds *DingTalk) GetSimpleEmployees(deptId, cursor, size int) (*response.ListUserSimpleResponse, error) {
	var (
		data = &response.ListUserSimpleResponse{}
		body = map[string]interface{}{
			"dept_id": deptId,
			"cursor":  cursor,
			"size":    size,
		}
		err error
	)
	if err = ds.Request(http.MethodPost, constant.GetDeptSimpleUserKey, nil, body, data); err != nil {
		return nil, err
	}
	return data, nil
}

//
//func (ds *DingTalk) GetAdminList() {
//	var (
//		data = &models.ListAdminResponse{}
//		err  error
//	)
//	if err = ds.Request(http.MethodPost, constant.GetUserCountKey, nil, nil, data); err != nil {
//		return 0, err
//	}
//	return data.Result.Count, nil
//}
