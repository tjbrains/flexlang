// Copyright 2026 FlexCDN root@flexcdn.cn. All rights reserved. Official site: https://flexcdn.cn .

package flexlang

import (
	"sync"

	"github.com/expr-lang/expr"
)

var shardBasicVM *BasicVM
var basicVMOnce = sync.Once{}

type BasicVM struct {
	ctxPool *BasicContextPool
}

func SharedBasicVM() *BasicVM {
	basicVMOnce.Do(func() {
		shardBasicVM = NewBasicVM(32)
	})

	return shardBasicVM
}

func NewBasicVM(concurrent int) *BasicVM {
	if concurrent <= 0 {
		concurrent = 1 << 10
	}

	return &BasicVM{
		ctxPool: NewBasicContextPool(concurrent),
	}
}

func (this *BasicVM) Compile(code string) (*Program, error) {
	var ctx = this.ctxPool.Get()
	program, err := expr.Compile(code, expr.Env(ctx), expr.Patch(NewVisitor()))
	this.ctxPool.Put(ctx)

	if err != nil {
		return nil, err
	}
	return NewProgram(program), nil
}

func (this *BasicVM) Run(program *Program) (any, error) {
	var ctx = this.ctxPool.Get()
	result, err := expr.Run(program.Raw(), ctx)
	this.ctxPool.Put(ctx)
	return result, err
}

func (this *BasicVM) Eval(code string) (any, error) {
	program, err := this.Compile(code)
	if err != nil {
		return nil, err
	}

	return this.Run(program)
}
