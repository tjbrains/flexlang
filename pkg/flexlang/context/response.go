// Copyright 2026 FlexCDN root@flexcdn.cn. All rights reserved. Official site: https://flexcdn.cn .

package context

import "net/http"

type Response interface {
	// SetHeader 设置响应报头
	SetHeader(name string, values ...string) bool

	// DeleteHeader 删除响应报头
	DeleteHeader(name string) bool

	// Header 读取所有响应报头
	Header() http.Header

	// Send 发送内容
	//
	// - status 为状态码，常用200
	Send(status int, body string) bool

	// SendResp 发送响应对象
	SendResp(resp *http.Response) (int64, error)

	// Redirect 跳转到某个URL
	//
	// - status 为跳转状态码，如307等
	Redirect(status int, url string) bool
}
