// Copyright 2026 FlexCDN root@flexcdn.cn. All rights reserved. Official site: https://flexcdn.cn .

package flexlang_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tjbrains/flexlang/pkg/flexlang"
)

func TestBasicVM_Compile(t *testing.T) {
	var vm = flexlang.SharedBasicVM()

	program, err := vm.Compile(`1+1`)
	if err != nil {
		t.Fatal(err)
	}

	result, err := vm.Run(program)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestBasicVM_CompileEmpty(t *testing.T) {
	var vm = flexlang.SharedBasicVM()
	program, err := vm.Compile("")
	if err != nil {
		t.Fatal(err)
	}

	result, err := vm.Run(program)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestBasicVM_Eval(t *testing.T) {
	var vm = flexlang.SharedBasicVM()
	result, err := vm.Eval(`1+1`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestBasicVM_String(t *testing.T) {
	var vm = flexlang.SharedBasicVM()

	{
		result, err := vm.Eval(`String.charAt('abc', 1)`)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(result)
	}

	{
		result, err := vm.Eval(`String.fromCharCode(97, 98, 99)`)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(result)
	}

	{
		result, err := vm.Eval(`String.endsWith('abc', 'c')`)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(result)
	}

	{
		result, err := vm.Eval(`String.endsWith('abc', 'c', 3)`)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(result)
	}

	{
		result, err := vm.Eval(`String.endsWith('abc', 'b', 2)`)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(result)
	}
}

func TestBasicVM_String_Concat(t *testing.T) {
	var vm = flexlang.SharedBasicVM()
	result, err := vm.Eval(`'abc' + 'efg'`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestBasicVM_Global_IsNaN(t *testing.T) {
	var vm = flexlang.SharedBasicVM()
	{
		result, err := vm.Eval(`isNaN(1)`)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(result)
		assert.Equal(t, false, result)
	}

	{
		result, err := vm.Eval(`isNaN(NaN)`)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(result)
		assert.Equal(t, true, result)
	}
}

func TestBasicVM_Math(t *testing.T) {
	var vm = flexlang.SharedBasicVM()
	{
		result, err := vm.Eval(`Math.PI`)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(result)
	}
}

func TestBasicVM_JSON(t *testing.T) {
	var vm = flexlang.SharedBasicVM()
	{
		result, err := vm.Eval(`JSON.parse('{ "a": 1, "b": 2 }')`)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(result)
	}

	{
		result, err := vm.Eval(`JSON.parse('1')`)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(result)
	}

	{
		result, err := vm.Eval(`JSON.stringify({"a":1, "b": 2})`)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(result)
	}
}

func TestBasicVM_Date(t *testing.T) {
	var vm = flexlang.SharedBasicVM()
	{
		result, err := vm.Eval(`let d = Date.new();
let a = {"b":d};
String.concat(string(a.b.getFullYear()), "-", String.padStart(string(a.b.getMonth()+1), 2, '0'), "-", string(a.b.getDate()))`)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(result)
	}
}

func TestBasicVM_Regexp(t *testing.T) {
	var vm = flexlang.SharedBasicVM()
	{
		result, err := vm.Eval(`if  "abc" matches "b" {
	1
} else {
	0
}`)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(result)
	}
}

func TestBasicVM_Regexp2(t *testing.T) {
	var vm = flexlang.SharedBasicVM()
	{
		result, err := vm.Eval(`if RegExp.new("\\w").test("abc") {
	1
} else {
	0
}`)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(result)
	}

	{
		result, err := vm.Eval(`if len(RegExp.new("\\w").exec("abc")) > 0 {
	1
} else {
	0
}`)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(result)
	}

	{
		result, err := vm.Eval(`if len(RegExp.new("\\w").exec("==")) > 0 {
	1
} else {
	0
}`)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(result)
	}

	{
		result, err := vm.Eval(`if len(String.match("abc", "\\w+")) > 0 {
	1
} else {
	0
}`)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(result)
	}

	{
		result, err := vm.Eval(`if len(String.match("abc", RegExp.new("\\w+"))) > 0 {
	1
} else {
	0
}`)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(result)
	}

	{
		result, err := vm.Eval(`if len(String.match("abc", nil)) > 0 {
	1
} else {
	0
}`)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(result)
	}
}

func TestBasicVM_New(t *testing.T) {
	var vm = flexlang.SharedBasicVM()

	{
		result, err := vm.Eval(`NewDate().getDate()`)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(result)
	}

	{
		result, err := vm.Eval(`NewRegExp("\\w+").test("abc")`)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(result)
	}

	{
		result, err := vm.Eval(`let u = NewURL("https://example.com/index.html");
u.path`)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(result)
	}
}

func TestBasicVM_Base64(t *testing.T) {
	var vm = flexlang.SharedBasicVM()

	{
		result, err := vm.Eval(`Base64.encode("abc")`)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(result)
	}

	{
		result, err := vm.Eval(`len(Base64.decode(Base64.encode("abc")))`)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(result)
	}

	{
		result, err := vm.Eval(`len(Base64.decode("123"))`)
		if err != nil {
			t.Log(err)
		}
		t.Log(result)
	}
}

func TestBasicVM_Sprintf(t *testing.T) {
	var vm = flexlang.SharedBasicVM()

	{
		result, err := vm.Eval(`String.sprintf("%s, %d", "abc", 1)`)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(result)
	}
}

func TestBasicVM_Overflow(t *testing.T) {
	var vm = flexlang.SharedBasicVM()

	{
		_, err := vm.Eval(`String.split("abc", "b")[3]`)
		assert.NotNil(t, err)
		t.Log(err)
	}

	{
		_, err := vm.Eval(`String.split("abc", "b")[-100]`)
		assert.NotNil(t, err)
		t.Log(err)
	}
}

func TestBasicVM_Hash(t *testing.T) {
	var vm = flexlang.SharedBasicVM()

	{
		result, err := vm.Eval(`md5("123456")`)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(result)
	}
}

func TestBasicVM_CryptoHMAC(t *testing.T) {
	var vm = flexlang.SharedBasicVM()

	{
		result, err := vm.Eval(`let h = Crypto.NewHMAC("sha1", "");
h.update("123456");
h.sum()
`)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(result)
	}
}

func TestBasicVM_IfElse(t *testing.T) {
	var vm = flexlang.SharedBasicVM()

	for _, code := range []string{
		`if true {
 2
}`,
		`if true {
 2
} else if false {
	3
}`,
		`if true {
	// 中文
	2
}`,
	} {
		program, err := vm.Compile(code)
		if err != nil {
			t.Fatal(err)
		}

		result, err := vm.Run(program)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(result)
	}
}

func TestBasicVM_Error(t *testing.T) {
	var vm = flexlang.SharedBasicVM()

	for _, code := range []string{
		`a`,
		`return`,
		`true; return`,
		`if true {
	true
}`,
	} {
		program, err := vm.Compile(code)
		if err == nil {
			_, err = vm.Run(program)
		}
		t.Log(code, "=>", err, "program:", program)
	}
}

func TestBasicVM_LongErr(t *testing.T) {
	var vm = flexlang.SharedBasicVM()
	_, err := vm.Eval(`'a' + String.length1()`)
	if err != nil {
		t.Log(err)
	}
}

func TestBasicVM_ConsoleLog(t *testing.T) {
	var vm = flexlang.SharedBasicVM()
	_, err := vm.Eval(`console.log("Hello", "World")`)
	if err != nil {
		t.Fatal(err)
	}
}

func BenchmarkBasicVM_Run(b *testing.B) {
	var vm = flexlang.NewBasicVM(1 << 10)

	program, err := vm.Compile(`let a = String.charAt('abc', 1); len(a)%2`)
	if err != nil {
		b.Fatal(err)
	}

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, runErr := vm.Run(program)
			if runErr != nil {
				b.Fatal(runErr)
			}
		}
	})
}

func BenchmarkBasicVM_Run_Date(b *testing.B) {
	var vm = flexlang.NewBasicVM(1 << 10)

	program, err := vm.Compile(`let d = NewDate();
let a = {"b":d};
String.concat(string(a.b.getFullYear()), "-", String.padStart(string(a.b.getMonth()+1), 2, '0'), "-", string(a.b.getDate()))`)
	if err != nil {
		b.Fatal(err)
	}

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, runErr := vm.Run(program)
			if runErr != nil {
				b.Fatal(runErr)
			}
		}
	})
}

func BenchmarkBasicVM_Eval_Date(b *testing.B) {
	var vm = flexlang.NewBasicVM(1 << 10)

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, runErr := vm.Eval(`let d = Date.new();
let a = {"b":d};
String.concat(string(a.b.getFullYear()), "-", String.padStart(string(a.b.getMonth()+1), 2, '0'), "-", string(a.b.getDate()))`)
			if runErr != nil {
				b.Fatal(runErr)
			}
		}
	})
}
