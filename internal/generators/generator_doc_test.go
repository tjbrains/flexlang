// Copyright 2026 FlexCDN root@flexcdn.cn. All rights reserved. Official site: https://flexcdn.cn .

package generators_test

import (
	"testing"

	"github.com/tjbrains/flexlang/internal/generators"
)

func TestDocGenerator_Run(t *testing.T) {
	var generator = generators.NewDocGenerator()
	err := generator.Run()
	if err != nil {
		t.Fatal(err)
	}
}
