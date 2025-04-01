package response

import (
	"github.com/pkg/errors"
)

// Response 响应
// {"errcode":40035,"errmsg":"缺少参数 corpid or appkey"}
type Response struct {
	Code      int    `json:"errcode"`
	Msg       string `json:"errmsg,omitempty"`
	Success   bool   `json:"success,omitempty"`
	RequestId string `json:"request_id,omitempty"`

	TraceId string `json:"requestId,omitempty"`

	// 调用结果
	Result bool `json:"result,omitempty"`
}

// Unmarshalled 统一检查返回异常异常
type Unmarshalled interface {
	CheckError(data []byte) error
}

func (res *Response) CheckError(data []byte) (err error) {
	if res.Ok() {
		return nil
	}
	return errors.New(string(data))
}

func (res *Response) Ok() bool {
	return res.Code == 0
}
