// Copyright 2026 FlexCDN root@flexcdn.cn. All rights reserved. Official site: https://flexcdn.cn .

package functions_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tjbrains/flexlang/internal/functions"
)

func TestURL(t *testing.T) {
	var u = functions.NewURL("https://user:pass@example.com:8080/docs?nav=1#link")

	t.Logf("%#v", u)

	assert.Equal(t, 8080, u.Port)
	assert.Equal(t, "link", u.Hash)
	assert.Equal(t, "example.com:8080", u.Host)
	assert.Equal(t, "nav=1", u.Query)
	assert.Equal(t, "https", u.Scheme)
	assert.Equal(t, "/docs", u.Path)
	assert.Equal(t, "", u.Opaque)
	assert.True(t, len(u.User) > 0)
}
