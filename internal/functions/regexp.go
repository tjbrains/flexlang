// Copyright 2026 FlexCDN root@flexcdn.cn. All rights reserved. Official site: https://flexcdn.cn .

package functions

import (
	"regexp"
)

type RegExpFunctions struct {
}

// Escape 转义特殊字符
//
// @internal
func (this RegExpFunctions) Escape(s string) string {
	return regexp.QuoteMeta(s)
}

// New 创建新正则表达式对象
//
// @internal
func (this RegExpFunctions) New(expr string) RegExp {
	return NewRegExp(expr)
}

// RegExp 正则表达式定义
type RegExp struct {
	// 测试某个字符串是否匹配当前正则表达式
	// 示例：
	// ~~~javascript
	// NewRegExp("\\w+").test("abc") // => true
	// ~~~
	Test func(s string) bool `expr:"test"`

	// 执行当前正则表达式匹配
	// 返回匹配的内容，圆括号视为匹配的子串
	// ~~~javascript
	// NewRegExp("\\d{3}").Exec("123456") // => [123]
	// NewRegExp("\\d+").Exec("123456") // => [123456]
	// NewRegExp("\\d+").Exec("123|456") // => [123]
	// NewRegExp("(\\d)(\\d)").Exec("123|456") // => [12, 1, 2]
	// ~~~
	Exec func(s string) []string `expr:"exec"`

	// 使用当前正则表达式分割字符串
	// ~~~javascript
	// NewRegExp("A").Split("123A456A789")  // => ["123", "456", "789"]
	// NewRegExp("A").Split("123456789") // => ["123456789"]
	// NewRegExp("\\|").Split("123|456|789") // => ["123", "456", "789"]
	// ~~~
	Split func(s string) []string `expr:"split"`
}

// NewRegExp 创建新正则表达式对象
//
// @internal
func NewRegExp(expr string) RegExp {
	reg, err := regexp.Compile(expr)
	if err != nil {
		reg = nil
	}
	return RegExp{
		Test: func(s string) bool {
			if reg == nil {
				return false
			}
			return reg.MatchString(s)
		},
		Exec: func(s string) []string {
			if reg == nil {
				return nil
			}
			return reg.FindStringSubmatch(s)
		},
		Split: func(s string) []string {
			if reg == nil {
				return nil
			}
			return reg.Split(s, -1)
		},
	}
}
