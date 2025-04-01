package response

// AccessToken 获取企业内部应用的access_token
// https://developers.dingtalk.com/document/app/obtain-orgapp-token
type AccessToken struct {
	Response
	// Expires 过期时间
	Expires int16 `json:"expires_in"`
	// Create 创建时间
	Create int64 `json:"create"`
	// Token
	Token string `json:"access_token"`
}

func (token *AccessToken) CreatedAt() int64 {
	return token.Create
}

func (token *AccessToken) ExpiresIn() int16 {
	return token.Expires
}

func (token *AccessToken) Get() string {
	return token.Token
}
