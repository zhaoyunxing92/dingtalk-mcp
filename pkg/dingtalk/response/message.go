package response

type SendCorpConversationResponse struct {
	Response
	TaskId int `json:"task_id"`
}
