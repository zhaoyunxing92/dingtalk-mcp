package dingtalk

import (
	"net/http"
	"strings"

	"github.com/zhaoyunxing92/dingtalk-mcp/pkg/dingtalk/constant"
	"github.com/zhaoyunxing92/dingtalk-mcp/pkg/dingtalk/models/message"
	"github.com/zhaoyunxing92/dingtalk-mcp/pkg/dingtalk/response"
)

// DoSendCorpConversation 发送工作通知
func (ds *DingTalk) DoSendCorpConversation(msg message.Message, userIds string) (int, error) {
	var (
		err  error
		body = map[string]interface{}{
			"agent_id":    ds.agentId,
			"userid_list": userIds,
			"msg":         msg,
		}
		data = response.SendCorpConversationResponse{}
	)
	if err = ds.Request(http.MethodPost, constant.SendCorpConversationKey, nil, body, &data); err != nil {
		return 0, err
	}
	return data.TaskId, nil
}

// SendCorpConversation 发送工作通知
func (ds *DingTalk) SendCorpConversation(msg message.Message, userIds []string) (int, error) {
	return ds.DoSendCorpConversation(msg, strings.Join(userIds, ","))
}

// SimpleSendCorpConversation 发送工作通知
func (ds *DingTalk) SimpleSendCorpConversation(msg message.Message, userId string, userIds ...string) (int, error) {
	return ds.SendCorpConversation(msg, append(userIds, userId))
}

// RecallCorpConversation 撤回工作通知
func (ds *DingTalk) RecallCorpConversation(taskId string) error {
	var (
		body = map[string]interface{}{
			"agent_id":    ds.agentId,
			"msg_task_id": taskId,
		}
		data = response.SendCorpConversationResponse{}
	)
	return ds.Request(http.MethodPost, constant.RecallCorpConvMessageKey, nil, body, &data)
}
