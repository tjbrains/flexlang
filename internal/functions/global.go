// Copyright 2026 FlexCDN root@flexcdn.cn. All rights reserved. Official site: https://flexcdn.cn .

package functions

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"hash/crc32"
	"math"
	"math/big"
	"net/url"
	"strconv"
	"strings"
)

// 已忽略的函数：
//  - eval ( x )
//  - isFinite ( number )
//  - decodeURI ( encodedURI )
//  - encodeURI ( uri )
//  - Encode ( string, extraUnescaped )
//  - Decode ( string, preserveEscapeSet )
//  - ParseHexOctet ( string, position )

type GlobalFunctions struct{}

var NaN = math.NaN()

func (this GlobalFunctions) IsNaN(n any) bool {
	switch x := n.(type) {
	case float64:
		return math.IsNaN(x)
	}
	return false
}

func (this GlobalFunctions) ParseFloat(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

func (this GlobalFunctions) ParseInt(s any, radix ...int) int64 {
	var str string

	switch x := s.(type) {
	case string:
		str = x
	case int:
		return int64(x)
	case int8:
		return int64(x)
	case int16:
		return int64(x)
	case int32:
		return int64(x)
	case int64:
		return x
	case uint:
		return int64(x)
	case uint8:
		return int64(x)
	case uint16:
		return int64(x)
	case uint32:
		return int64(x)
	case uint64:
		return int64(x) // maybe overflow
	case float32:
		return int64(x)
	case float64:
		return int64(x)
	case bool:
		if x {
			return 1
		}
		return 0
	default:
		return 0
	}

	if len(str) == 0 {
		return 0
	}

	var dotIndex = strings.Index(str, ".")
	if dotIndex >= 0 {
		str = str[:dotIndex]
	}

	if len(radix) == 0 {
		result, _ := strconv.ParseInt(str, 10, 64)
		return result
	}

	var radix0 = radix[0]
	if radix0 <= 0 {
		radix0 = 10
	}
	result, _ := strconv.ParseInt(str, radix0, 64)
	return result
}

func (this GlobalFunctions) DecodeURIComponent(encodedURIComponent string) string {
	s, _ := url.QueryUnescape(encodedURIComponent)
	return s
}

func (this GlobalFunctions) EncodeURIComponent(uriComponent string) string {
	return url.QueryEscape(uriComponent)
}

func (this GlobalFunctions) TypeOf(v any) string {
	if v == nil {
		return "undefined"
	}

	switch v.(type) {
	case bool:
		return "boolean"
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		return "number"
	case string:
		return "string"
	case big.Int, *big.Int:
		return "bigint"
	}

	return "object"
}

func (this GlobalFunctions) MD5(s string) string {
	var data = md5.Sum([]byte(s))
	return hex.EncodeToString(data[:])
}

func (this GlobalFunctions) Sha1(s string) string {
	var data = sha1.Sum([]byte(s))
	return hex.EncodeToString(data[:])
}

func (this GlobalFunctions) Sha256(s string) string {
	var data = sha256.Sum256([]byte(s))
	return hex.EncodeToString(data[:])
}

func (this GlobalFunctions) Crc32(s string) uint32 {
	return crc32.ChecksumIEEE([]byte(s))
}
