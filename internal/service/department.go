package service

import "github.com/zhaoyunxing92/dingtalk-mcp/pkg/dingtalk"

type Department struct {
	client *dingtalk.DingTalk
}

func NewDepartmentService(client *dingtalk.DingTalk) *Department {
	return &Department{client: client}
}
