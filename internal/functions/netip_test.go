// Copyright 2026 FlexCDN root@flexcdn.cn. All rights reserved. Official site: https://flexcdn.cn .

package functions_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tjbrains/flexlang/internal/functions"
)

func TestNetIPFunctions_IsValid(t *testing.T) {
	var netIPFunctions functions.NetIPFunctions

	assert.True(t, netIPFunctions.IsValid("127.0.0.1"))
	assert.True(t, netIPFunctions.IsValid("::1"))
	assert.False(t, netIPFunctions.IsValid("127.0.0.1.1"))
	assert.False(t, netIPFunctions.IsValid("::1.1"))
	assert.False(t, netIPFunctions.IsValid(""))
	assert.False(t, netIPFunctions.IsValid("127.0.0.256"))
}

func TestNetIPFunctions_IsIPv4(t *testing.T) {
	var netIPFunctions functions.NetIPFunctions

	assert.True(t, netIPFunctions.IsIPv4("127.0.0.1"))
	assert.False(t, netIPFunctions.IsIPv4("::1"))
}

func TestNetIPFunctions_IsIPv6(t *testing.T) {
	var netIPFunctions functions.NetIPFunctions

	assert.False(t, netIPFunctions.IsIPv6("127.0.0.1"))
	assert.True(t, netIPFunctions.IsIPv6("::1"))
}

func TestNetIPFunctions_IsBetween(t *testing.T) {
	var netIPFunctions functions.NetIPFunctions

	assert.True(t, netIPFunctions.IsBetween("127.0.0.1", "127.0.0.1", "127.0.0.2"))
	assert.False(t, netIPFunctions.IsBetween("127.0.0.1", "127.0.0.2", "127.0.0.3"))
	assert.True(t, netIPFunctions.IsBetween("127.0.1.2", "127.0.0.2", "127.0.2.3"))
}

func TestNetIPFunctions_IsInCIDR(t *testing.T) {
	var netIPFunctions functions.NetIPFunctions

	assert.True(t, netIPFunctions.IsInCIDR("127.0.0.1", "127.0.0.0/24"))
	assert.True(t, netIPFunctions.IsInCIDR("127.0.0.1", "127.0.0.0/16"))
	assert.True(t, netIPFunctions.IsInCIDR("127.0.1.2", "127.0.0.0/8"))
	assert.False(t, netIPFunctions.IsInCIDR("127.0.1.1", "127.0.0.0/32"))
}

func TestNetIPFunctions_IsInRanges(t *testing.T) {
	var netIPFunctions functions.NetIPFunctions

	assert.True(t, netIPFunctions.IsInRanges("127.0.0.1", []any{[]any{"127.0.0.0", "127.0.0.255"}}))
	assert.True(t, netIPFunctions.IsInRanges("127.0.0.1", []any{[]any{"127.0.0.0", "127.0.0.127"}}))
	assert.True(t, netIPFunctions.IsInRanges("127.0.1.2", []any{
		[]any{"127.0.0.0", "127.0.0.255"},
		[]any{"127.0.1.0", "127.0.1.255"},
	}))
}
