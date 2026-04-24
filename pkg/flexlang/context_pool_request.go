// Copyright 2026 FlexCDN root@flexcdn.cn. All rights reserved. Official site: https://flexcdn.cn .

package flexlang

type RequestContextPool struct {
	ch chan *RequestContext
}

func NewRequestContextPool(poolSize int) *RequestContextPool {
	if poolSize <= 0 {
		poolSize = 32
	}
	return &RequestContextPool{
		ch: make(chan *RequestContext, poolSize),
	}
}

func (this *RequestContextPool) Get() *RequestContext {
	select {
	case ctx := <-this.ch:
		return ctx
	default:
		return NewRequestContext()
	}
}

func (this *RequestContextPool) Put(ctx *RequestContext) {
	select {
	case this.ch <- ctx:
	default:
	}
}
