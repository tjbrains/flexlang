// Copyright 2026 FlexCDN root@flexcdn.cn. All rights reserved. Official site: https://flexcdn.cn .

package context

type RequestContext struct {
	*BasicContext

	// 请求对象
	//
	// 可以通过此对象读取请求相关信息
	//
	// 示例：
	// ~~~javascript
	// $req.url() // 返回当前URL，类似于 "https://example.com/hello?name=lily"
	// $req.query().get("name") // 获取参数值，在上面URL中，结果就是 "lily"
	// ~~~
	// @prototype $req [Request](#request对象)
	Req Request `expr:"$req"`

	// 响应对象
	//
	// 可以通过此对象设置响应的内容
	//
	// 示例：
	// ~~~javascript
	// $resp.send(200, "Hello, World!")
	// ~~~
	// @prototype $resp [Response](#response对象)
	Resp Response `expr:"$resp"`
}

// NewRequestContext 构造新请求上下文对象
//
// @internal
func NewRequestContext() *RequestContext {
	return &RequestContext{
		BasicContext: NewBasicContext(),
	}
}
