// Copyright 2026 FlexCDN root@flexcdn.cn. All rights reserved. Official site: https://flexcdn.cn .

package functions_test

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tjbrains/flexlang/internal/functions"
)

func TestGlobalFunctions_IsNaN(t *testing.T) {
	var globalFunctions functions.GlobalFunctions

	assert.True(t, globalFunctions.IsNaN(functions.NaN))
	assert.False(t, globalFunctions.IsNaN(1))
}

func TestGlobalFunctions_ParseFloat(t *testing.T) {
	var globalFunctions functions.GlobalFunctions

	assert.Equal(t, 0.0, globalFunctions.ParseFloat("0"))
	assert.Equal(t, 0.0, globalFunctions.ParseFloat("abc"))
	assert.Equal(t, 123.0, globalFunctions.ParseFloat("123"))
	assert.Equal(t, 123.456, globalFunctions.ParseFloat("123.456"))
	assert.NotEqual(t, 123.45, globalFunctions.ParseFloat("123.456"))
}

func TestGlobalFunctions_ParseInt(t *testing.T) {
	var globalFunctions functions.GlobalFunctions

	var a = assert.New(t)

	a.Equal(int64(0), globalFunctions.ParseInt(""))
	a.Equal(int64(0), globalFunctions.ParseInt("abc"))
	a.Equal(int64(0), globalFunctions.ParseInt("0"))
	a.Equal(int64(123), globalFunctions.ParseInt("123"))
	a.Equal(int64(123), globalFunctions.ParseInt("123.456"))
	a.Equal(int64(1), globalFunctions.ParseInt(int(1)))
	a.Equal(int64(12345), globalFunctions.ParseInt(uint16(12345)))
	a.Equal(int64(12345), globalFunctions.ParseInt(float32(12345.456)))
	a.Equal(int64(12345), globalFunctions.ParseInt(float64(12345.456)))
	a.Equal(int64(1), globalFunctions.ParseInt(true))
	a.Equal(int64(0), globalFunctions.ParseInt(false))
}

func TestGlobalFunctions_DecodeURIComponent(t *testing.T) {
	var globalFunctions functions.GlobalFunctions

	assert.Equal(t, "", globalFunctions.DecodeURIComponent(""))
	assert.Equal(t, "%", globalFunctions.DecodeURIComponent(`%25`))
	assert.Equal(t, "=", globalFunctions.DecodeURIComponent(`%3D`))
}

func TestGlobalFunctions_EncodeURIComponent(t *testing.T) {
	var globalFunctions functions.GlobalFunctions

	assert.Equal(t, "", globalFunctions.EncodeURIComponent(""))
	assert.Equal(t, "%25", globalFunctions.EncodeURIComponent(`%`))
	assert.Equal(t, "%3D", globalFunctions.EncodeURIComponent(`=`))
}

func TestGlobalFunctions_TypeOf(t *testing.T) {
	var globalFunctions functions.GlobalFunctions

	assert.Equal(t, "string", globalFunctions.TypeOf(""))
	assert.Equal(t, "string", globalFunctions.TypeOf("abc"))
	assert.Equal(t, "number", globalFunctions.TypeOf(123))
	assert.Equal(t, "number", globalFunctions.TypeOf(123.0))
	assert.Equal(t, "boolean", globalFunctions.TypeOf(true))
	assert.Equal(t, "boolean", globalFunctions.TypeOf(false))
	assert.Equal(t, "bigint", globalFunctions.TypeOf(big.NewInt(1234)))
}

func TestGlobalFunctions_Hash(t *testing.T) {
	var globalFunctions functions.GlobalFunctions

	t.Log(globalFunctions.MD5(""))
	t.Log(globalFunctions.MD5("123456"))
	t.Log(globalFunctions.Sha1("123456"))
	t.Log(globalFunctions.Sha256("123456"))
	t.Log(globalFunctions.Crc32("123456"))
	t.Log(globalFunctions.Crc32("1234567"))
}
