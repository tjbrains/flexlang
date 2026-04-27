// Copyright 2026 FlexCDN root@flexcdn.cn. All rights reserved. Official site: https://flexcdn.cn .

package context_test

import (
	"net/url"
	"testing"

	"github.com/tjbrains/flexlang/pkg/flexlang"
)

func TestRequestQuery_ToJSON(t *testing.T) {
	var query = flexlang.NewURLQuery(url.Values{})
	query.Set("name", "Lily")
	query.Set("v", "1")
	t.Log(query.Encode())
	t.Log(query.ToJSON())
	t.Log(query.ToMap())
}