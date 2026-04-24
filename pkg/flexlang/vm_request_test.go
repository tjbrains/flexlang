// Copyright 2026 FlexCDN root@flexcdn.cn. All rights reserved. Official site: https://flexcdn.cn .

package flexlang_test

import (
	"testing"

	"github.com/tjbrains/flexlang/pkg/flexlang"
)

func TestRequestVM(t *testing.T) {
	var vm = flexlang.NewRequestVM(1 << 10)
	program, err := vm.Compile("1+1")
	if err != nil {
		t.Fatal(err)
	}

	result, err := vm.Run(program, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestRequestVM_Req(t *testing.T) {
	var vm = flexlang.NewRequestVM(1 << 10)

	{
		program, err := vm.Compile("ctx.req.serverInfo().id")
		if err != nil {
			t.Fatal(err)
		}

		result, err := vm.Run(program, NewFakeRequest(), nil)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%#v", result)
	}

	{
		result, err := vm.Eval("ctx.req.url()", NewFakeRequest(), nil)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%#v", result)
	}

	{
		result, err := vm.Eval(`ctx.req.setHeader("Hello", "World", "Universe"); ctx.req.header()`, NewFakeRequest(), nil)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%#v", result)
	}
}

func TestRequestVM_String(t *testing.T) {
	var vm = flexlang.NewRequestVM(1 << 10)
	program, err := vm.Compile("String.charAt('abc', 1)")
	if err != nil {
		t.Fatal(err)
	}

	result, err := vm.Run(program, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestRequestVM_Rand(t *testing.T) {
	var vm = flexlang.NewRequestVM(1 << 10)
	program, err := vm.Compile("Math.randN(10)")
	if err != nil {
		t.Fatal(err)
	}

	result, err := vm.Run(program, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func BenchmarkRequestVM_Req(b *testing.B) {
	var vm = flexlang.NewRequestVM(1 << 10)
	program, err := vm.Compile("ctx.req.path()")
	if err != nil {
		b.Fatal(err)
	}

	var req = NewFakeRequest()

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			result, runErr := vm.Run(program, req, nil)
			if runErr != nil {
				b.Fatal(runErr)
			}
			_ = result
		}
	})
}
