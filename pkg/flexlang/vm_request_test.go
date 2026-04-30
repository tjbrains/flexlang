// Copyright 2026 FlexCDN root@flexcdn.cn. All rights reserved. Official site: https://flexcdn.cn .

package flexlang_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tjbrains/flexlang/pkg/flexlang"
	"github.com/tjbrains/flexlang/pkg/flexlang/context"
)

func TestRequestVM(t *testing.T) {
	var vm = flexlang.NewRequestVM(1 << 10)
	program, err := vm.Compile("1+1")
	if err != nil {
		t.Fatal(err)
	}

	result, err := vm.Run(program, flexlang.RequestEnv{})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestRequestVM_Empty(t *testing.T) {
	var vm = flexlang.NewRequestVM(1 << 10)
	program, err := vm.Compile("")
	if err != nil {
		t.Fatal(err)
	}

	result, err := vm.Run(program, flexlang.RequestEnv{})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestRequestVM_Req(t *testing.T) {
	var vm = flexlang.NewRequestVM(1 << 10)

	{
		program, err := vm.Compile("$req.serverInfo().id")
		if err != nil {
			t.Fatal(err)
		}

		result, err := vm.Run(program, flexlang.RequestEnv{
			Req: context.NewFakeRequest(),
		})
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%#v", result)
	}

	{
		result, err := vm.Eval("$req.url()", flexlang.RequestEnv{
			Req: context.NewFakeRequest(),
		})
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%#v", result)
	}

	{
		result, err := vm.Eval(`$req.setHeader("Hello", "World", "Universe"); $req.header()`, flexlang.RequestEnv{
			Req: context.NewFakeRequest(),
		})
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%#v", result)
	}

	{
		result, err := vm.Eval(`$req.query().get("name")`, flexlang.RequestEnv{
			Req: context.NewFakeRequest(),
		})
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%#v", result)
		assert.Equal(t, "Lily", result)
	}

	{
		result, err := vm.Eval(`$req.query().values("name")`, flexlang.RequestEnv{
			Req: context.NewFakeRequest(),
		})
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%#v", result)
	}
}

func TestRequestVM_Query(t *testing.T) {
	var vm = flexlang.NewRequestVM(1 << 10)

	{
		result, err := vm.Eval(`
let query = $req.query();
query.get("name")`, flexlang.RequestEnv{
			Req: context.NewFakeRequest(),
		})
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%#v", result)
	}

	{
		result, err := vm.Eval(`
let query = $req.query().toKV();
query.name + ' ' + query['name']`, flexlang.RequestEnv{
			Req: context.NewFakeRequest(),
		})
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%#v", result)
	}

	{
		result, err := vm.Eval(`
let jsonData = $req.query().toJSON();
jsonData`, flexlang.RequestEnv{
			Req: context.NewFakeRequest(),
		})
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%#v", result)
	}

	{
		result, err := vm.Eval(`
let jsonData = $req.query().toKVJSON();
jsonData`, flexlang.RequestEnv{
			Req: context.NewFakeRequest(),
		})
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

	result, err := vm.Run(program, flexlang.RequestEnv{})
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

	result, err := vm.Run(program, flexlang.RequestEnv{})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestRequestVM_ConsoleLog(t *testing.T) {
	var vm = flexlang.NewRequestVM(1 << 10)
	_, err := vm.Eval(`console.log("Hello")`, flexlang.RequestEnv{
		Printer: func(s ...string) {
			t.Log(strings.Join(s, " "))
		},
	})
	if err != nil {
		t.Fatal(err)
	}
}

func BenchmarkRequestVM_Req(b *testing.B) {
	var vm = flexlang.NewRequestVM(1 << 10)
	program, err := vm.Compile(`$req.path()`)
	if err != nil {
		b.Fatal(err)
	}

	var req = context.NewFakeRequest()

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			result, runErr := vm.Run(program, flexlang.RequestEnv{
				Req: req,
			})
			if runErr != nil {
				b.Fatal(runErr)
			}
			_ = result
		}
	})
}
