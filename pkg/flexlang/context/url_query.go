// Copyright 2026 FlexCDN root@flexcdn.cn. All rights reserved. Official site: https://flexcdn.cn .

package context

import (
	"encoding/json"
	"net/url"
)

type URLQuery struct {
	urlQuery url.Values
}

// NewURLQuery 构造新的查询
//
// @internal
func NewURLQuery(urlQuery url.Values) URLQuery {
	return URLQuery{
		urlQuery: urlQuery,
	}
}

// Get 获取某个参数值
//
// @expr get
func (this URLQuery) Get(name string) string {
	return this.urlQuery.Get(name)
}

// Values 获取某个参数的所有值
//
// 在一个参数有很多值的时候很有用
//
// 示例：
// 对于：
// ~~~
// https://example.com?name=lily&name=lucy&name=jim
// ~~~
// 来说：
// ~~~javascript
// $req.query().values("name") // => ["lily", "lucy", "jim"]
// ~~~
//
// @expr values
func (this URLQuery) Values(name string) []string {
	return this.urlQuery[name]
}

// Has 判断某个参数值是否存在
//
// @expr has
func (this URLQuery) Has(name string) bool {
	return this.urlQuery.Has(name)
}

// Set 设置参数值
//
// @expr set
func (this URLQuery) Set(name string, value string) bool {
	this.urlQuery.Set(name, value)
	return true
}

// Delete 删除某个参数值
//
// @expr delete
func (this URLQuery) Delete(name string) bool {
	this.urlQuery.Del(name)
	return true
}

// Encode 将所有参数值编码为一个字符串
//
// @expr encode
func (this URLQuery) Encode() string {
	return this.urlQuery.Encode()
}

// ToJSON 将所有参数值转为为JSON
//
// 其中每个参数值对应一个字符串数组
//
// @expr toJSON
func (this URLQuery) ToJSON() string {
	data, _ := json.Marshal(this.urlQuery)
	return string(data)
}

// ToKV 将所有参数值转换为键值对
//
// @expr toKV
func (this URLQuery) ToKV() map[string]string {
	var kv = map[string]string{}
	for k, v := range this.urlQuery {
		if len(v) > 0 {
			kv[k] = v[0]
		}
	}
	return kv
}

// ToKVJSON 将所有参数值转换为键值对JSON
//
// 其中每个参数值对应一个字符串
//
// @expr ToKVJSON
func (this URLQuery) ToKVJSON() string {
	data, _ := json.Marshal(this.ToKV())
	return string(data)
}

// ToMap 将所有参数值转换为键值对Map
//
// @expr toMap
func (this URLQuery) ToMap() map[string][]string {
	return this.urlQuery
}
