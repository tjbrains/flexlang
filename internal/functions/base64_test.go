// Copyright 2026 FlexCDN root@flexcdn.cn. All rights reserved. Official site: https://flexcdn.cn .

package functions_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tjbrains/flexlang/internal/functions"
)

func TestBase64Functions_Encode(t *testing.T) {
	var base64Functions functions.Base64Functions

	var s = base64Functions.Encode("Hello, World!")
	t.Log(s)

	assert.Equal(t, "Hello, World!", base64Functions.Decode(s))
}
