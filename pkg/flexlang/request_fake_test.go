// Copyright 2021 FlexCDN root@flexcdn.cn. All rights reserved.

package flexlang_test

import (
	"net/http"
	"strconv"
	"time"

	"github.com/tjbrains/flexlang/pkg/flexlang"
)

type FakeRequest struct {
	host       string
	remoteAddr string
	header     http.Header
	uri        string
}

func NewFakeRequest() *FakeRequest {
	return &FakeRequest{}
}

func (this *FakeRequest) Id() string {
	return strconv.FormatInt(time.Now().UnixMicro(), 10)
}

func (this *FakeRequest) ServerInfo() flexlang.RequestServerInfo {
	return flexlang.RequestServerInfo{
		Id: 123,
	}
}

func (this *FakeRequest) NodeInfo() flexlang.RequestNodeInfo {
	return flexlang.RequestNodeInfo{
		Id: 456,
	}
}

func (this *FakeRequest) URL() string {
	return "https://example.com/hello?name=Lily"
}

func (this *FakeRequest) Path() string {
	return "/hello"
}

func (this *FakeRequest) URI() string {
	return this.uri
}

func (this *FakeRequest) SetURI(uri string) bool {
	this.uri = uri
	return true
}

func (this *FakeRequest) Host() string {
	return this.host
}

func (this *FakeRequest) RemoteAddr() string {
	return this.remoteAddr
}

func (this *FakeRequest) RawRemoteAddr() string {
	return "127.0.0.1:12345"
}

func (this *FakeRequest) RemotePort() int {
	return 80
}

func (this *FakeRequest) Method() string {
	return http.MethodGet
}

func (this *FakeRequest) ContentLength() int64 {
	return 1024
}
func (this *FakeRequest) TransferEncoding() string {
	return "gzip"
}

func (this *FakeRequest) Proto() string {
	return "HTTP/1.2"
}

func (this *FakeRequest) ProtoMajor() int {
	return 1
}

func (this *FakeRequest) ProtoMinor() int {
	return 2
}

func (this *FakeRequest) Cookie(name string) string {
	return "cookie"
}

func (this *FakeRequest) Header() http.Header {
	return this.header
}

func (this *FakeRequest) SetHeader(name string, values ...string) bool {
	if this.header == nil {
		this.header = http.Header{}
	}
	this.header[name] = values

	return true
}

func (this *FakeRequest) DeleteHeader(name string) bool {
	if this.header == nil {
		return true
	}
	delete(this.header, name)

	return true
}

func (this *FakeRequest) SetAttr(name string, value string) bool {
	return true
}

func (this *FakeRequest) SetVar(name string, value string) bool {
	return true
}

func (this *FakeRequest) Format(s string) string {
	return s
}

func (this *FakeRequest) Done() bool {
	return true
}

func (this *FakeRequest) Allow() bool {
	return true
}

func (this *FakeRequest) Close() bool {
	return true
}
