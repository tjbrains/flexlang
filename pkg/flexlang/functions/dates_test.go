// Copyright 2026 FlexCDN root@flexcdn.cn. All rights reserved. Official site: https://flexcdn.cn .

package functions_test

import (
	"testing"

	"github.com/tjbrains/flexlang/pkg/flexlang/functions"
)

func TestDateFunctions_New(t *testing.T) {
	var dateFunctions functions.DateFunctions

	var date = dateFunctions.New()
	t.Log(date.GetDate(), date.GetDay(), date.GetFullYear(), date.GetHours(), date.GetMinutes(), date.GetSeconds(), date.GetMilliseconds(), date.GetMonth(), date.GetTime())
}
