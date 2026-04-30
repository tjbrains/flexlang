// Copyright 2026 FlexCDN root@flexcdn.cn. All rights reserved. Official site: https://flexcdn.cn .

package main

import (
	"fmt"

	"github.com/tjbrains/flexlang/internal/generators"
)

func main() {
	var generator = generators.NewDocGenerator()
	err := generator.Run()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("success")
}
