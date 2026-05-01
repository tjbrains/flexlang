# FlexLang

An expression language based on ExprLang.

~~~go
var vm = flexlang.SharedBasicVM()
program, err := vm.Compile(`Math.randN(10)`)
if err != nil {
    log.Fatal(err)
}

result, err := vm.Run(program)
if err != nil {
    log.Fatal(err)
}
log.Println(result)
~~~