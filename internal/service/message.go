package service

import (
	"context"
	"strconv"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	"github.com/zhaoyunxing92/dingtalk-mcp/pkg/dingtalk"
	"github.com/zhaoyunxing92/dingtalk-mcp/pkg/dingtalk/models/message"
)

type Message struct {
	client *dingtalk.DingTalk
}

func NewMessageService(client *dingtalk.DingTalk) *Message {
	return &Message{client: client}
}

// SendCorpConversation 发送工作通知
func (msg *Message) SendCorpConversation(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	userIds := req.Params.Arguments["userIds"].(string)
	text := req.Params.Arguments["context"].(string)
	if taskId, err := msg.client.DoSendCorpConversation(message.NewTextMessage(text), userIds); err != nil {
		return nil, err
	} else {
		return mcp.NewToolResultText(strconv.Itoa(taskId)), nil
	}
}

// SendMarkDownCorpConversation 发送markdown格式的工作通知
func (msg *Message) SendMarkDownCorpConversation(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	userIds := req.Params.Arguments["userIds"].(string)
	title := req.Params.Arguments["title"].(string)
	content := req.Params.Arguments["content"].(string)
	if taskId, err := msg.client.DoSendCorpConversation(message.NewMarkDownMessage(title, content), userIds); err != nil {
		return nil, err
	} else {
		return mcp.NewToolResultText(strconv.Itoa(taskId)), nil
	}
}

// RecallCorpConversation 撤回工作通知
func (msg *Message) RecallCorpConversation(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	taskId := req.Params.Arguments["taskId"].(string)
	if err := msg.client.RecallCorpConversation(taskId); err != nil {
		return nil, err
	} else {
		return mcp.NewToolResultText("撤回消息成功"), nil
	}
}

// AddTools 添加工具
func (msg *Message) AddTools(svc *server.MCPServer) {
	sendCorpConversation := mcp.NewTool("send_corp_conversation",
		mcp.WithDescription("企业用户发送消息"),
		mcp.WithString("userIds",
			mcp.Required(),
			mcp.Description(`接收者的userId列表，用英文逗号分隔，最大用户列表长度：100。`)),
		mcp.WithString("context",
			mcp.Required(),
			mcp.Description("消息内容，最长不超过2048个字节"),
		))

	sendMarkDownCorpConversation := mcp.NewTool("send_markdown_corp_conversation",
		mcp.WithDescription("企业用户发送消息"),
		mcp.WithString("userIds",
			mcp.Required(),
			mcp.Description(`接收者的userId列表，用英文逗号分隔，最大用户列表长度：100。`)),
		mcp.WithString("content",
			mcp.Required(),
			mcp.Description("markdown格式的消息，最大不超过5000字符"),
		),
		mcp.WithString("title",
			mcp.Required(),
			mcp.Description("消息标题，最长不超过128个字节"),
		))

	recallCorpConversation := mcp.NewTool("recall_corp_conversation",
		mcp.WithDescription("撤回给员工的消息"),
		mcp.WithString("taskId",
			mcp.Required(),
			mcp.Description(`发送消息时钉钉返回的任务ID`)),
	)

	svc.AddTool(sendCorpConversation, msg.SendCorpConversation)
	svc.AddTool(recallCorpConversation, msg.RecallCorpConversation)
	svc.AddTool(sendMarkDownCorpConversation, msg.SendMarkDownCorpConversation)
}
