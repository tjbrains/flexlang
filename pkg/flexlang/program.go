// Copyright 2026 FlexCDN root@flexcdn.cn. All rights reserved. Official site: https://flexcdn.cn .

package flexlang

import "github.com/expr-lang/expr/vm"

type Program struct {
	rawProgram *vm.Program
}

func NewProgram(rawProgram *vm.Program) *Program {
	return &Program{
		rawProgram: rawProgram,
	}
}

func (this *Program) Raw() *vm.Program {
	return this.rawProgram
}
