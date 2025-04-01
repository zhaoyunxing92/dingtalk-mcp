package dingtalk

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/zhaoyunxing92/dingtalk-mcp/pkg/dingtalk/cache"
	"github.com/zhaoyunxing92/dingtalk-mcp/pkg/dingtalk/cache/memory"
	"github.com/zhaoyunxing92/dingtalk-mcp/pkg/dingtalk/constant"
	"github.com/zhaoyunxing92/dingtalk-mcp/pkg/dingtalk/request"
	"github.com/zhaoyunxing92/dingtalk-mcp/pkg/dingtalk/response"

	"github.com/pkg/errors"
)

type DingTalk struct {
	// 企业内部应用对应:AgentId
	agentId int

	// 企业内部应用对应:AppKey
	key string
	// 企业内部对应:AppSecret
	secret string
	// 请求客户端
	client http.Client
	// 缓存
	cache cache.Cache
}

func NewClient(agentId int, key, secret string) (ding *DingTalk) {
	if agentId == -1 {
		panic("agentId and key must be set")
	}
	if key == "" {
		panic("key must be set")
	}
	if secret == "" {
		panic("secret must be set")
	}

	return &DingTalk{
		key:     key,
		secret:  secret,
		agentId: agentId,
		client:  http.Client{Timeout: 10 * time.Second},
		cache:   memory.NewCache(),
	}
}

func (ds *DingTalk) GetAccessToken() (token string, err error) {
	var (
		ch  = ds.cache
		res = &response.AccessToken{}
	)

	if token, err = ch.Get(ds.key); err == nil {
		return token, nil
	}

	// 读取本地文件
	args := url.Values{}
	args.Set("appkey", ds.key)
	args.Set("appsecret", ds.secret)

	if err = ds.Request(http.MethodGet, constant.GetTokenKey, args, nil, res); err != nil {
		return "", err
	}
	res.Create = time.Now().Unix()
	if err = ch.Set(ds.key, res); err != nil {
		return res.Token, err
	}
	return res.Token, nil
}

func (ds *DingTalk) Request(method, path string, query url.Values, body interface{}, data response.Unmarshalled) (err error) {

	if query == nil {
		query = url.Values{}
	}

	if query.Get("access_token") == "" && path != constant.GetTokenKey && path != constant.CorpAccessToken &&
		path != constant.SuiteAccessToken && path != constant.GetAuthInfo && path != constant.GetAgentKey &&
		path != constant.ActivateSuiteKey && path != constant.GetSSOTokenKey && path != constant.GetUnactiveCorpKey &&
		path != constant.ReauthCorpKey && path != constant.GetCorpPermanentCodeKey && path != constant.GetUserAccessToken {

		var token string
		if token, err = ds.GetAccessToken(); err != nil {
			return err
		}
		// set token
		query.Set("access_token", token)
	}
	return ds.HttpRequest(method, path, query, body, data)
}

func (ds *DingTalk) HttpRequest(method, path string, query url.Values, body interface{},
	response response.Unmarshalled) error {
	var (
		req    *http.Request
		res    *http.Response
		err    error
		form   []byte // body 数据
		data   []byte // 返回数据
		client = ds.client
		uri    *url.URL
		token  string
		newApi = ds.isNewApi(path)
	)

	if newApi {
		token = query.Get("access_token")
		query.Del("access_token")
		uri, _ = url.Parse(constant.NewApi + path)
	} else {
		uri, _ = url.Parse(constant.Api + path)
	}
	uri.RawQuery = query.Encode()

	if body != nil {
		// 检查提交表单类型
		switch body.(type) {
		case request.UploadFile:
			var b bytes.Buffer
			var fw io.Writer

			w := multipart.NewWriter(&b)
			file := body.(request.UploadFile)
			fw, err = w.CreateFormFile(file.FieldName, file.FileName)
			if err != nil {
				return err
			}
			if _, err = io.Copy(fw, file.Reader); err != nil {
				return err
			}
			if err = w.Close(); err != nil {
				return err
			}
			req, _ = http.NewRequest(method, uri.String(), &b)
			req.Header.Set("Content-Type", w.FormDataContentType())
		default:
			// 表单不为空
			form, _ = json.Marshal(body)
			req, _ = http.NewRequest(method, uri.String(), bytes.NewReader(form))
			req.Header.Set("Content-Type", "application/json; charset=utf-8")
		}
	} else {
		req, _ = http.NewRequest(method, uri.String(), nil)
	}

	if newApi {
		req.Header.Set("x-acs-dingtalk-access-token", token)
	}

	if res, err = client.Do(req); err != nil {
		return err
	}

	defer func(b io.ReadCloser) { _ = b.Close() }(res.Body)

	if data, err = io.ReadAll(res.Body); err != nil {
		return err
	}

	switch res.StatusCode {
	case 200:
		if err = json.Unmarshal(data, response); err != nil {
			return err
		}
		return response.CheckError(data)
	case 400:
		return errors.Errorf("dingtalk server error,res:%s", data)
	case 404:
		return errors.Errorf("dingtalk server error,res:%s", data)
	case 500:
		return errors.Errorf("dingtalk server error,res:%s", data)
	}
	return errors.New("dingtalk seed fail")
}

func (ds *DingTalk) isNewApi(path string) bool {
	return strings.HasPrefix(path, "/v1.0/") || strings.HasPrefix(path, "/v2.0/")
}
