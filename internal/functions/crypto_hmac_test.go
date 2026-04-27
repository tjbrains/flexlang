// Copyright 2026 FlexCDN root@flexcdn.cn. All rights reserved. Official site: https://flexcdn.cn .

package functions_test

import (
	"testing"

	"github.com/tjbrains/flexlang/internal/functions"
)

func TestCryptoHMAC(t *testing.T) {
	h, err := functions.NewCryptoHMAC("sha1", "")
	if err != nil {
		t.Fatal(err)
	}
	h.Update("123456")
	t.Log(h.Sum())
}
