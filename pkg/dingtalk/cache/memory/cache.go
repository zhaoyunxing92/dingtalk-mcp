package memory

import (
	"time"

	"github.com/pkg/errors"

	"github.com/zhaoyunxing92/dingtalk-mcp/pkg/dingtalk/cache"
)

type Cache struct {
	token map[string]cache.Data
}

func NewCache() *Cache {
	return &Cache{token: make(map[string]cache.Data)}
}

func (c *Cache) Set(key string, data cache.Data) error {
	c.token[key] = data
	return nil
}

func (c *Cache) Get(key string) (string, error) {
	if data, ok := c.token[key]; ok {
		created := data.CreatedAt()
		expires := data.ExpiresIn()
		if time.Now().Unix() > created+int64(expires-60) {
			return "", errors.New("token is already expired")
		} else {
			return data.Get(), nil
		}
	} else {
		return "", errors.New("token not found")
	}
}
