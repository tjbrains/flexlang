// Copyright 2026 FlexCDN root@flexcdn.cn. All rights reserved. Official site: https://flexcdn.cn .

package context

import (
	"github.com/tjbrains/flexlang/internal/functions"
	"github.com/tjbrains/flexlang/internal/visitors"
)

// BasicContext 为表达式求值提供内置全局 API，字段名与 expr 标签对应表达式中的标识符。
type BasicContext struct {
	// IEEE 754 非数字常量
	// 可以使用 IsNaN(number) 来判断一个变量是否为非数字
	NaN float64 `expr:"NaN"`

	// 判断是否为非数字（NaN）
	//
	// 示例：
	// ~~~javascript
	// isNaN(NaN) // => true
	// isNaN(10)  // => false
	// ~~~
	// @prototype isNaN(number any) bool
	IsNaN func(number any) bool `expr:"isNaN"`

	// 将字符串解析为浮点数
	//
	// 示例：
	// ~~~javascript
	// parseFloat("0") // => 0.0
	// parseFloat("abc") // => 0.0
	// parseFloat("123") // => 123.0
	// parseFloat("123.456") // => 123.456
	// ~~~
	// @prototype parseFloat(s string) float64
	ParseFloat func(s string) float64 `expr:"parseFloat"`

	// 将字符串按进制解析为整数
	//
	// 示例：
	// ~~~javascript
	// parseInt("") // => 0
	// parseInt("abc") // => 0
	// parseInt("123") // => 123
	// parseInt("123.456") // => 123
	// ~~~
	// @prototype parseInt(s string[, radix int]) int64
	ParseInt func(s string, radix ...int) int64 `expr:"parseInt"`

	// URI组件解码
	//
	// 示例：
	// ~~~javascript
	// decodeURIComponent("") // => ""
	// decodeURIComponent(`%25`) // => "%"
	// decodeURIComponent(`%3D`) // => "="
	// ~~~
	DecodeURIComponent func(encodedURIComponent string) string `expr:"decodeURIComponent"`

	// URI组件编码
	//
	// 示例：
	// ~~~javascript
	// encodeURIComponent("") // => ""
	// encodeURIComponent(`%`) // => "%25"
	// encodeURIComponent(`=`) // => "%3D"
	// ~~~
	EncodeURIComponent func(uriComponent string) string `expr:"encodeURIComponent"`

	// 返回值的类型名称
	//
	// 可能的值为undefined、boolean、number、string、bigint、object
	//
	// 示例：
	// ~~~javascript
	// typeOf("") // => string
	// typeOf("abc") // => string
	// typeOf(123) // => number
	// typeOf(123.0) // => number
	// typeOf(true) // => boolean
	// typeOf(false) // => boolean
	// ~~~
	TypeOf func(v any) string `expr:"typeOf"`

	// 计算MD5摘
	//
	// 返回十六进制字符串
	//
	// 示例：
	// ~~~javascript
	// md5("") // => "d41d8cd98f00b204e9800998ecf8427e"
	// md5("123456") // => "e10adc3949ba59abbe56e057f20f883e"
	// ~~~
	MD5 func(s string) string `expr:"md5"`

	// 计算SHA-1摘要
	//
	// 返回十六进制字符串
	//
	// 示例：
	// ~~~javascript
	// sha1("123456") // => 7c4a8d09ca3762af61e59520943dc26494f8941b
	// ~~~
	Sha1 func(s string) string `expr:"sha1"`

	// 计算 SHA-256摘要
	//
	// 返回十六进制字符串
	//
	// 示例：
	// ~~~javascript
	// sha256("123456") // => 8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92
	// ~~~
	Sha256 func(s string) string `expr:"sha256"`

	// 计算CRC32校验和
	//
	// 返回一个数字
	//
	// 示例：
	// ~~~javascript
	// crc32("123456") // => 158520161
	// ~~~
	Crc32 func(s string) uint32 `expr:"crc32"`

	// 字符串相关操作
	String struct {
		// 根据一组ASCII码生成字符串
		//
		// 示例：
		// ~~~javascript
		// String.fromCharCode() // => ""
		// String.fromCharCode(97, 98, 99) // => "abc"
		// ~~~
		FromCharCode func(codeUnits ...rune) string `expr:"fromCharCode"`

		// 返回指定位置的字符（以字符串的形式返回）
		//
		// 示例：
		// ~~~javascript
		// String.at("abc", 0) // => "a"
		// String.at("abc", 1) // => "b"
		// String.at("abc", 2) // => "c"
		// String.at("abc", 3) // => ""
		// String.at("abc", -1) // => "c"
		// String.at("abc", -2) // => "b"
		// String.at("abc", -3) // => "a"
		// String.at("abc", -4) // => ""
		// ~~~
		At func(s string, index int) string `expr:"at"`

		// 返回指定位置的字符（以字符串的形式返回）
		//
		// pos 不支持负数
		//
		// 示例：
		// ~~~javascript
		// String.charAt("abc", 0) // => "a"
		// String.charAt("abc", 1) // => "b"
		// String.charAt("abc", 2) // => "c"
		// String.charAt("abc", 3) // => ""
		// String.charAt("abc", -1) // => ""
		// ~~~
		CharAt func(s string, pos int) string `expr:"charAt"`

		// 读取指定位置的ASCII码
		//
		// 超出范围范围 NaN
		//
		// 示例：
		// ~~~javascript
		// String.charCodeAt("abc", 0) // => 97
		// String.charCodeAt("abc", 1) // => 98
		// String.charCodeAt("abc", 2) // => 99
		// String.charCodeAt("abc", 3) // => NaN
		// String.charCodeAt("abc", -1) // => NaN
		// ~~~
		CharCodeAt func(s string, pos int) any `expr:"charCodeAt"`

		// 拼接多个字符串
		//
		// 示例：
		// ~~~javascript
		// String.concat("") // => ""
		// String.concat("", "a", "b", "c") // => "abc"
		// String.concat("ABC") // => "ABC"
		// String.concat("ABC", "a", "b", "c") // => "ABCabc"
		// ~~~
		Concat func(s string, args ...string) string `expr:"concat"`

		// 检查字符串是否以某个子串结尾
		//
		// 示例：
		// ~~~javascript
		// String.endsWith("abc", "c") // => true
		// String.endsWith("abc", "b") // => false
		// String.endsWith("abc", "c", -1) // => false
		// String.endsWith("abc", "c", 3) // => true
		// String.endsWith("abc", "c", 4) // => true
		// String.endsWith("abc", "b", 2) // => true
		// String.endsWith("abc", "a", 1) // => true
		// String.endsWith("abc", "", 0) // => true
		// ~~~
		// @prototype endsWith(s string, searchString string[, endPosition int]) bool
		EndsWith func(s string, searchString string, endPosition ...int) bool `expr:"endsWith"`

		// 检查字符串是否包含某个子串
		//
		// 示例：
		// ~~~javascript
		// String.includes("abcdefg", "cde") // => true
		// String.includes("abcdefg", "bde") // => false
		// String.includes("abcdefg", "cde", 6)) // => true
		// String.includes("abcdefg", "cde", 3) // => false
		// ~~~
		// @prototype includes(s string, searchString string[, position int]) bool
		Includes func(s string, searchString string, position ...int) bool `expr:"includes"`

		// 查找子串首次出现位置
		//
		// 如果未找到则返回 -1
		//
		// 示例：
		// ~~~javascript
		// String.indexOf("abc", "defg") // => -1
		// String.indexOf("abcdefg", "abc") // => 0
		// String.indexOf("abcdefg", "bcd") // => 1
		// String.indexOf("abcdefg", "") // => 0
		// String.indexOf("abcdefg", "", 1) // => 1
		// String.indexOf("abcdefg", "cd", 1) // => 2
		// String.indexOf("abcdefg", "cd", 3) // => -1
		// String.indexOf("abcdefg", "cd", 10) // => -1
		// ~~~
		// @prototype indexOf(s string, searchString string[, position int]) int
		IndexOf func(s string, searchString string, position ...int) int `expr:"indexOf"`

		// 查找子串最后一次出现的位置
		//
		// 如果未找到则返回 -1
		//
		// 示例：
		// ~~~javascript
		// String.lastIndexOf("abcdefg", "ff") // => -1
		// String.lastIndexOf("abcdefg", "abc") // => 0
		// String.lastIndexOf("abcdefg", "fg") // => 5
		// String.lastIndexOf("abcdefg", "fg", 5) // => 5
		// String.lastIndexOf("abcdefg", "fg", 6) // => 5
		// String.lastIndexOf("abcdefg", "fg", 10) // => 5
		// String.lastIndexOf("abcdefg", "") // => 7
		// String.lastIndexOf("abcdefg", "", 3) // => 3
		// String.lastIndexOf("abcdefg", "", 10) // => 7
		// String.lastIndexOf("abcdefg", "", -1) // => -1
		// ~~~
		// @prototype lastIndexOf(s string, searchString string[, position int]) int
		LastIndexOf func(s string, searchString string, position ...int) int `expr:"lastIndexOf"`

		// 匹配正则
		//
		// 返回匹配的片段
		//
		//示例：
		// ~~~javascript
		// String.match("abc", "\\w+") // => ["abc"]
		// String.match("ab|c", "\\w+") // => ["ab"]
		// String.match("()*&)", "\\w+") // => nil
		// String.match("abc", NewRegExp("\\w+")) // => ["abc"]
		// String.match("ab|c", NewRegExp("\\w+")) // => ["ab"]
		// ~~~
		Match func(s string, regexp any) []string `expr:"match"`

		// 重复字符串
		//
		// 示例：
		// ~~~javascript
		// String.repeat("abc", 0) // => ""
		// String.repeat("abc", -1) // => ""
		// String.repeat("abc", 10) // => "abcabcabcabcabcabcabcabcabcabc"
		// ~~~
		Repeat func(s string, count int) string `expr:"repeat"`

		// 替换首个匹配子串
		//
		// 示例：
		// ~~~javascript
		// String.replace("abcdefg", "", "") // => "abcdefg"
		// String.replace("abcdefg", "a", "1") // => "1bcdefg"
		// String.replace("abcdefg", "b", "1") // => "a1cdefg"
		// String.replace("abcabc", "b", "1") // => "a1cabc"
		// ~~~
		Replace func(s string, searchValue string, replaceValue string) string `expr:"replace"`

		// 替换全部匹配子串
		//
		// 示例：
		// ~~~javascript
		// String.replaceAll("abcabc", "b", "1") // => "a1ca1c"
		// ~~~
		ReplaceAll func(s string, searchValue string, replaceValue string) string `expr:"replaceAll"`

		// 从尾部填充字符串至指定长度
		//
		// 示例：
		// ~~~javascript
		// String.padEnd("abc", -1, "ABC") // => "abc"
		// String.padEnd("abc", 0, "ABC") // => "abc"
		// String.padEnd("abc", 3, "ABC") // => "abc"
		// String.padEnd("abc", 4, "ABC") // => "abcA"
		// String.padEnd("abc", 6, "ABC") // => "abcABC"
		// String.padEnd("abc", 9, "ABC") // => "abcABCABC"
		// String.padEnd("abc", 14, "ABC") // => "abcABCABCABCAB"
		// String.padEnd("abc", 14, "") // => "abc"
		// String.padEnd("abc", 14, " ") // => "abc           "
		// ~~~
		PadEnd func(s string, maxLength int, fillString string) string `expr:"padEnd"`

		// 从首部填充至指定长度
		//
		// 示例：
		// ~~~javascript
		// String.padStart("abc", -1, "ABC") // => "abc"
		// String.padStart("abc", 0, "ABC") // => "abc"
		// String.padStart("abc", 3, "ABC") // => "abc"
		// String.padStart("abc", 4, "ABC") // => "Aabc"
		// String.padStart("abc", 6, "ABC") // => "ABCabc"
		// String.padStart("abc", 9, "ABC") // => "ABCABCabc"
		// String.padStart("abc", 14, "ABC") // => "ABCABCABCABabc"
		// String.padStart("abc", 14, "") // => "abc"
		// String.padStart("abc", 14, " ") // => "           abc"
		// ~~~
		PadStart func(s string, maxLength int, fillString string) string `expr:"padStart"`

		// 截取字符串一部分
		//
		// 示例：
		// ~~~javascript
		// String.slice("", 0, 0) // => ""
		// String.slice("abc", -3, -2) // => "a"
		// String.slice("abc", -3, -1) // => "ab"
		// String.slice("abc", 0, 10) // => "abc"
		// String.slice("abc", 1, 3) // => "bc"
		// String.slice("abc", 1, 1) // => ""
		// String.slice("abc", 1, 2) // => "b"
		// String.slice("abcdefg", -3, 7) // => "efg"
		// String.slice("abcdefg", 10, 100) // => ""
		// ~~~
		Slice func(s string, start int, end int) string `expr:"slice"`

		// 使用分隔符分割字符串
		//
		// 示例：
		// ~~~javascript
		// String.split("", "") // => []
		// String.split("abc", "") // => ["a", "b", "c"]
		// String.split("abc", "b") // => ["a", "c"]
		// String.split("abc", "b", -1) // => ["a", "c"]
		// String.split("abc", "b", 1) // => ["a"]
		// String.split("abc", "b", 2) // => ["a", "c"]
		// String.split("abc", "d", 2) // => ["abc"]
		// ~~~
		// @prototype split(s string, separator string[, limit int]) []string
		Split func(s string, separator string, limit ...int) []string `expr:"split"`

		// 检查字符串是否以某个子串开头
		//
		// 示例：
		// ~~~javascript
		// String.startsWith("abc", "a") // => true
		// String.startsWith("abc", "ab") // => true
		// String.startsWith("abc", "b") // => false
		// String.startsWith("abc", "a", -1) // => true
		// String.startsWith("abc", "a", 0) // => true
		// String.startsWith("abc", "b", 1) // => true
		// String.startsWith("abcdefg", "ef", 4) // => true
		// String.startsWith("abc", "", 1) // => true
		// String.startsWith("abc", "", 0) // => true
		// String.startsWith("abc", "", -1) // => true
		// ~~~
		// @prototype startsWith(s string, searchString string[, startPosition int]) bool
		StartsWith func(s string, searchString string, startPosition ...int) bool `expr:"startsWith"`

		// 获取字符串一部分
		//
		// 示例：
		// ~~~javascript
		// String.substring("", 1, 2) // => ""
		// String.substring("abcdefg", 1, 3) // => "bc"
		// String.substring("abcdefg", 0, 3) // => "abc"
		// String.substring("abcdefg", -1, 3) // => "abc"
		// String.substring("abcdefg", 3, 0) // => "abc"
		// String.substring("abcdefg", 3, 1) // => "bc"
		// String.substring("abcdefg", 0, 100) // => "abcdefg"
		// String.substring("abcdefg", 6, 100) // "g"
		// String.substring("abcdefg", 7, 100) // => ""
		// String.substring("abcdefg", 10, 100) // => ""
		// ~~~
		Substring func(s string, start int, end int) string `expr:"substring"`

		// 将字符串转换为小写
		//
		// 示例：
		// ~~~javascript
		// String.toLowerCase("") // => ""
		// String.toLowerCase("abc") // => "abc"
		// String.toLowerCase("ABC") // => "abc"
		// String.toLowerCase("aBC") // => "abc"
		// String.toLowerCase("ABC中文") // => "abc中文"
		// ~~~
		ToLowerCase func(s string) string `expr:"toLowerCase"`

		// 将字符串转换为大写
		//
		// 示例：
		// ~~~javascript
		// String.toUpperCase("ABC") // => ABC
		// String.toUpperCase("abc") // => ABC
		// String.toUpperCase("Abc") // => ABC
		// String.toUpperCase("abc中文") // => ABC中文
		// ~~~
		ToUpperCase func(s string) string `expr:"toUpperCase"`

		// 去除字符串首尾空白
		//
		// 示例：
		// ~~~javascript
		// String.trim("") // => ""
		// String.trim(" a ") // => "a"
		// String.trim(" a b \t\n\r\t") // => "a b"
		// ~~~
		Trim func(s string) string `expr:"trim"`

		// 去除字符串尾部空白
		//
		// 示例：
		// ~~~javascript
		// String.trimEnd("") // => ""
		// String.trimEnd(" a ") // => " a"
		// String.trimEnd(" a b \t\n\r\t") // => " a b"
		// ~~~
		TrimEnd func(s string) string `expr:"trimEnd"`

		// 去除字符串前缀
		//
		// 示例：
		// ~~~javascript
		// String.trimPrefix("abcdefg", "ab") // => "cdefg"
		// String.trimPrefix("abcdefg", "") // => "abcdefg"
		// String.trimPrefix("ababcdefg", "ab") // => "abcdefg"
		// ~~~
		TrimPrefix func(s string, prefix string) string `expr:"trimPrefix"`

		// 去除字符串首部空白
		//
		// 示例：
		// ~~~javascript
		// String.trimStart("") // => ""
		// String.trimStart(" a ") // => "a "
		// String.trimStart(" \t\n\r\t a b") // => "a b"
		// ~~~
		TrimStart func(s string) string `expr:"trimStart"`

		// 去除字符串后缀
		//
		// 示例：
		// ~~~javascript
		// String.trimSuffix("abcdefg", "fg") // => "abcde"
		// String.trimSuffix("abcdefg", "") // => "abcdefg"
		// String.trimSuffix("abcdefgfg", "fg") // => "abcdefg"
		// ~~~
		TrimSuffix func(s string, suffix string) string `expr:"trimSuffix"`

		// 获取字符串长度
		//
		// 示例：
		// ~~~javascript
		// String.length("") // => 0
		// String.length("abc") // => 3
		// String.length("abc中文") // => 5
		// ~~~
		Length func(s string) int `expr:"length"`

		// 格式化字符串
		//
		// 示例：
		// ~~~javascript
		// String.sprintf("%s %d, %.2f", "abc", 1, 2.3456) // => "abc 1, 2.35"
		// ~~~
		Sprintf func(s string, args ...any) string `expr:"sprintf"`
	} `expr:"String"`

	// 数学相关操作
	Math struct {
		// 自然常数 e
		E float64 `expr:"E"`

		// ln(10)
		LN10 float64 `expr:"LN10"`

		// log10(e)
		LOG10E float64 `expr:"LOG10E"`

		// 圆周率 π
		PI float64 `expr:"PI"`

		// 1/√2
		MathSQRT1_2 float64 `expr:"MathSQRT1_2"`

		// √2
		SQRT2 float64 `expr:"SQRT2"`

		// 计算绝对值
		//
		// 示例：
		// ~~~javascript
		// Math.abs(2.0) // => 2.0
		// Math.abs(-2.123) // => 2.123
		// ~~~
		Abs func(x float64) float64 `expr:"abs"`

		// 计算立方根
		//
		// 示例：
		// ~~~javascript
		// Math.cbrt(8.0) // => 2.0
		// Math.cbrt(-27.0) // => -3.0
		// ~~~
		Cbrt func(x float64) float64 `expr:"cbrt"`

		// 向上取整
		//
		// 示例：
		// ~~~javascript
		// Math.ceil(1.0) // => 1
		// Math.ceil(1.0234) // => 2
		// Math.ceil(1.5) // => 2
		// Math.ceil(123.678) // => 124
		// ~~~
		Ceil func(x float64) int64 `expr:"ceil"`

		// 计算uint32数值的二进制前导零数量
		//
		// 示例：
		// ~~~javascript
		// Math.clz32(1) // => 31
		// Math.clz32(4) // => 29
		// Math.clz32(0) // => 32
		// Math.clz32(2147483648) // => 0
		// ~~~
		Clz32 func(x uint32) int `expr:"clz32"`

		// 计算余弦值
		//
		// 示例：
		// ~~~javascript
		// Math.cos(Math.PI/3) // => 0.49999999999999994
		// ~~~
		Cos func(x float64) float64 `expr:"cos"`

		// 计算双曲余弦值
		//
		// 示例：
		// ~~~javascript
		// Math.cosh(30) // => 5.343237290762231e+12
		// Math.cosh(60) // => 5.710036949078421e+25
		// ~~~
		Cosh func(x float64) float64 `expr:"cosh"`

		// 计算 e^x
		//
		// 示例：
		// ~~~javascript
		// Math.exp(100.0) // => 2.6881171418161356e+43
		// ~~~
		Exp func(x float64) float64 `expr:"exp"`

		// 计算 e^x - 1
		//
		// 示例：
		// ~~~javascript
		// Math.expm1(100.0) // => 2.6881171418161356e+43
		// ~~~
		Expm1 func(x float64) float64 `expr:"expm1"`

		// 向下取整
		//
		// 示例：
		// ~~~javascript
		// Math.floor(1.0) // => 1
		// Math.floor(1.0234) // => 1
		// Math.floor(123.678) // => 123
		// ~~~
		Floor func(x float64) int64 `expr:"floor"`

		// 计算平方和开方
		//
		// 示例：
		// ~~~javascript
		// Math.hypot(3, 4) // => 5.0
		// Math.hypot(-3, -4) // => 5.0
		// Math.hypot(0, 0) // => 0.0
		// ~~~
		Hypot func(args ...float64) float64 `expr:"hypot"`

		// 计算自然对数 ln(x)
		//
		// 示例：
		// ~~~javascript
		// Math.log(2) // => 0.6931471805599453
		// ~~~
		Log func(x float64) float64 `expr:"log"`

		// 计算 ln(1+x)
		//
		// 示例：
		// ~~~javascript
		// Math.log1p(2) // => 1.0986122886681096
		// ~~~
		Log1p func(x float64) float64 `expr:"log1p"`

		// 计算常用对数 log10(x)
		//
		// 示例：
		// ~~~javascript
		// Math.log10(2) // => 0.3010299956639812
		// ~~~
		Log10 func(x float64) float64 `expr:"log10"`

		// 计算以2为底的对数
		//
		// 示例：
		// ~~~javascript
		// Math.log2(2) // => 1
		// Math.log2(10) // => 3.321928094887362
		// ~~~
		Log2 func(x float64) float64 `expr:"log2"`

		// 幂运算
		//
		// 示例：
		// ~~~javascript
		// Math.pow(2.0, 0.0) // => 1.0
		// Math.pow(4.0, 0.5 // => 2.0
		// isNaN(Math.pow(-4.0, 0.5)) // => true
		// ~~~
		Pow func(base float64, exponent float64) float64 `expr:"pow"`

		// 生成随机浮点数
		//
		// 范围在 [0,1) 之间
		//
		// 示例：
		// ~~~javascript
		// Math.random() // => 0.6826845041002553 (类似的小数，非固定值)
		// ~~~
		Random func() float64 `expr:"random"`

		// 生成随机整数
		//
		// 范围在 [0, n) 之间
		//
		// 示例：
		// ~~~javascript
		// Math.randN(10) // => 5 （类似的整数，在0-9之间（不含10），非固定值）
		// Math.randN(0)) // => 0
		// Math.randN(-10) // => 0
		// ~~~
		RandN func(n int) int `expr:"randN"`

		// 四舍五入到最近浮点数形式的整数
		//
		// 示例：
		// ~~~javascript
		// Math.round(2.0) // => 2.0
		// Math.round(2.0123) // => 2.0
		// Math.round(2.6123) // => 3.0
		// ~~~
		Round func(x float64) float64 `expr:"round"`

		// 根据数值正负生成符号数
		//
		// 返回-1.0、0.0、1.0
		//
		// 示例：
		// ~~~javascript
		// Math.sign(1.0) // => 1.0
		// Math.sign(-123456.123) // => -1.0
		// Math.sign(0) // => 0.0
		// ~~~
		Sign func(x float64) float64 `expr:"sign"`

		// 计算正弦值
		//
		// 示例：
		// ~~~javascript
		// Math.sin(Math.PI/6) // => 0.49999999999999994
		// ~~~
		Sin func(x float64) float64 `expr:"sin"`

		// 计算双曲正弦值
		//
		// 示例：
		// ~~~javascript
		// Math.sinh(30) // => 5.343237290762231e+12
		// Math.sinh(60) // => 5.710036949078421e+25
		// ~~~
		Sinh func(x float64) float64 `expr:"sinh"`

		// 计算平方根
		//
		// 示例：
		// ~~~javascript
		// Math.sqrt(4) // => 2.0
		// Math.sqrt(9) // => 3.0
		// ~~~
		Sqrt func(x float64) float64 `expr:"sqrt"`

		// 计算正切值
		//
		// 示例：
		// ~~~javascript
		// Math.tan(math.Pi / 4) // => 0.9999999999999998
		// ~~~
		Tan func(x float64) float64 `expr:"tan"`

		// 计算双曲正切值
		//
		// 示例：
		// ~~~javascript
		// Math.tanh(45) // => 1
		// ~~~
		Tanh func(x float64) float64 `expr:"tanh"`

		// 向零截断取整
		//
		// 示例：
		// ~~~javascript
		// Math.trunc(2.0) // => 2
		// Math.trunc(2.0123) // => 2
		// Math.trunc(34.6789) // => 34
		// ~~~
		Trunc func(x float64) int64 `expr:"trunc"`
	} `expr:"Math"`

	// 日期相关操作
	Date struct {
		// 构造当前时间的日期对象
		//
		// 示例：
		// ~~~javascript
		// let d = Date.new();
		// let year = string(d.getFullYear());
		// let month = String.padStart(string(d.getMonth()+1), 2, '0');
		// let day = String.padStart(string(d.getDate()), 2, '0');
		// year + "-" + month + "-" + day // => 类似于 "2026-04-28" 的当前日期
		// ~~~
		New func() functions.Date `expr:"new"`
	} `expr:"Date"`

	// 构造日期对象
	//
	// 同 Date.new()
	NewDate func() functions.Date `expr:"NewDate"`

	// 正则表达式相关操作
	RegExp struct {
		// 转义用于正则表达式的特殊字符
		//
		// 示例：
		// ~~~javascript
		// RegExp.escape("#$%*()-[]{}|") // => "#\$%\*\(\)-\[\]\{\}\|"
		// ~~~
		Escape func(s string) string `expr:"escape"`

		// 构造新正则表达式
		//
		// 示例：
		// ~~~javascript
		// RegExp.new("\\w+").test("abc") // => true
		// ~~~
		New func(expr string) functions.RegExp `expr:"new"`
	} `expr:"RegExp"`

	// 构造新正则表达式
	//
	// 同 RegExp.New(expr)
	//
	// 示例：
	// ~~~javascript
	// NewRegExp("\\w+").test("abc") // => true
	// ~~~
	NewRegExp func(expr string) functions.RegExp `expr:"NewRegExp"`

	// JSON相关操作
	JSON struct {
		// 解析JSON文本
		//
		// 示例：
		// ~~~javascript
		// let v = JSON.parse('{ "a": 1, "b": 2 }');
		// v.a // => 1
		// ~~~
		Parse func(text string) any `expr:"parse"`

		// 将数据序列化为JSON文本
		//
		// 示例：
		// ~~~javascript
		// JSON.stringify({"a":1, "b": 2} // => "{"a":1,"b":2}"
		// ~~~
		Stringify func(value any) string `expr:"stringify"`
	} `expr:"JSON"`

	// 构造新URL对象
	//
	// 示例：
	// ~~~javascript
	// let u = NewURL("https://user:pass@example.com:8080/docs?nav=1#link");
	// u.port // => 8080
	// u.hash // => "link"
	// u.host // => "example.com:8080"
	// u.query // => "nav=1"
	// u.scheme // => "https"
	// u.path // => "/docs"
	// u.opaque // => ""
	// u.user // => {"password":"pass", "username":"user"}
	// ~~~
	NewURL func(url string) functions.URL `expr:"NewURL"`

	// Base64相关操作
	Base64 struct {
		// Base64编码
		//
		// 示例：
		// ~~~javascript
		// Base64.encode("Hello, World!") // => "SGVsbG8sIFdvcmxkIQ=="
		// ~~~
		Encode func(s string) string `expr:"encode"`

		// Base64 解码
		//
		// 示例：
		// ~~~javascript
		// Base64.decode("SGVsbG8sIFdvcmxkIQ==") // => "Hello, World!"
		// ~~~
		Decode func(s string) string `expr:"decode"`
	} `expr:"Base64"`

	// 加密相关操作
	Crypto struct {
		// 创建HMAC摘要器
		//
		// algorithm 支持md5、sha1、sha256、sha512
		//
		// 示例：
		// ~~~javascript
		// let h = Crypto.NewHMAC("sha1", "123456");
		// h.update("ABCDEFG");
		// h.sum() // => "02ce4f53c007cb7c90f39b953e1fe71a3c6a178d"
		// ~~~
		NewHMAC func(algorithm string, key string) (functions.CryptoHMACHash, error) `expr:"NewHMAC"`
	} `expr:"Crypto"`

	// 控制台相关操作
	Console struct {
		// 输出日志
		//
		// 示例：
		// ~~~javascript
		// console.log("Hello", "World")
		// ~~~
		Log func(data ...any) bool `expr:"log"`
	} `expr:"console"`

	visitor *visitors.Visitor
	printer func(s ...string)
}

// NewBasicContext 创建基础上下文对象
//
// @internal
func NewBasicContext() *BasicContext {
	var ctx = &BasicContext{
		visitor: visitors.NewVisitor(),
	}

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

	ctx.Math.Abs = mathFunctions.Abs
	ctx.Math.Cbrt = mathFunctions.Cbrt
	ctx.Math.Ceil = mathFunctions.Ceil
	ctx.Math.Clz32 = mathFunctions.Clz32
	ctx.Math.Cos = mathFunctions.Cos
	ctx.Math.Cosh = mathFunctions.Cosh
	ctx.Math.Exp = mathFunctions.Exp
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

	// Console
	var consoleFunctions functions.ConsoleFunctions
	ctx.Console.Log = func(data ...any) bool {
		var printer = ctx.printer
		if printer == nil {
			return false
		}
		consoleFunctions.Log(printer, data...)
		return true
	}

	return ctx
}

// WithPrinter 设置打印器
//
// @internal
func (this *BasicContext) WithPrinter(printer func(s ...string)) {
	this.printer = printer
}

// Visitor 读取节点访问器
//
// @internal
func (this *BasicContext) Visitor() *visitors.Visitor {
	return this.visitor
}

// Reset 重置
//
// @internal
func (this *BasicContext) Reset() {
	this.visitor.Reset()
}
