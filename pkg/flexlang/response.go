// Copyright 2026 FlexCDN root@flexcdn.cn. All rights reserved. Official site: https://flexcdn.cn .

package flexlang

import "net/http"

type Response interface {
	SetHeader(name string, values ...string) bool
	DeleteHeader(name string) bool
	Header() http.Header
	Send(status int, body string) bool
	SendResp(resp *http.Response) (int64, error)
	Redirect(status int, url string) bool
}
