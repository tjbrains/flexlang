// Copyright 2026 FlexCDN root@flexcdn.cn. All rights reserved. Official site: https://flexcdn.cn .

package functions

import "time"

// 忽略的函数
//  - Date.prototype.getTimezoneOffset ( )
//  - Date.prototype.getUTCDate ( )
//  - Date.prototype.getUTCDay ( )
//  - Date.prototype.getUTCFullYear ( )
//  - Date.prototype.getUTCHours ( )
//  - Date.prototype.getUTCMilliseconds ( )
//  - Date.prototype.getUTCMinutes ( )
//  - Date.prototype.getUTCMonth ( )
//  - Date.prototype.getUTCSeconds ( )
//  - Date.prototype.setDate ( date )
//  - Date.prototype.setFullYear ( year [ , month [ , date ] ] )
//  - Date.prototype.setHours ( hour [ , min [ , sec [ , ms ] ] ] )
//  - Date.prototype.setMilliseconds ( ms )
//  - Date.prototype.setMinutes ( min [ , sec [ , ms ] ] )
//  - Date.prototype.setMonth ( month [ , date ] )
//  - Date.prototype.setSeconds ( sec [ , ms ] )
//  - Date.prototype.setTime ( time )
//  - Date.prototype.setUTCDate ( date )
//  - Date.prototype.setUTCFullYear ( year [ , month [ , date ] ] )
//  - Date.prototype.setUTCHours ( hour [ , min [ , sec [ , ms ] ] ] )
//  - Date.prototype.setUTCMilliseconds ( ms )
//  - Date.prototype.setUTCMinutes ( min [ , sec [ , ms ] ] )
//  - Date.prototype.setUTCMonth ( month [ , date ] )
//  - Date.prototype.setUTCSeconds ( sec [ , ms ] )
//  - Date.prototype.toDateString ( )
//  - Date.prototype.toISOString ( )
//  - Date.prototype.toJSON ( key )
//  - Date.prototype.toLocaleDateString ( [ reserved1 [ , reserved2 ] ] )
//  - Date.prototype.toLocaleString ( [ reserved1 [ , reserved2 ] ] )
//  - Date.prototype.toLocaleTimeString ( [ reserved1 [ , reserved2 ] ] )
//  - Date.prototype.toString ( )
//  - Date.prototype.toTimeString ( )
//  - Date.prototype.toUTCString ( )
//  - Date.prototype.valueOf ( )

type DateFunctions struct{}

// New 创建新日期对象
//
// @internal
func (this DateFunctions) New() Date {
	return NewDate()
}

// NewDate 创建新日期对象
//
// @internal
func NewDate() Date {
	var now = time.Now()

	return Date{
		GetDate: func() int {
			return now.Day()
		},
		GetDay: func() int {
			return int(now.Weekday())
		},
		GetFullYear: func() int {
			return now.Year()
		},
		GetHours: func() int {
			return now.Hour()
		},
		GetMilliseconds: func() int {
			return int(now.UnixMilli() % 1000)
		},
		GetMinutes: func() int {
			return now.Minute()
		},
		GetMonth: func() int {
			return int(now.Month()) - 1
		},
		GetSeconds: func() int {
			return now.Second()
		},
		GetTime: func() int64 {
			return now.UnixMilli()
		},
	}
}

type Date struct {
	// 获取日期
	//
	// 1-31
	GetDate func() int `expr:"getDate"`

	// 获取一周中的天
	//
	// 0-6
	GetDay func() int `expr:"getDay"`

	// 获取年份
	//
	// 类似于 2006
	GetFullYear func() int `expr:"getFullYear"`

	// 获取24制小时数
	//
	// 类似于 1、11、16等
	GetHours func() int `expr:"getHours"`

	// 获取当前时间戳毫秒部分
	//
	// 比如 150、320 等
	GetMilliseconds func() int `expr:"getMilliseconds"`

	// 获取当前分钟数
	//
	// 类似于 1、15、45
	GetMinutes func() int `expr:"getMinutes"`

	// 获取当前月数
	//
	// 值为 0-11
	GetMonth func() int `expr:"getMonth"`

	// 获取当前秒数
	//
	// 类似于 1、10、35 等
	GetSeconds func() int `expr:"getSeconds"`

	// 获取当前时间戳
	//
	// 含毫秒
	//
	// 类似于 1777369354150
	GetTime func() int64 `expr:"getTime"`
}
