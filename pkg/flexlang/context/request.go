// Copyright 2026 FlexCDN root@flexcdn.cn. All rights reserved. Official site: https://flexcdn.cn .

package context

import (
	"net/http"
)

type Request interface {
	// Id 获取当前请求ID
	//
	// @expr id
	Id() string

	// ServerInfo 获取网站信息
	//
	// @expr serverInfo
	ServerInfo() RequestServerInfo

	// NodeInfo 获取节点信息
	//
	// @expr nodeInfo
	NodeInfo() RequestNodeInfo

	// URL 获取当前URL
	//
	// @expr url
	URL() string

	// Path 获取当前URL路径
	//
	// @expr path
	Path() string

	// URI 获取当前URI
	//
	// 一般为路径加参数
	//
	// @expr uri
	URI() string

	// SetURI 设置当前回源的URI
	//
	// @expr setURI
	SetURI(uri string) bool

	// Query 读取当前查询参数对象
	//
	// @expr query
	Query() URLQuery

	// Host 读取当前主机地址
	//
	// 如果正在访问的地址使用非标准端口号，那么包含端口
	//
	// @expr host
	Host() string

	// RemoteAddr 读取客户端地址
	//
	// @expr remoteAddr
	RemoteAddr() string

	// RawRemoteAddr 读取直接连接的客户端地址
	//
	// 如果用户使用了代理，那么将会读取到代理的地址
	//
	// @expr rawRemoteAddr
	RawRemoteAddr() string

	// RemotePort 读取客户端端口
	//
	// @expr remotePort
	RemotePort() int

	// Method 读取请求方法
	//
	// 返回如GET、POST、HEAD之类的请求方法
	//
	// @expr method
	Method() string

	// ContentLength 请求内容长度
	//
	// 即客户端发送的请求内容长度，GET方法请求的内容通常长度为0
	//
	// @expr contentLength
	ContentLength() int64

	// TransferEncoding 请求使用的编码
	TransferEncoding() string

	// Proto 请求使用的协议
	//
	// 比如 HTTP/1.1
	Proto() string

	// ProtoMajor 请求使用的协议主版本
	//
	// 比如 HTTP/1.0 中的 1
	ProtoMajor() int

	// ProtoMinor 请求使用的协议小版本
	//
	// 比如 HTTP/1.0 中的 0
	ProtoMinor() int

	// Cookie 读取请求发送的Cookie值
	Cookie(name string) string

	// Header 读取所有请求发送的报头
	Header() http.Header

	// SetHeader 设置请求报头
	//
	// 用来修改发送到源站的请求报头
	SetHeader(name string, values ...string) bool

	// DeleteHeader 删除请求报头
	//
	// 用来修改发送到源站的请求报头
	DeleteHeader(name string) bool

	// SetAttr 设置请求属性
	//
	// 可以记录到访问日志的额外信息
	SetAttr(name string, value string) bool

	// SetVar 设置自定义变量
	//
	// 以便可以在 $req.format(name) 函数中使用
	SetVar(name string, value string) bool

	// Format 根据请求信息格式化内容
	//
	// 在FlexCDN中，可以使用请求变量，以下是几个例子：
	// ~~~javascript
	// $req.format("${requestId}") // => 请求ID
	// $req.format("${geo.country.name}") // => 类似于“中国”
	// $req.format("${geo.province.name}") // => 类似于“湖北省”
	// ~~~
	Format(format string) string

	// Done 完成请求
	//
	// 不再继续执行
	Done() bool

	// Close 关闭请求
	Close() bool

	// Allow 允许访问
	//
	// 即将当前请求设置为白名单直接通过WAF检查
	Allow() bool
}
