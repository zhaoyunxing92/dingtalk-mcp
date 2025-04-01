package response

type SimpleEmployee struct {
	UserId string `json:"userid"`
	Name   string `json:"name"`
}

type CountUserResponse struct {
	Response

	Result struct {
		Count int `json:"count"`
	} `json:"result"`
}

type ListAdminResponse struct {
	Response

	Result []struct {
		UserId string `json:"userid"`
		// SysLevel 管理员等级，1表示主管理员，2表示子管理员
		Level int `json:"sys_level"`
	} `json:"result"`
}

type ListUserSimpleResponse struct {
	Response
	Result struct {
		HasMore    bool             `json:"has_more"`
		NextCursor int              `json:"next_cursor"`
		List       []SimpleEmployee `json:"list"`
	} `json:"result"`
}

type GetListSimple struct {
	DeptId int `json:"dept_id"`

	Cursor int `json:"cursor"`

	Size int `json:"-"`
}

func NewGetListSimple(deptId, cursor int) *GetListSimple {
	return &GetListSimple{deptId, cursor, 100}
}
