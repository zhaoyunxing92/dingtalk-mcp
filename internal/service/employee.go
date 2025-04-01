package service

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	"github.com/zhaoyunxing92/dingtalk-mcp/pkg/dingtalk"
)

type Employee struct {
	client *dingtalk.DingTalk
}

func NewEmployeeService(client *dingtalk.DingTalk) *Employee {
	return &Employee{client: client}
}

// GetSimpleEmployees 获取员工列表
func (emp *Employee) GetSimpleEmployees(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	resp, err := emp.client.GetSimpleEmployees(1, 0, 100)
	if err != nil {
		return nil, err
	}
	if marshal, err := json.Marshal(resp.Result.List); err != nil {
		return nil, err
	} else {
		return mcp.NewToolResultText(string(marshal)), nil
	}
}

// GetEmployeesCount 获取员工数量
func (emp *Employee) GetEmployeesCount(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	oa := req.Params.Arguments["only_active"].(bool)
	count, err := emp.client.GetEmployeesCount(oa)
	if err != nil {
		return nil, err
	}
	return mcp.NewToolResultText(strconv.Itoa(count)), nil
}

// AddTools 添加工具
func (emp *Employee) AddTools(svc *server.MCPServer) {
	getEmployeesCount := mcp.NewTool("get_employees_count",
		mcp.WithDescription("获取企业员工人数"),
		mcp.WithBoolean("only_active",
			mcp.Required(),
			mcp.Description(`是否包含未激活钉钉人数：
* false：包含未激活钉钉的人员数量。
* true：只包含激活钉钉的人员数量。`)))

	getSimpleEmployees := mcp.NewTool("get_simple_employees",
		mcp.WithDescription("获取企业的员工基础信息，返回用户名称和userId列表"))

	svc.AddTool(getEmployeesCount, emp.GetEmployeesCount)
	svc.AddTool(getSimpleEmployees, emp.GetSimpleEmployees)
}
