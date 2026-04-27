// Copyright 2026 FlexCDN root@flexcdn.cn. All rights reserved. Official site: https://flexcdn.cn .

package flexlang

import "github.com/expr-lang/expr/vm"

type Program struct {
	rawProgram *vm.Program
	isEmpty    bool
}

func NewProgram(rawProgram *vm.Program) *Program {
	return &Program{
		rawProgram: rawProgram,
	}
}

func NewEmptyProgram() *Program {
	return &Program{
		rawProgram: nil,
		isEmpty:    true,
	}
}

func (this *Program) IsEmpty() bool {
	return this.isEmpty
}

func (this *Program) Raw() *vm.Program {
	return this.rawProgram
}
