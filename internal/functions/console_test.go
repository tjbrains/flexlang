// Copyright 2026 FlexCDN root@flexcdn.cn. All rights reserved. Official site: https://flexcdn.cn .

package functions_test

import (
	"math"
	"testing"

	"github.com/tjbrains/flexlang/internal/functions"
)

func TestConsoleLog(t *testing.T) {
	var f functions.ConsoleFunctions
	f.Log(func(s ...string) {
		t.Log(s)
	}, true, false, 1, 2, 1.23, math.Pi, []byte("Hello"), map[string]string{"a": "b"})
}
