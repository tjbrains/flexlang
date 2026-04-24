// Copyright 2026 FlexCDN root@flexcdn.cn. All rights reserved. Official site: https://flexcdn.cn .

package functions_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tjbrains/flexlang/pkg/flexlang/functions"
)

func TestStringFunctions_FromCharCode(t *testing.T) {
	var f functions.StringFunctions
	assert.Equal(t, "", f.FromCharCode())
	assert.Equal(t, "abc", f.FromCharCode('a', 'b', 'c'))
}

func TestStringFunctions_At(t *testing.T) {
	var f functions.StringFunctions

	assert.Equal(t, "a", f.At("abc", 0))
	assert.Equal(t, "b", f.At("abc", 1))
	assert.Equal(t, "c", f.At("abc", 2))
	assert.Equal(t, "", f.At("abc", 3))
	assert.Equal(t, "c", f.At("abc", -1))
	assert.Equal(t, "b", f.At("abc", -2))
	assert.Equal(t, "a", f.At("abc", -3))
	assert.Equal(t, "", f.At("abc", -4))
}

func TestStringFunctions_CharAt(t *testing.T) {
	var f functions.StringFunctions

	assert.Equal(t, "a", f.CharAt("abc", 0))
	assert.Equal(t, "b", f.CharAt("abc", 1))
	assert.Equal(t, "c", f.CharAt("abc", 2))

	assert.Equal(t, "", f.CharAt("abc", -11))
	assert.Equal(t, "", f.CharAt("abc", 3))
}

func TestStringFunctions_CharCodeAt(t *testing.T) {
	var f functions.StringFunctions

	var gf functions.GlobalFunctions

	assert.Equal(t, 97, f.CharCodeAt("abc", 0))
	assert.Equal(t, 98, f.CharCodeAt("abc", 1))
	assert.Equal(t, 99, f.CharCodeAt("abc", 2))
	assert.True(t, gf.IsNaN(f.CharCodeAt("abc", -1)))
	assert.True(t, gf.IsNaN(f.CharCodeAt("abc", 3)))
}

func TestStringFunctions_Concat(t *testing.T) {
	var f functions.StringFunctions

	assert.Equal(t, "", f.Concat(""))
	assert.Equal(t, "abc", f.Concat("", "a", "b", "c"))
	assert.Equal(t, "ABC", f.Concat("ABC"))
	assert.Equal(t, "ABCabc", f.Concat("ABC", "a", "b", "c"))
}

func TestStringFunctions_EndsWith(t *testing.T) {
	var f functions.StringFunctions

	assert.True(t, f.EndsWith("abc", "c"))
	assert.False(t, f.EndsWith("abc", "b"))

	assert.False(t, f.EndsWith("abc", "c", -1))
	assert.True(t, f.EndsWith("abc", "c", 3))
	assert.True(t, f.EndsWith("abc", "c", 4))
	assert.True(t, f.EndsWith("abc", "b", 2))
	assert.True(t, f.EndsWith("abc", "a", 1))
	assert.True(t, f.EndsWith("abc", "", 0))
}

func TestStringFunctions_Includes(t *testing.T) {
	var f functions.StringFunctions

	assert.True(t, f.Includes("abcdefg", "cde"))
	assert.False(t, f.Includes("abcdefg", "bde"))
	assert.True(t, f.Includes("abcdefg", "cde", 6))
	assert.False(t, f.Includes("abcdefg", "cde", 3))
}

func TestStringFunctions_IndexOf(t *testing.T) {
	var f functions.StringFunctions

	assert.Equal(t, -1, f.IndexOf("abc", "defg"))
	assert.Equal(t, 0, f.IndexOf("abcdefg", "abc"))
	assert.Equal(t, 1, f.IndexOf("abcdefg", "bcd"))
	assert.Equal(t, 0, f.IndexOf("abcdefg", ""))
	assert.Equal(t, 1, f.IndexOf("abcdefg", "", 1))
	assert.Equal(t, 7, f.IndexOf("abcdefg", "", 10))
	assert.Equal(t, 2, f.IndexOf("abcdefg", "cd", 1))
	assert.Equal(t, -1, f.IndexOf("abcdefg", "cd", 3))
	assert.Equal(t, -1, f.IndexOf("abcdefg", "cd", 10))
}

func TestStringFunctions_Match(t *testing.T) {
	var f functions.StringFunctions
	assert.Equal(t, []string{"abc"}, f.Match("abc", "\\w+"))
	assert.Equal(t, []string{"ab"}, f.Match("ab|c", "\\w+"))
	assert.Equal(t, ([]string)(nil), f.Match("()*&)", "\\w+"))

	var fr functions.RegExpFunctions
	assert.Equal(t, []string{"abc"}, f.Match("abc", fr.New("\\w+")))
	assert.Equal(t, []string{"ab"}, f.Match("ab|c", fr.New("\\w+")))
}

func TestStringFunctions_LastIndexOf(t *testing.T) {
	var f functions.StringFunctions

	assert.Equal(t, -1, f.LastIndexOf("abcdefg", "ff"))
	assert.Equal(t, 0, f.LastIndexOf("abcdefg", "abc"))
	assert.Equal(t, 5, f.LastIndexOf("abcdefg", "fg"))
	assert.Equal(t, 5, f.LastIndexOf("abcdefg", "fg", 5))
	assert.Equal(t, 5, f.LastIndexOf("abcdefg", "fg", 6))
	assert.Equal(t, -1, f.LastIndexOf("abcdefg", "fg", 4))
	assert.Equal(t, 5, f.LastIndexOf("abcdefg", "fg", 10))

	assert.Equal(t, 7, f.LastIndexOf("abcdefg", ""))
	assert.Equal(t, 3, f.LastIndexOf("abcdefg", "", 3))
	assert.Equal(t, 7, f.LastIndexOf("abcdefg", "", 10))
	assert.Equal(t, -1, f.LastIndexOf("abcdefg", "", -1))
}

func TestStringFunctions_Repeat(t *testing.T) {
	var f functions.StringFunctions

	assert.Equal(t, "", f.Repeat("abc", 0))
	assert.Equal(t, "", f.Repeat("abc", -1))
	assert.Equal(t, "abcabcabcabcabcabcabcabcabcabc", f.Repeat("abc", 10))
}

func TestStringFunctions_Replace(t *testing.T) {
	var f functions.StringFunctions

	assert.Equal(t, "abcdefg", f.Replace("abcdefg", "", ""))
	assert.Equal(t, "1bcdefg", f.Replace("abcdefg", "a", "1"))
	assert.Equal(t, "a1cdefg", f.Replace("abcdefg", "b", "1"))
	assert.Equal(t, "a1cabc", f.Replace("abcabc", "b", "1"))

	assert.Equal(t, "a1ca1c", f.ReplaceAll("abcabc", "b", "1"))
}

func TestStringFunctions_PadEnd(t *testing.T) {
	var f functions.StringFunctions

	assert.Equal(t, "abc", f.PadEnd("abc", -1, "ABC"))
	assert.Equal(t, "abc", f.PadEnd("abc", 0, "ABC"))
	assert.Equal(t, "abc", f.PadEnd("abc", 3, "ABC"))
	assert.Equal(t, "abcA", f.PadEnd("abc", 4, "ABC"))
	assert.Equal(t, "abcABC", f.PadEnd("abc", 6, "ABC"))
	assert.Equal(t, "abcABCABC", f.PadEnd("abc", 9, "ABC"))
	assert.Equal(t, "abcABCABCABCAB", f.PadEnd("abc", 14, "ABC"))
	assert.Equal(t, "abc", f.PadEnd("abc", 14, ""))
	assert.Equal(t, "abc           ", f.PadEnd("abc", 14, " "))
}

func TestStringFunctions_PadStart(t *testing.T) {
	var f functions.StringFunctions

	assert.Equal(t, "abc", f.PadStart("abc", -1, "ABC"))
	assert.Equal(t, "abc", f.PadStart("abc", 0, "ABC"))
	assert.Equal(t, "abc", f.PadStart("abc", 3, "ABC"))
	assert.Equal(t, "Aabc", f.PadStart("abc", 4, "ABC"))
	assert.Equal(t, "ABCabc", f.PadStart("abc", 6, "ABC"))
	assert.Equal(t, "ABCABCabc", f.PadStart("abc", 9, "ABC"))
	assert.Equal(t, "ABCABCABCABabc", f.PadStart("abc", 14, "ABC"))
	assert.Equal(t, "abc", f.PadStart("abc", 14, ""))
	assert.Equal(t, "           abc", f.PadStart("abc", 14, " "))
}

func TestStringFunctions_Slice(t *testing.T) {
	var f functions.StringFunctions

	assert.Equal(t, "", f.Slice("", 0, 0))

	assert.Equal(t, "a", f.Slice("abc", -3, -2))
	assert.Equal(t, "ab", f.Slice("abc", -3, -1))
	assert.Equal(t, "abc", f.Slice("abc", 0, 10))
	assert.Equal(t, "bc", f.Slice("abc", 1, 3))
	assert.Equal(t, "", f.Slice("abc", 1, 1))
	assert.Equal(t, "b", f.Slice("abc", 1, 2))
	assert.Equal(t, "efg", f.Slice("abcdefg", -3, 7))
	assert.Equal(t, "", f.Slice("abcdefg", 10, 100))
}

func TestStringFunctions_Split(t *testing.T) {
	var f functions.StringFunctions

	assert.Equal(t, []string{}, f.Split("", ""))
	assert.Equal(t, []string{"a", "b", "c"}, f.Split("abc", ""))
	assert.Equal(t, []string{"a", "c"}, f.Split("abc", "b"))
	assert.Equal(t, []string{"a", "c"}, f.Split("abc", "b", -1))
	assert.Equal(t, []string{"a"}, f.Split("abc", "b", 1))
	assert.Equal(t, []string{"a", "c"}, f.Split("abc", "b", 2))
	assert.Equal(t, []string{"abc"}, f.Split("abc", "d", 2))
}

func TestStringFunctions_StartsWith(t *testing.T) {
	var f functions.StringFunctions

	assert.True(t, f.StartsWith("abc", "a"))
	assert.True(t, f.StartsWith("abc", "ab"))
	assert.False(t, f.StartsWith("abc", "b"))

	assert.True(t, f.StartsWith("abc", "a", -1))
	assert.True(t, f.StartsWith("abc", "a", 0))
	assert.True(t, f.StartsWith("abc", "b", 1))
	assert.True(t, f.StartsWith("abcdefg", "ef", 4))
	assert.True(t, f.StartsWith("abc", "", 1))
	assert.True(t, f.StartsWith("abc", "", 0))
	assert.True(t, f.StartsWith("abc", "", -1))
}

func TestStringFunctions_Substring(t *testing.T) {
	var f functions.StringFunctions

	assert.Equal(t, "", f.Substring("", 1, 2))
	assert.Equal(t, "bc", f.Substring("abcdefg", 1, 3))
	assert.Equal(t, "abc", f.Substring("abcdefg", 0, 3))
	assert.Equal(t, "abc", f.Substring("abcdefg", -1, 3))
	assert.Equal(t, "abc", f.Substring("abcdefg", 3, 0))
	assert.Equal(t, "bc", f.Substring("abcdefg", 3, 1))
	assert.Equal(t, "abcdefg", f.Substring("abcdefg", 0, 100))
	assert.Equal(t, "g", f.Substring("abcdefg", 6, 100))
	assert.Equal(t, "", f.Substring("abcdefg", 7, 100))
	assert.Equal(t, "", f.Substring("abcdefg", 10, 100))
}

func TestStringFunctions_ToLowerCase(t *testing.T) {
	var f functions.StringFunctions

	assert.Equal(t, "", f.ToLowerCase(""))
	assert.Equal(t, "abc", f.ToLowerCase("abc"))
	assert.Equal(t, "abc", f.ToLowerCase("ABC"))
	assert.Equal(t, "abc", f.ToLowerCase("aBC"))
	assert.Equal(t, "abc中文", f.ToLowerCase("ABC中文"))
}

func TestStringFunctions_ToUpperCase(t *testing.T) {
	var f functions.StringFunctions

	assert.Equal(t, "", f.ToUpperCase(""))
	assert.Equal(t, "ABC", f.ToUpperCase("ABC"))
	assert.Equal(t, "ABC", f.ToUpperCase("abc"))
	assert.Equal(t, "ABC", f.ToUpperCase("Abc"))
	assert.Equal(t, "ABC中文", f.ToUpperCase("abc中文"))
}

func TestStringFunctions_Trim(t *testing.T) {
	var f functions.StringFunctions

	assert.Equal(t, "", f.Trim(""))
	assert.Equal(t, "a", f.Trim(" a "))
	assert.Equal(t, "a b", f.Trim(" a b \t\n\r\t"))
}

func TestStringFunctions_TrimEnd(t *testing.T) {
	var f functions.StringFunctions

	assert.Equal(t, "", f.TrimEnd(""))
	assert.Equal(t, " a", f.TrimEnd(" a "))
	assert.Equal(t, " a b", f.TrimEnd(" a b \t\n\r\t"))
}

func TestStringFunctions_TrimPrefix(t *testing.T) {
	var f functions.StringFunctions

	assert.Equal(t, "cdefg", f.TrimPrefix("abcdefg", "ab"))
	assert.Equal(t, "abcdefg", f.TrimPrefix("abcdefg", ""))
}

func TestStringFunctions_TrimStart(t *testing.T) {
	var f functions.StringFunctions

	assert.Equal(t, "", f.TrimStart(""))
	assert.Equal(t, "a ", f.TrimStart(" a "))
	assert.Equal(t, "a b", f.TrimStart(" \t\n\r\t a b"))
}

func TestStringFunctions_TrimSuffix(t *testing.T) {
	var f functions.StringFunctions

	assert.Equal(t, "abcde", f.TrimSuffix("abcdefg", "fg"))
	assert.Equal(t, "abcdefg", f.TrimSuffix("abcdefg", ""))
}

func TestStringFunctions_Length(t *testing.T) {
	var f functions.StringFunctions

	assert.Equal(t, 0, f.Length(""))
	assert.Equal(t, 3, f.Length("abc"))
}

func TestStringFunctions_Sprintf(t *testing.T) {
	var f functions.StringFunctions

	t.Log(f.Sprintf("abc %d, %.2f", 1, 2.3456))
}
