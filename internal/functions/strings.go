// Copyright 2026 FlexCDN root@flexcdn.cn. All rights reserved. Official site: https://flexcdn.cn .

package functions

import (
	"fmt"
	goregexp "regexp"
	"strings"
	"unicode"
)

// 已忽略的函数：
//  - String.fromCodePoint(...codePoints)
//  - String.raw(template, ...substitutions)
//	- codePointAt(pos)
//  - localeCompare(that [, reserved1 [ , reserved2 ] ] ))
//  - matchAll( regexp )
//  - normalize([ form ])
//  - isWellFormed()
//  - search ( regexp )
//  - toLocaleLowerCase ( [ reserved1 [ , reserved2 ] ] )
//  - toLocaleUpperCase ( [ reserved1 [ , reserved2 ] ] )
//  - toWellFormed ( )

type StringFunctions struct {
}

func (this StringFunctions) FromCharCode(codeUnits ...rune) string {
	return string(codeUnits)
}

func (this StringFunctions) At(s string, index int) string {
	if len(s) == 0 {
		return ""
	}

	if index < 0 {
		index += len(s)
	}

	if index >= 0 && index < len(s) {
		return s[index : index+1]
	}

	return ""
}

func (this StringFunctions) CharAt(s string, pos int) string {
	if pos < 0 || pos >= len(s) {
		return ""
	}
	return s[pos : pos+1]
}

func (this StringFunctions) CharCodeAt(s string, pos int) any {
	var runes = []rune(s)

	if pos < 0 || pos >= len(runes) {
		return NaN
	}

	return int(runes[pos])
}

func (this StringFunctions) Concat(s string, args ...string) string {
	var n = len(s)
	for _, arg := range args {
		n += len(arg)
	}

	if n == 0 {
		return ""
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteString(s)

	for _, arg := range args {
		b.WriteString(arg)
	}

	return b.String()
}

func (this StringFunctions) EndsWith(s string, searchString string, endPosition ...int) bool {
	if len(searchString) == 0 {
		return true
	}

	if len(s) == 0 {
		return false
	}

	if len(endPosition) == 0 {
		return strings.HasSuffix(s, searchString)
	}

	var position = endPosition[0]

	if position < 0 {
		return false
	}

	return strings.HasSuffix(s[:min(len(s), position)], searchString)
}

func (this StringFunctions) Includes(s string, searchString string, endPosition ...int) bool {
	if len(searchString) == 0 {
		return true
	}

	if len(s) == 0 {
		return false
	}

	if len(endPosition) == 0 {
		return strings.Contains(s, searchString)
	}

	var position = endPosition[0]
	if position < 0 || position >= len(s) {
		return false
	}

	return strings.Contains(s[:position], searchString)
}

func (this StringFunctions) IndexOf(s string, searchString string, startPosition ...int) int {
	if len(searchString) == 0 {
		if len(startPosition) == 0 {
			return 0
		}

		var position = startPosition[0]
		if position < 0 {
			return 0
		}

		return min(position, len(s))
	}

	if len(s) == 0 {
		return -1
	}

	if len(startPosition) == 0 {
		return strings.Index(s, searchString)
	}

	var position = startPosition[0]
	if position < 0 || position >= len(s) {
		return -1
	}

	var index = strings.Index(s[position:], searchString)
	if index < 0 {
		return -1
	}

	return index + position
}

func (this StringFunctions) LastIndexOf(s string, searchString string, endPosition ...int) int {
	if len(searchString) == 0 {
		if len(endPosition) == 0 {
			return len(s)
		}

		var position = endPosition[0]

		if position < 0 {
			return -1
		}

		return min(position, len(s))
	}

	if len(endPosition) == 0 {
		return strings.LastIndex(s, searchString)
	}

	var position = endPosition[0]
	if position < 0 {
		return -1
	}

	var index = strings.LastIndex(s, searchString)
	if index <= position {
		return index
	}

	return -1
}

func (this StringFunctions) Match(s string, regexp any) []string {
	if regexp == nil {
		return nil
	}
	switch reg := regexp.(type) {
	case string:
		r, err := goregexp.Compile(reg)
		if err != nil {
			return nil
		}
		return r.FindStringSubmatch(s)
	case RegExp:
		return reg.Exec(s)
	}

	return nil
}

// Repeat
// 注意：count 最大值为1M
// @DANGER
func (this StringFunctions) Repeat(s string, count int) string {
	if count <= 0 {
		return ""
	}

	const maxRepeats = 1 << 20
	return strings.Repeat(s, min(count, maxRepeats))
}

func (this StringFunctions) Replace(s string, searchValue string, replaceValue string) string {
	return strings.Replace(s, searchValue, replaceValue, 1)
}

func (this StringFunctions) ReplaceAll(s string, searchValue string, replaceValue string) string {
	return strings.ReplaceAll(s, searchValue, replaceValue)
}

// PadEnd
// @DANGER
func (this StringFunctions) PadEnd(s string, maxLength int, fillString string) string {
	if maxLength <= 0 {
		return s
	}

	var l = len(s)
	if maxLength <= l {
		return s
	}

	if len(fillString) == 0 {
		return s
	}

	var result = s
	var repeats = (maxLength - l) / len(fillString)
	if repeats > 0 {
		result += strings.Repeat(fillString, repeats)
	}
	if len(result) < maxLength {
		result += fillString[:maxLength-len(result)]
	}

	return result
}

// PadStart
// @DANGER
func (this StringFunctions) PadStart(s string, maxLength int, fillString string) string {
	if maxLength <= 0 {
		return s
	}

	var l = len(s)
	if maxLength <= l {
		return s
	}

	if len(fillString) == 0 {
		return s
	}

	var result string
	var repeats = (maxLength - l) / len(fillString)
	if repeats > 0 {
		result = strings.Repeat(fillString, repeats)
	}
	if len(result)+l < maxLength {
		result += fillString[:maxLength-len(result)-l]
	}

	result += s

	return result
}

func (this StringFunctions) Slice(s string, start int, end int) string {
	var l = len(s)
	if l == 0 {
		return ""
	}

	if start < 0 {
		start += l
	}

	if end < 0 {
		end += l
	}

	if start < 0 || end < 0 || start >= l || start >= end {
		return ""
	}

	return s[start:min(l, end)]
}

func (this StringFunctions) Split(s string, separator string, limit ...int) []string {
	if len(limit) == 0 {
		if len(s) == 0 {
			if len(separator) == 0 {
				return []string{}
			}
			return []string{""}
		}

		return strings.Split(s, separator)
	}

	var n = limit[0]
	if n == 0 {
		return []string{}
	}

	if len(s) == 0 {
		if len(separator) == 0 {
			return []string{}
		}
		return []string{""}
	}

	if n < 0 {
		return strings.Split(s, separator)
	}

	var result = strings.Split(s, separator)
	return result[:min(n, len(result))]
}

func (this StringFunctions) StartsWith(s string, searchString string, startPosition ...int) bool {
	if len(searchString) == 0 {
		return true
	}

	if len(s) == 0 {
		return false
	}

	if len(startPosition) == 0 {
		return strings.HasPrefix(s, searchString)
	}

	var position = startPosition[0]

	if position < 0 {
		position = 0
	}

	if position >= len(s) {
		return false
	}

	return strings.HasPrefix(s[position:], searchString)
}

func (this StringFunctions) Substring(s string, start int, end int) string {
	var l = len(s)
	if l == 0 {
		return ""
	}

	if start < 0 {
		start = 0
	}

	if end < 0 {
		return ""
	}

	if start == end {
		return ""
	}

	if start > end {
		start, end = end, start
	}

	if start > l {
		return ""
	}

	if end > l {
		end = l
	}

	return s[start:end]
}

func (this StringFunctions) ToLowerCase(s string) string {
	if len(s) == 0 {
		return ""
	}

	return strings.ToLower(s)
}

func (this StringFunctions) ToUpperCase(s string) string {
	if len(s) == 0 {
		return ""
	}

	return strings.ToUpper(s)
}

func (this StringFunctions) Trim(s string) string {
	return strings.TrimSpace(s)
}

func (this StringFunctions) TrimEnd(s string) string {
	return strings.TrimRightFunc(s, unicode.IsSpace)
}

func (this StringFunctions) TrimPrefix(s string, prefix string) string {
	return strings.TrimPrefix(s, prefix)
}

func (this StringFunctions) TrimStart(s string) string {
	return strings.TrimLeftFunc(s, unicode.IsSpace)
}

func (this StringFunctions) TrimSuffix(s string, suffix string) string {
	return strings.TrimSuffix(s, suffix)
}

func (this StringFunctions) Length(s string) int {
	return len([]rune(s))
}

func (this StringFunctions) Sprintf(s string, args ...any) string {
	return fmt.Sprintf(s, args...)
}
