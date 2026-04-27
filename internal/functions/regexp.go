// Copyright 2026 FlexCDN root@flexcdn.cn. All rights reserved. Official site: https://flexcdn.cn .

package functions

import "regexp"

type RegExpFunctions struct {
}

func (this RegExpFunctions) Escape(s string) string {
	return regexp.QuoteMeta(s)
}

func (this RegExpFunctions) New(expr string) RegExp {
	return NewRegExp(expr)
}

type RegExp struct {
	Test  func(s string) bool     `expr:"test"`
	Exec  func(s string) []string `expr:"exec"`
	Split func(s string) []string `expr:"split"`
}

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
	}
}
