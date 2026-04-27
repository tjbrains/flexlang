// Copyright 2026 FlexCDN root@flexcdn.cn. All rights reserved. Official site: https://flexcdn.cn .

package flexlang

import (
	"strings"
	"sync"
	"unicode"

	"github.com/expr-lang/expr"
	"github.com/tjbrains/flexlang/internal/context"
)

var shardBasicVM *BasicVM
var basicVMOnce = sync.Once{}

type BasicVM struct {
	ctxPool *context.BasicContextPool

	ctxHandler func(ctx *context.BasicContext)
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
		ctxPool: context.NewBasicContextPool(concurrent),
	}
}

func (this *BasicVM) Compile(code string) (*Program, error) {
	code = strings.TrimRightFunc(code, unicode.IsSpace)

	if len(code) == 0 {
		return NewEmptyProgram(), nil
	}

	var ctx = this.ctxPool.Get()
	program, err := expr.Compile(code, expr.Env(ctx), expr.Patch(ctx.Visitor()))
	this.ctxPool.Put(ctx)

	if err != nil {
		var ok bool
		code, ok = FixError(code, err)
		if ok {
			ctx = this.ctxPool.Get()
			program, err = expr.Compile(code, expr.Env(ctx), expr.Patch(ctx.Visitor()))
			this.ctxPool.Put(ctx)
		}
	}

	if err != nil {
		return nil, TrimError(err)
	}

	return NewProgram(program), nil
}

func (this *BasicVM) Run(program *Program) (any, error) {
	if program.IsEmpty() {
		return "", nil
	}

	var ctx = this.ctxPool.Get()
	if this.ctxHandler != nil {
		this.ctxHandler(ctx)
	}
	result, err := expr.Run(program.Raw(), ctx)
	this.ctxPool.Put(ctx)
	return result, TrimError(err)
}

func (this *BasicVM) Eval(code string) (any, error) {
	program, err := this.Compile(code)
	if err != nil {
		return nil, err
	}

	return this.Run(program)
}

func (this *BasicVM) WithCtxHandler(handler func(ctx *context.BasicContext)) {
	this.ctxHandler = handler
}
