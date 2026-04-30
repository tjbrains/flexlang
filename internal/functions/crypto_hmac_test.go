// Copyright 2026 FlexCDN root@flexcdn.cn. All rights reserved. Official site: https://flexcdn.cn .

package functions_test

import (
	"testing"

	"github.com/tjbrains/flexlang/internal/functions"
)

func TestCryptoHMAC(t *testing.T) {
	h, err := functions.NewCryptoHMAC("sha1", "123456")
	if err != nil {
		t.Fatal(err)
	}
	h.Update("ABCDEFG")
	t.Log(h.Sum())
}
