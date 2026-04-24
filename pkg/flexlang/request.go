// Copyright 2026 FlexCDN root@flexcdn.cn. All rights reserved. Official site: https://flexcdn.cn .

package flexlang

import (
	"net/http"
)

type Request interface {
	Id() string
	ServerInfo() RequestServerInfo
	NodeInfo() RequestNodeInfo
	URL() string
	Path() string
	URI() string
	SetURI(uri string) bool
	Host() string
	RemoteAddr() string
	RawRemoteAddr() string
	RemotePort() int
	Method() string
	ContentLength() int64
	TransferEncoding() string
	Proto() string
	ProtoMajor() int
	ProtoMinor() int
	Cookie(name string) string
	Header() http.Header
	SetHeader(name string, values ...string) bool
	DeleteHeader(name string) bool
	SetAttr(name string, value string) bool
	SetVar(name string, value string) bool
	Format(format string) string
	Done() bool
	Close() bool
	Allow() bool
}
