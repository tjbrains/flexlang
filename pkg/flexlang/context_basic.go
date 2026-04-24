// Copyright 2026 FlexCDN root@flexcdn.cn. All rights reserved. Official site: https://flexcdn.cn .

package flexlang

import (
	"github.com/tjbrains/flexlang/pkg/flexlang/functions"
)

type BasicContext struct {
	NaN                float64                                 `expr:"NaN"`
	IsNaN              func(number any) bool                   `expr:"isNaN"`
	ParseFloat         func(s string) float64                  `expr:"parseFloat"`
	ParseInt           func(s string, radix ...int) int64      `expr:"parseInt"`
	DecodeURIComponent func(encodedURIComponent string) string `expr:"decodeURIComponent"`
	EncodeURIComponent func(uriComponent string) string        `expr:"encodeURIComponent"`
	TypeOf             func(v any) string                      `expr:"typeOf"`

	MD5    func(s string) string `expr:"md5"`
	Sha1   func(s string) string `expr:"sha1"`
	Sha256 func(s string) string `expr:"sha256"`
	Crc32  func(s string) uint32 `expr:"crc32"`

	String struct {
		FromCharCode func(codeUnits ...rune) string                                 `expr:"fromCharCode"`
		At           func(s string, index int) string                               `expr:"at"`
		CharAt       func(s string, pos int) string                                 `expr:"charAt"`
		CharCodeAt   func(s string, pos int) any                                    `expr:"charCodeAt"`
		Concat       func(s string, args ...string) string                          `expr:"concat"`
		EndsWith     func(s string, searchString string, endPosition ...int) bool   `expr:"endsWith"`
		Includes     func(s string, searchString string, position ...int) bool      `expr:"includes"`
		IndexOf      func(s string, searchString string, position ...int) int       `expr:"indexOf"`
		LastIndexOf  func(s string, searchString string, position ...int) int       `expr:"lastIndexOf"`
		Match        func(s string, regexp any) []string                            `expr:"match"`
		Repeat       func(s string, count int) string                               `expr:"repeat"`
		Replace      func(s string, searchValue string, replaceValue string) string `expr:"replace"`
		ReplaceAll   func(s string, searchValue string, replaceValue string) string `expr:"replaceAll"`
		PadEnd       func(s string, maxLength int, fillString string) string        `expr:"padEnd"`
		PadStart     func(s string, maxLength int, fillString string) string        `expr:"padStart"`
		Slice        func(s string, start int, end int) string                      `expr:"slice"`
		Split        func(s string, separator string, limit ...int) []string        `expr:"split"`
		StartsWith   func(s string, searchString string, startPosition ...int) bool `expr:"startsWith"`
		Substring    func(s string, start int, end int) string                      `expr:"substring"`
		ToLowerCase  func(s string) string                                          `expr:"toLowerCase"`
		ToUpperCase  func(s string) string                                          `expr:"toUpperCase"`
		Trim         func(s string) string                                          `expr:"trim"`
		TrimEnd      func(s string) string                                          `expr:"trimEnd"`
		TrimPrefix   func(s string, prefix string) string                           `expr:"trimPrefix"`
		TrimStart    func(s string) string                                          `expr:"trimStart"`
		TrimSuffix   func(s string, suffix string) string                           `expr:"trimSuffix"`
		Length       func(s string) int                                             `expr:"length"`
		Sprintf      func(s string, args ...any) string                             `expr:"sprintf"`
	} `expr:"String"`

	Math struct {
		E           float64 `expr:"E"`
		LN10        float64 `expr:"LN10"`
		LOG10E      float64 `expr:"LOG10E"`
		PI          float64 `expr:"PI"`
		MathSQRT1_2 float64 `expr:"MathSQRT1_2"`
		SQRT2       float64 `expr:"SQRT2"`

		Expm1  func(x float64) float64                      `expr:"expm1"`
		Floor  func(x float64) int64                        `expr:"floor"`
		Hypot  func(args ...float64) float64                `expr:"hypot"`
		Log    func(x float64) float64                      `expr:"log"`
		Log1p  func(x float64) float64                      `expr:"log1p"`
		Log10  func(x float64) float64                      `expr:"log10"`
		Log2   func(x float64) float64                      `expr:"log2"`
		Pow    func(base float64, exponent float64) float64 `expr:"pow"`
		Random func() float64                               `expr:"random"`
		RandN  func(n int) int                              `expr:"randN"`
		Round  func(x float64) float64                      `expr:"round"`
		Sign   func(x float64) float64                      `expr:"sign"`
		Sin    func(x float64) float64                      `expr:"sin"`
		Sinh   func(x float64) float64                      `expr:"sinh"`
		Sqrt   func(x float64) float64                      `expr:"sqrt"`
		Tan    func(x float64) float64                      `expr:"tan"`
		Tanh   func(x float64) float64                      `expr:"tanh"`
		Trunc  func(x float64) int64                        `expr:"trunc"`
	} `expr:"Math"`

	Date struct {
		New func() functions.Date `expr:"new"`
	} `expr:"Date"`

	NewDate func() functions.Date `expr:"NewDate"`

	RegExp struct {
		Escape func(s string) string              `expr:"escape"`
		New    func(expr string) functions.RegExp `expr:"new"`
	} `expr:"RegExp"`

	NewRegExp func(expr string) functions.RegExp `expr:"NewRegExp"`

	JSON struct {
		Parse     func(text string) any  `expr:"parse"`
		Stringify func(value any) string `expr:"stringify"`
	} `expr:"JSON"`

	NewURL func(url string) functions.URL `expr:"NewURL"`

	Base64 struct {
		Encode func(s string) string `expr:"encode"`
		Decode func(s string) string `expr:"decode"`
	} `expr:"Base64"`

	Crypto struct {
		NewHMAC func(algorithm string, key string) (functions.CryptoHMACHash, error) `expr:"NewHMAC"`
	} `expr:"Crypto"`
}

func NewBasicContext() *BasicContext {
	var ctx = &BasicContext{}

	// Global
	var globalFunctions functions.GlobalFunctions
	ctx.NaN = functions.NaN
	ctx.IsNaN = globalFunctions.IsNaN
	ctx.ParseFloat = globalFunctions.ParseFloat
	ctx.ParseInt = globalFunctions.ParseInt
	ctx.DecodeURIComponent = globalFunctions.DecodeURIComponent
	ctx.EncodeURIComponent = globalFunctions.EncodeURIComponent
	ctx.TypeOf = globalFunctions.TypeOf

	ctx.MD5 = globalFunctions.MD5
	ctx.Sha1 = globalFunctions.Sha1
	ctx.Sha256 = globalFunctions.Sha256
	ctx.Crc32 = globalFunctions.Crc32

	// String
	var stringFunctions functions.StringFunctions
	ctx.String.FromCharCode = stringFunctions.FromCharCode
	ctx.String.At = stringFunctions.At
	ctx.String.CharAt = stringFunctions.CharAt
	ctx.String.CharCodeAt = stringFunctions.CharCodeAt
	ctx.String.Concat = stringFunctions.Concat
	ctx.String.EndsWith = stringFunctions.EndsWith
	ctx.String.Includes = stringFunctions.Includes
	ctx.String.IndexOf = stringFunctions.IndexOf
	ctx.String.LastIndexOf = stringFunctions.LastIndexOf
	ctx.String.Match = stringFunctions.Match
	ctx.String.Repeat = stringFunctions.Repeat
	ctx.String.Replace = stringFunctions.Replace
	ctx.String.ReplaceAll = stringFunctions.ReplaceAll
	ctx.String.PadEnd = stringFunctions.PadEnd
	ctx.String.PadStart = stringFunctions.PadStart
	ctx.String.Slice = stringFunctions.Slice
	ctx.String.Split = stringFunctions.Split
	ctx.String.StartsWith = stringFunctions.StartsWith
	ctx.String.Substring = stringFunctions.Substring
	ctx.String.ToLowerCase = stringFunctions.ToLowerCase
	ctx.String.ToUpperCase = stringFunctions.ToUpperCase
	ctx.String.Trim = stringFunctions.Trim
	ctx.String.TrimEnd = stringFunctions.TrimEnd
	ctx.String.TrimPrefix = stringFunctions.TrimPrefix
	ctx.String.TrimStart = stringFunctions.TrimStart
	ctx.String.TrimSuffix = stringFunctions.TrimSuffix
	ctx.String.Length = stringFunctions.Length
	ctx.String.Sprintf = stringFunctions.Sprintf

	// Math
	var mathFunctions functions.MathFunctions
	ctx.Math.E = functions.MathE
	ctx.Math.LN10 = functions.MathLN10
	ctx.Math.LOG10E = functions.MathLOG10E
	ctx.Math.PI = functions.MathPI
	ctx.Math.MathSQRT1_2 = functions.MathSQRT1_2
	ctx.Math.SQRT2 = functions.MathSQRT2

	ctx.Math.Expm1 = mathFunctions.Expm1
	ctx.Math.Floor = mathFunctions.Floor
	ctx.Math.Hypot = mathFunctions.Hypot
	ctx.Math.Log = mathFunctions.Log
	ctx.Math.Log1p = mathFunctions.Log1p
	ctx.Math.Log10 = mathFunctions.Log10
	ctx.Math.Log2 = mathFunctions.Log2
	ctx.Math.Pow = mathFunctions.Pow
	ctx.Math.Random = mathFunctions.Random
	ctx.Math.RandN = mathFunctions.RandN
	ctx.Math.Round = mathFunctions.Round
	ctx.Math.Sign = mathFunctions.Sign
	ctx.Math.Sin = mathFunctions.Sin
	ctx.Math.Sinh = mathFunctions.Sinh
	ctx.Math.Sqrt = mathFunctions.Sqrt
	ctx.Math.Tan = mathFunctions.Tan
	ctx.Math.Tanh = mathFunctions.Tanh
	ctx.Math.Trunc = mathFunctions.Trunc

	// Date
	var dateFunctions functions.DateFunctions
	ctx.Date.New = dateFunctions.New

	ctx.NewDate = functions.NewDate

	// RegExp
	var regexpFunctions functions.RegExpFunctions
	ctx.RegExp.Escape = regexpFunctions.Escape
	ctx.RegExp.New = regexpFunctions.New

	ctx.NewRegExp = functions.NewRegExp

	// JSON
	var jsonFunctions functions.JSONFunctions
	ctx.JSON.Parse = jsonFunctions.Parse
	ctx.JSON.Stringify = jsonFunctions.Stringify

	// URL
	ctx.NewURL = functions.NewURL

	// Base64
	var base64Functions functions.Base64Functions
	ctx.Base64.Encode = base64Functions.Encode
	ctx.Base64.Decode = base64Functions.Decode

	// Crypto
	ctx.Crypto.NewHMAC = functions.NewCryptoHMAC

	return ctx
}
