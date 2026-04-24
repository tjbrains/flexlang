// Copyright 2026 FlexCDN root@flexcdn.cn. All rights reserved. Official site: https://flexcdn.cn .

package flexlang

import "github.com/expr-lang/expr"

type RequestVM struct {
	ctxPool *RequestContextPool
}

func NewRequestVM(concurrent int) *RequestVM {
	if concurrent <= 0 {
		concurrent = 1 << 10
	}

	return &RequestVM{
		ctxPool: NewRequestContextPool(concurrent),
	}
}

func (this *RequestVM) Compile(code string) (*Program, error) {
	var ctx = this.ctxPool.Get()
	program, err := expr.Compile(code, expr.Env(ctx), expr.Patch(NewVisitor()))
	this.ctxPool.Put(ctx)

	if err != nil {
		return nil, err
	}
	return NewProgram(program), nil
}

func (this *RequestVM) Run(program *Program, req Request, resp Response) (any, error) {
	var ctx = this.ctxPool.Get()
	ctx.Ctx.Req = req
	ctx.Ctx.Resp = resp
	result, err := expr.Run(program.Raw(), ctx)
	this.ctxPool.Put(ctx)
	return result, err
}

func (this *RequestVM) Eval(code string, req Request, resp Response) (any, error) {
	program, err := this.Compile(code)
	if err != nil {
		return nil, err
	}

	return this.Run(program, req, resp)
}
