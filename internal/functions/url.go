// Copyright 2026 FlexCDN root@flexcdn.cn. All rights reserved. Official site: https://flexcdn.cn .

package functions

import (
	"net/url"
	"strconv"
)

type URL struct {
	// 主机名
	//
	// 如果有端口号的话，此值也包含端口号
	Host string `json:"host" expr:"host"`

	// 查询参数
	//
	// 不包括开始的问号
	Query string `json:"query" expr:"query"`

	// 端口号
	Port int `json:"port" expr:"port"`

	// 路径
	//
	// 从正斜杠开始
	Path string `json:"path" expr:"path"`

	// 锚点
	//
	// 不包括井号（#）部分
	Hash string `json:"hash" expr:"hash"`

	// 协议
	Scheme string `json:"scheme" expr:"scheme"`

	// 非透明数据
	//
	// 比如 `mailto:user@example.com` 中的 `opaque` 为 `user@example.com`
	Opaque string `json:"opaque" expr:"opaque"`

	// 用户信息内容
	User map[string]string `json:"user" expr:"user"`
}

// NewURL 构造新URL
//
// @internal
func NewURL(urlString string) URL {
	u, err := url.Parse(urlString)
	if err != nil {
		return URL{}
	}
	var user = map[string]string{}
	if u.User != nil {
		password, _ := u.User.Password()
		user = map[string]string{
			"username": u.User.Username(),
			"password": password,
		}
	}

	port, _ := strconv.Atoi(u.Port())

	return URL{
		Host:   u.Host,
		Query:  u.RawQuery,
		Port:   port,
		Path:   u.Path,
		Hash:   u.Fragment,
		Scheme: u.Scheme,
		Opaque: u.Opaque,
		User:   user,
	}
}
