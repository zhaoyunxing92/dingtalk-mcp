/*
 * Copyright 2020 zhaoyunxing.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package file

import (
	"encoding/json"
	"os"
	"strings"
	"time"

	"github.com/zhaoyunxing92/dingtalk-mcp/pkg/dingtalk/cache"

	"github.com/pkg/errors"
)

type Cache struct {
	path string // 文件路径
	file string // 文件
}

// NewFileCache 文件缓存
// path 缓存文件路径
// file 文件名称
func NewFileCache(path, file string) *Cache {
	file = strings.Join([]string{path, file}, "/")
	return &Cache{path, file}
}

// Set 缓存
func (cache *Cache) Set(data cache.Data) (err error) {
	path := cache.path
	if _, err := os.Stat(path); os.IsNotExist(err) {
		_ = os.MkdirAll(path, os.ModePerm)
	}
	var bytes []byte
	if bytes, err = json.Marshal(data); err != nil {
		return err
	}
	if err = os.WriteFile(cache.file, bytes, 0o644); err != nil {
		return err
	}
	return err
}

// Get 获取
func (cache *Cache) Get(data cache.Data) error {
	bytes, err := os.ReadFile(cache.file)
	if err == nil {
		err = json.Unmarshal(bytes, data)
		if err == nil {
			created := data.CreatedAt()
			expires := data.ExpiresIn()
			if time.Now().Unix() > created+int64(expires-60) {
				err = errors.New("token is already expired")
			}
		}
	}
	return err
}
