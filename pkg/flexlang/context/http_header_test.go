// Copyright 2026 FlexCDN root@flexcdn.cn. All rights reserved. Official site: https://flexcdn.cn .

package context_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tjbrains/flexlang/pkg/flexlang/context"
)

func TestHTTPHeader(t *testing.T) {
	var header = context.HTTPHeader{}
	header.Set("Content-Type", "text/html; charset=utf-8")

	t.Log(header)

	assert.True(t, header.Has("Content-Type"))

	header.Delete("Content-Type")

	assert.False(t, header.Has("Content-Type"))

	header.Add("Cookie", "a=b")
	header.Add("Cookie", "c=d")
	assert.Equal(t, 2, len(header.Values("Cookie")))

	t.Log(header.Values("Cookie"))

	t.Log("JSON:", header.ToJSON())
}
