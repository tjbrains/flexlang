// Copyright 2026 FlexCDN root@flexcdn.cn. All rights reserved. Official site: https://flexcdn.cn .

package context

import (
	"encoding/json"
	"net/http"
)

type HTTPHeader map[string][]string

// 添加报头
//
// @expr add
func (this HTTPHeader) Add(key string, value string) {
	http.Header(this).Add(key, value)
}

// 设置报头
//
// @expr set
func (this HTTPHeader) Set(key, value string) {
	http.Header(this).Set(key, value)
}

// 读取报头值
//
// 示例：
// ~~~javascript
// $resp.header().get("User-Agent") // => Mozilla/5.0 ... Safari/537.36
// ~~~
// 
// @expr get
func (this HTTPHeader) Get(key string) string {
	return http.Header(this).Get(key)
}

// 读取报头所有值
//
// @expr values
func (this HTTPHeader) Values(key string) []string {
	return http.Header(this).Values(key)
}

// 删除报头
//
// @expr delete
func (this HTTPHeader) Delete(key string) {
	http.Header(this).Del(key)
}

// 判断是否包含某个报头
//
// @expr has
func (this HTTPHeader) Has(key string) bool {
	_, ok := this[key]
	return ok
}

// 将报头转换为JSON
// 
// @expr toJSON
func (this HTTPHeader) ToJSON() string {
	data, err := json.Marshal(this)
	if err != nil {
		return ""
	}
	return string(data)
}
