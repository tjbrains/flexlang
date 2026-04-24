// Copyright 2026 FlexCDN root@flexcdn.cn. All rights reserved. Official site: https://flexcdn.cn .

package functions_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tjbrains/flexlang/pkg/flexlang/functions"
)

func TestRegExpFunctions_Escape(t *testing.T) {
	var regexpFunctions functions.RegExpFunctions
	t.Log(regexpFunctions.Escape("#$%*()-[]{}|"))
}

func TestRegExpFunctions_New(t *testing.T) {
	var regexpFunctions functions.RegExpFunctions

	{
		var reg = regexpFunctions.New("\\w+")
		assert.True(t, reg.Test("123abc"))
	}

	{
		var reg = regexpFunctions.New("+++")
		assert.False(t, reg.Test("123abc"))
	}

	{
		var reg = regexpFunctions.New("^\\w+$")
		assert.True(t, reg.Test("123abc"))
	}

	{
		var reg = regexpFunctions.New("^\\w+")
		assert.False(t, reg.Test("#123abc"))
	}

	{
		var reg = regexpFunctions.New("\\w+$")
		assert.False(t, reg.Test("123abc#"))
	}
}

func TestRegExpFunctions_Exec(t *testing.T) {
	var regexpFunctions functions.RegExpFunctions

	t.Log(regexpFunctions.New(`\d{3}`).Exec("123456"))
	t.Log(regexpFunctions.New(`\d+`).Exec("123456"))
	t.Log(regexpFunctions.New(`\d+`).Exec("123|456"))
	t.Log(regexpFunctions.New(`(\d)(\d)`).Exec("123|456"))
}
