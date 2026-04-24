// Copyright 2026 FlexCDN root@flexcdn.cn. All rights reserved. Official site: https://flexcdn.cn .

package flexlang

type RequestContext struct {
	BasicContext

	Req  Request  `expr:"$req"`
	Resp Response `expr:"$resp"`
}

func NewRequestContext() *RequestContext {
	return &RequestContext{
		BasicContext: *NewBasicContext(),
	}
}
