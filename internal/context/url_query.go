// Copyright 2026 FlexCDN root@flexcdn.cn. All rights reserved. Official site: https://flexcdn.cn .

package context

import (
	"encoding/json"
	"net/url"
)

type URLQuery struct {
	urlQuery url.Values
}

func NewURLQuery(urlQuery url.Values) URLQuery {
	return URLQuery{
		urlQuery: urlQuery,
	}
}

func (this URLQuery) Get(name string) string {
	return this.urlQuery.Get(name)
}

func (this URLQuery) Values(name string) []string {
	return this.urlQuery[name]
}

func (this URLQuery) Has(name string) bool {
	return this.urlQuery.Has(name)
}

func (this URLQuery) Set(name string, value string) bool {
	this.urlQuery.Set(name, value)
	return true
}
func (this URLQuery) Delete(name string) bool {
	this.urlQuery.Del(name)
	return true
}

func (this URLQuery) Encode() string {
	return this.urlQuery.Encode()
}

func (this URLQuery) ToJSON() string {
	data, _ := json.Marshal(this.urlQuery)
	return string(data)
}

func (this URLQuery) ToPairJSON() string {
	data, _ := json.Marshal(this.ToPair())
	return string(data)
}

func (this URLQuery) ToMap() map[string][]string {
	return this.urlQuery
}

func (this URLQuery) ToPair() map[string]string {
	var pair = map[string]string{}
	for k, v := range this.urlQuery {
		if len(v) > 0 {
			pair[k] = v[0]
		}
	}
	return pair
}
