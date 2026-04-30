// Copyright 2026 FlexCDN root@flexcdn.cn. All rights reserved. Official site: https://flexcdn.cn .

package flexlang

import (
	"strings"
	"unicode"

	"github.com/expr-lang/expr"
	"github.com/tjbrains/flexlang/pkg/flexlang/context"
)

type RequestEnv struct {
	Req     context.Request
	Resp    context.Response
	Printer func(s ...string)
}

type RequestVM struct {
	ctxPool *context.RequestContextPool

	ctxHandler func(ctx *context.RequestContext)
}

func NewRequestVM(concurrent int) *RequestVM {
	if concurrent <= 0 {
		concurrent = 1 << 10
	}

	return &RequestVM{
		ctxPool: context.NewRequestContextPool(concurrent),
	}
}

func (this *RequestVM) Compile(code string) (*Program, error) {
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

func (this *RequestVM) Run(program *Program, env RequestEnv) (any, error) {
	if program.IsEmpty() {
		return "", nil
	}

	var ctx = this.ctxPool.Get()
	ctx.Req = env.Req
	ctx.Resp = env.Resp

	if env.Printer != nil {
		ctx.WithPrinter(env.Printer)
	} else {
		ctx.WithPrinter(nil) // reset
	}

	if this.ctxHandler != nil {
		this.ctxHandler(ctx)
	}

	result, err := expr.Run(program.Raw(), ctx)
	this.ctxPool.Put(ctx)
	return result, TrimError(err)
}

func (this *RequestVM) Eval(code string, requestEnv RequestEnv) (any, error) {
	program, err := this.Compile(code)
	if err != nil {
		return nil, err
	}

	return this.Run(program, requestEnv)
}

func (this *RequestVM) WithCtxHandler(handler func(ctx *context.RequestContext)) {
	this.ctxHandler = handler
}
