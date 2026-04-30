// Copyright 2026 FlexCDN root@flexcdn.cn. All rights reserved. Official site: https://flexcdn.cn .

package context

type BasicContextPool struct {
	ch chan *BasicContext
}

func NewBasicContextPool(poolSize int) *BasicContextPool {
	if poolSize <= 0 {
		poolSize = 32
	}
	return &BasicContextPool{
		ch: make(chan *BasicContext, poolSize),
	}
}

func (this *BasicContextPool) Get() *BasicContext {
	select {
	case ctx := <-this.ch:
		return ctx
	default:
		return NewBasicContext()
	}
}

func (this *BasicContextPool) Put(ctx *BasicContext) {
	select {
	case this.ch <- ctx:
	default:
	}
}
