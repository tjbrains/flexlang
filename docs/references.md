# 变量和函数
## $req
> $req [Request](#request对象)

请求对象

可以通过此对象读取请求相关信息

示例：
~~~javascript
$req.url() // 返回当前URL，类似于 "https://example.com/hello?name=lily"
$req.query().get("name") // 获取参数值，在上面URL中，结果就是 "lily"
~~~

## $resp
> $resp [Response](#response对象)

响应对象

可以通过此对象设置响应的内容

示例：
~~~javascript
$resp.send(200, "Hello, World!")
~~~

## NaN
> NaN

IEEE 754 非数字常量
可以使用 IsNaN(number) 来判断一个变量是否为非数字

## isNaN
> isNaN(number any) bool

判断是否为非数字（NaN）

示例：
~~~javascript
isNaN(NaN) // => true
isNaN(10)  // => false
~~~

## parseFloat
> parseFloat(s string) float64

将字符串解析为浮点数

示例：
~~~javascript
parseFloat("0") // => 0.0
parseFloat("abc") // => 0.0
parseFloat("123") // => 123.0
parseFloat("123.456") // => 123.456
~~~

## parseInt
> parseInt(s string\[, radix int]) int64

将字符串按进制解析为整数

示例：
~~~javascript
parseInt("") // => 0
parseInt("abc") // => 0
parseInt("123") // => 123
parseInt("123.456") // => 123
~~~

## decodeURIComponent
> decodeURIComponent(encodedURIComponent string) string

URI组件解码

示例：
~~~javascript
decodeURIComponent("") // => ""
decodeURIComponent(`%25`) // => "%"
decodeURIComponent(`%3D`) // => "="
~~~

## encodeURIComponent
> encodeURIComponent(uriComponent string) string

URI组件编码

示例：
~~~javascript
encodeURIComponent("") // => ""
encodeURIComponent(`%`) // => "%25"
encodeURIComponent(`=`) // => "%3D"
~~~

## typeOf
> typeOf(v any) string

返回值的类型名称

可能的值为undefined、boolean、number、string、bigint、object

示例：
~~~javascript
typeOf("") // => string
typeOf("abc") // => string
typeOf(123) // => number
typeOf(123.0) // => number
typeOf(true) // => boolean
typeOf(false) // => boolean
~~~

## md5
> md5(s string) string

计算MD5摘

返回十六进制字符串

示例：
~~~javascript
md5("") // => "d41d8cd98f00b204e9800998ecf8427e"
md5("123456") // => "e10adc3949ba59abbe56e057f20f883e"
~~~

## sha1
> sha1(s string) string

计算SHA-1摘要

返回十六进制字符串

示例：
~~~javascript
sha1("123456") // => 7c4a8d09ca3762af61e59520943dc26494f8941b
~~~

## sha256
> sha256(s string) string

计算 SHA-256摘要

返回十六进制字符串

示例：
~~~javascript
sha256("123456") // => 8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92
~~~

## crc32
> crc32(s string) uint32

计算CRC32校验和

返回一个数字

示例：
~~~javascript
crc32("123456") // => 158520161
~~~

## String
> String

字符串相关操作

### String.fromCharCode
> String.fromCharCode(codeUnits ...rune) string

根据一组ASCII码生成字符串

示例：
~~~javascript
String.fromCharCode() // => ""
String.fromCharCode(97, 98, 99) // => "abc"
~~~

### String.at
> String.at(s string, index int) string

返回指定位置的字符（以字符串的形式返回）

示例：
~~~javascript
String.at("abc", 0) // => "a"
String.at("abc", 1) // => "b"
String.at("abc", 2) // => "c"
String.at("abc", 3) // => ""
String.at("abc", -1) // => "c"
String.at("abc", -2) // => "b"
String.at("abc", -3) // => "a"
String.at("abc", -4) // => ""
~~~

### String.charAt
> String.charAt(s string, pos int) string

返回指定位置的字符（以字符串的形式返回）

pos 不支持负数

示例：
~~~javascript
String.charAt("abc", 0) // => "a"
String.charAt("abc", 1) // => "b"
String.charAt("abc", 2) // => "c"
String.charAt("abc", 3) // => ""
String.charAt("abc", -1) // => ""
~~~

### String.charCodeAt
> String.charCodeAt(s string, pos int) any

读取指定位置的ASCII码

超出范围范围 NaN

示例：
~~~javascript
String.charCodeAt("abc", 0) // => 97
String.charCodeAt("abc", 1) // => 98
String.charCodeAt("abc", 2) // => 99
String.charCodeAt("abc", 3) // => NaN
String.charCodeAt("abc", -1) // => NaN
~~~

### String.concat
> String.concat(s string, args ...string) string

拼接多个字符串

示例：
~~~javascript
String.concat("") // => ""
String.concat("", "a", "b", "c") // => "abc"
String.concat("ABC") // => "ABC"
String.concat("ABC", "a", "b", "c") // => "ABCabc"
~~~

### String.endsWith
> String.endsWith(s string, searchString string\[, endPosition int]) bool

检查字符串是否以某个子串结尾

示例：
~~~javascript
String.endsWith("abc", "c") // => true
String.endsWith("abc", "b") // => false
String.endsWith("abc", "c", -1) // => false
String.endsWith("abc", "c", 3) // => true
String.endsWith("abc", "c", 4) // => true
String.endsWith("abc", "b", 2) // => true
String.endsWith("abc", "a", 1) // => true
String.endsWith("abc", "", 0) // => true
~~~

### String.includes
> String.includes(s string, searchString string\[, position int]) bool

检查字符串是否包含某个子串

示例：
~~~javascript
String.includes("abcdefg", "cde") // => true
String.includes("abcdefg", "bde") // => false
String.includes("abcdefg", "cde", 6)) // => true
String.includes("abcdefg", "cde", 3) // => false
~~~

### String.indexOf
> String.indexOf(s string, searchString string\[, position int]) int

查找子串首次出现位置

如果未找到则返回 -1

示例：
~~~javascript
String.indexOf("abc", "defg") // => -1
String.indexOf("abcdefg", "abc") // => 0
String.indexOf("abcdefg", "bcd") // => 1
String.indexOf("abcdefg", "") // => 0
String.indexOf("abcdefg", "", 1) // => 1
String.indexOf("abcdefg", "cd", 1) // => 2
String.indexOf("abcdefg", "cd", 3) // => -1
String.indexOf("abcdefg", "cd", 10) // => -1
~~~

### String.lastIndexOf
> String.lastIndexOf(s string, searchString string\[, position int]) int

查找子串最后一次出现的位置

如果未找到则返回 -1

示例：
~~~javascript
String.lastIndexOf("abcdefg", "ff") // => -1
String.lastIndexOf("abcdefg", "abc") // => 0
String.lastIndexOf("abcdefg", "fg") // => 5
String.lastIndexOf("abcdefg", "fg", 5) // => 5
String.lastIndexOf("abcdefg", "fg", 6) // => 5
String.lastIndexOf("abcdefg", "fg", 10) // => 5
String.lastIndexOf("abcdefg", "") // => 7
String.lastIndexOf("abcdefg", "", 3) // => 3
String.lastIndexOf("abcdefg", "", 10) // => 7
String.lastIndexOf("abcdefg", "", -1) // => -1
~~~

### String.match
> String.match(s string, regexp any) \[]string

匹配正则

返回匹配的片段

示例：
~~~javascript
String.match("abc", "\\w+") // => ["abc"]
String.match("ab|c", "\\w+") // => ["ab"]
String.match("()*&)", "\\w+") // => nil
String.match("abc", NewRegExp("\\w+")) // => ["abc"]
String.match("ab|c", NewRegExp("\\w+")) // => ["ab"]
~~~

### String.repeat
> String.repeat(s string, count int) string

重复字符串

示例：
~~~javascript
String.repeat("abc", 0) // => ""
String.repeat("abc", -1) // => ""
String.repeat("abc", 10) // => "abcabcabcabcabcabcabcabcabcabc"
~~~

### String.replace
> String.replace(s string, searchValue string, replaceValue string) string

替换首个匹配子串

示例：
~~~javascript
String.replace("abcdefg", "", "") // => "abcdefg"
String.replace("abcdefg", "a", "1") // => "1bcdefg"
String.replace("abcdefg", "b", "1") // => "a1cdefg"
String.replace("abcabc", "b", "1") // => "a1cabc"
~~~

### String.replaceAll
> String.replaceAll(s string, searchValue string, replaceValue string) string

替换全部匹配子串

示例：
~~~javascript
String.replaceAll("abcabc", "b", "1") // => "a1ca1c"
~~~

### String.padEnd
> String.padEnd(s string, maxLength int, fillString string) string

从尾部填充字符串至指定长度

示例：
~~~javascript
String.padEnd("abc", -1, "ABC") // => "abc"
String.padEnd("abc", 0, "ABC") // => "abc"
String.padEnd("abc", 3, "ABC") // => "abc"
String.padEnd("abc", 4, "ABC") // => "abcA"
String.padEnd("abc", 6, "ABC") // => "abcABC"
String.padEnd("abc", 9, "ABC") // => "abcABCABC"
String.padEnd("abc", 14, "ABC") // => "abcABCABCABCAB"
String.padEnd("abc", 14, "") // => "abc"
String.padEnd("abc", 14, " ") // => "abc           "
~~~

### String.padStart
> String.padStart(s string, maxLength int, fillString string) string

从首部填充至指定长度

示例：
~~~javascript
String.padStart("abc", -1, "ABC") // => "abc"
String.padStart("abc", 0, "ABC") // => "abc"
String.padStart("abc", 3, "ABC") // => "abc"
String.padStart("abc", 4, "ABC") // => "Aabc"
String.padStart("abc", 6, "ABC") // => "ABCabc"
String.padStart("abc", 9, "ABC") // => "ABCABCabc"
String.padStart("abc", 14, "ABC") // => "ABCABCABCABabc"
String.padStart("abc", 14, "") // => "abc"
String.padStart("abc", 14, " ") // => "           abc"
~~~

### String.slice
> String.slice(s string, start int, end int) string

截取字符串一部分

示例：
~~~javascript
String.slice("", 0, 0) // => ""
String.slice("abc", -3, -2) // => "a"
String.slice("abc", -3, -1) // => "ab"
String.slice("abc", 0, 10) // => "abc"
String.slice("abc", 1, 3) // => "bc"
String.slice("abc", 1, 1) // => ""
String.slice("abc", 1, 2) // => "b"
String.slice("abcdefg", -3, 7) // => "efg"
String.slice("abcdefg", 10, 100) // => ""
~~~

### String.split
> String.split(s string, separator string\[, limit int]) \[]string

使用分隔符分割字符串

示例：
~~~javascript
String.split("", "") // => []
String.split("abc", "") // => ["a", "b", "c"]
String.split("abc", "b") // => ["a", "c"]
String.split("abc", "b", -1) // => ["a", "c"]
String.split("abc", "b", 1) // => ["a"]
String.split("abc", "b", 2) // => ["a", "c"]
String.split("abc", "d", 2) // => ["abc"]
~~~

### String.startsWith
> String.startsWith(s string, searchString string\[, startPosition int]) bool

检查字符串是否以某个子串开头

示例：
~~~javascript
String.startsWith("abc", "a") // => true
String.startsWith("abc", "ab") // => true
String.startsWith("abc", "b") // => false
String.startsWith("abc", "a", -1) // => true
String.startsWith("abc", "a", 0) // => true
String.startsWith("abc", "b", 1) // => true
String.startsWith("abcdefg", "ef", 4) // => true
String.startsWith("abc", "", 1) // => true
String.startsWith("abc", "", 0) // => true
String.startsWith("abc", "", -1) // => true
~~~

### String.substring
> String.substring(s string, start int, end int) string

获取字符串一部分

示例：
~~~javascript
String.substring("", 1, 2) // => ""
String.substring("abcdefg", 1, 3) // => "bc"
String.substring("abcdefg", 0, 3) // => "abc"
String.substring("abcdefg", -1, 3) // => "abc"
String.substring("abcdefg", 3, 0) // => "abc"
String.substring("abcdefg", 3, 1) // => "bc"
String.substring("abcdefg", 0, 100) // => "abcdefg"
String.substring("abcdefg", 6, 100) // "g"
String.substring("abcdefg", 7, 100) // => ""
String.substring("abcdefg", 10, 100) // => ""
~~~

### String.toLowerCase
> String.toLowerCase(s string) string

将字符串转换为小写

示例：
~~~javascript
String.toLowerCase("") // => ""
String.toLowerCase("abc") // => "abc"
String.toLowerCase("ABC") // => "abc"
String.toLowerCase("aBC") // => "abc"
String.toLowerCase("ABC中文") // => "abc中文"
~~~

### String.toUpperCase
> String.toUpperCase(s string) string

将字符串转换为大写

示例：
~~~javascript
String.toUpperCase("ABC") // => ABC
String.toUpperCase("abc") // => ABC
String.toUpperCase("Abc") // => ABC
String.toUpperCase("abc中文") // => ABC中文
~~~

### String.trim
> String.trim(s string) string

去除字符串首尾空白

示例：
~~~javascript
String.trim("") // => ""
String.trim(" a ") // => "a"
String.trim(" a b \t\n\r\t") // => "a b"
~~~

### String.trimEnd
> String.trimEnd(s string) string

去除字符串尾部空白

示例：
~~~javascript
String.trimEnd("") // => ""
String.trimEnd(" a ") // => " a"
String.trimEnd(" a b \t\n\r\t") // => " a b"
~~~

### String.trimPrefix
> String.trimPrefix(s string, prefix string) string

去除字符串前缀

示例：
~~~javascript
String.trimPrefix("abcdefg", "ab") // => "cdefg"
String.trimPrefix("abcdefg", "") // => "abcdefg"
String.trimPrefix("ababcdefg", "ab") // => "abcdefg"
~~~

### String.trimStart
> String.trimStart(s string) string

去除字符串首部空白

示例：
~~~javascript
String.trimStart("") // => ""
String.trimStart(" a ") // => "a "
String.trimStart(" \t\n\r\t a b") // => "a b"
~~~

### String.trimSuffix
> String.trimSuffix(s string, suffix string) string

去除字符串后缀

示例：
~~~javascript
String.trimSuffix("abcdefg", "fg") // => "abcde"
String.trimSuffix("abcdefg", "") // => "abcdefg"
String.trimSuffix("abcdefgfg", "fg") // => "abcdefg"
~~~

### String.length
> String.length(s string) int

获取字符串长度

示例：
~~~javascript
String.length("") // => 0
String.length("abc") // => 3
String.length("abc中文") // => 5
~~~

### String.sprintf
> String.sprintf(s string, args ...any) string

格式化字符串

示例：
~~~javascript
String.sprintf("%s %d, %.2f", "abc", 1, 2.3456) // => "abc 1, 2.35"
~~~

## Math
> Math

数学相关操作

### Math.E
> Math.E

自然常数 e

### Math.LN10
> Math.LN10

ln(10)

### Math.LOG10E
> Math.LOG10E

log10(e)

### Math.PI
> Math.PI

圆周率 π

### Math.MathSQRT1_2
> Math.MathSQRT1_2

1/√2

### Math.SQRT2
> Math.SQRT2

√2

### Math.abs
> Math.abs(x float64) float64

计算绝对值

示例：
~~~javascript
Math.abs(2.0) // => 2.0
Math.abs(-2.123) // => 2.123
~~~

### Math.cbrt
> Math.cbrt(x float64) float64

计算立方根

示例：
~~~javascript
Math.cbrt(8.0) // => 2.0
Math.cbrt(-27.0) // => -3.0
~~~

### Math.ceil
> Math.ceil(x float64) int64

向上取整

示例：
~~~javascript
Math.ceil(1.0) // => 1
Math.ceil(1.0234) // => 2
Math.ceil(1.5) // => 2
Math.ceil(123.678) // => 124
~~~

### Math.clz32
> Math.clz32(x uint32) int

计算uint32数值的二进制前导零数量

示例：
~~~javascript
Math.clz32(1) // => 31
Math.clz32(4) // => 29
Math.clz32(0) // => 32
Math.clz32(2147483648) // => 0
~~~

### Math.cos
> Math.cos(x float64) float64

计算余弦值

示例：
~~~javascript
Math.cos(Math.PI/3) // => 0.49999999999999994
~~~

### Math.cosh
> Math.cosh(x float64) float64

计算双曲余弦值

示例：
~~~javascript
Math.cosh(30) // => 5.343237290762231e+12
Math.cosh(60) // => 5.710036949078421e+25
~~~

### Math.exp
> Math.exp(x float64) float64

计算 e^x

示例：
~~~javascript
Math.exp(100.0) // => 2.6881171418161356e+43
~~~

### Math.expm1
> Math.expm1(x float64) float64

计算 e^x - 1

示例：
~~~javascript
Math.expm1(100.0) // => 2.6881171418161356e+43
~~~

### Math.floor
> Math.floor(x float64) int64

向下取整

示例：
~~~javascript
Math.floor(1.0) // => 1
Math.floor(1.0234) // => 1
Math.floor(123.678) // => 123
~~~

### Math.hypot
> Math.hypot(args ...float64) float64

计算平方和开方

示例：
~~~javascript
Math.hypot(3, 4) // => 5.0
Math.hypot(-3, -4) // => 5.0
Math.hypot(0, 0) // => 0.0
~~~

### Math.log
> Math.log(x float64) float64

计算自然对数 ln(x)

示例：
~~~javascript
Math.log(2) // => 0.6931471805599453
~~~

### Math.log1p
> Math.log1p(x float64) float64

计算 ln(1+x)

示例：
~~~javascript
Math.log1p(2) // => 1.0986122886681096
~~~

### Math.log10
> Math.log10(x float64) float64

计算常用对数 log10(x)

示例：
~~~javascript
Math.log10(2) // => 0.3010299956639812
~~~

### Math.log2
> Math.log2(x float64) float64

计算以2为底的对数

示例：
~~~javascript
Math.log2(2) // => 1
Math.log2(10) // => 3.321928094887362
~~~

### Math.pow
> Math.pow(base float64, exponent float64) float64

幂运算

示例：
~~~javascript
Math.pow(2.0, 0.0) // => 1.0
Math.pow(4.0, 0.5 // => 2.0
isNaN(Math.pow(-4.0, 0.5)) // => true
~~~

### Math.random
> Math.random() float64

生成随机浮点数

范围在 [0,1) 之间

示例：
~~~javascript
Math.random() // => 0.6826845041002553 (类似的小数，非固定值)
~~~

### Math.randN
> Math.randN(n int) int

生成随机整数

范围在 [0, n) 之间

示例：
~~~javascript
Math.randN(10) // => 5 （类似的整数，在0-9之间（不含10），非固定值）
Math.randN(0)) // => 0
Math.randN(-10) // => 0
~~~

### Math.round
> Math.round(x float64) float64

四舍五入到最近浮点数形式的整数

示例：
~~~javascript
Math.round(2.0) // => 2.0
Math.round(2.0123) // => 2.0
Math.round(2.6123) // => 3.0
~~~

### Math.sign
> Math.sign(x float64) float64

根据数值正负生成符号数

返回-1.0、0.0、1.0

示例：
~~~javascript
Math.sign(1.0) // => 1.0
Math.sign(-123456.123) // => -1.0
Math.sign(0) // => 0.0
~~~

### Math.sin
> Math.sin(x float64) float64

计算正弦值

示例：
~~~javascript
Math.sin(Math.PI/6) // => 0.49999999999999994
~~~

### Math.sinh
> Math.sinh(x float64) float64

计算双曲正弦值

示例：
~~~javascript
Math.sinh(30) // => 5.343237290762231e+12
Math.sinh(60) // => 5.710036949078421e+25
~~~

### Math.sqrt
> Math.sqrt(x float64) float64

计算平方根

示例：
~~~javascript
Math.sqrt(4) // => 2.0
Math.sqrt(9) // => 3.0
~~~

### Math.tan
> Math.tan(x float64) float64

计算正切值

示例：
~~~javascript
Math.tan(math.Pi / 4) // => 0.9999999999999998
~~~

### Math.tanh
> Math.tanh(x float64) float64

计算双曲正切值

示例：
~~~javascript
Math.tanh(45) // => 1
~~~

### Math.trunc
> Math.trunc(x float64) int64

向零截断取整

示例：
~~~javascript
Math.trunc(2.0) // => 2
Math.trunc(2.0123) // => 2
Math.trunc(34.6789) // => 34
~~~

## Date
> Date

日期相关操作

### Date.new
> Date.new() Date

构造当前时间的日期对象

示例：
~~~javascript
let d = Date.new();
let year = string(d.getFullYear());
let month = String.padStart(string(d.getMonth()+1), 2, '0');
let day = String.padStart(string(d.getDate()), 2, '0');
year + "-" + month + "-" + day // => 类似于 "2026-04-28" 的当前日期
~~~

## NewDate
> NewDate() Date

构造日期对象

同 Date.new()

## RegExp
> RegExp

正则表达式相关操作

### RegExp.escape
> RegExp.escape(s string) string

转义用于正则表达式的特殊字符

示例：
~~~javascript
RegExp.escape("#$%*()-[]{}|") // => "#\$%\*\(\)-\[\]\{\}\|"
~~~

### RegExp.new
> RegExp.new(expr string) RegExp

构造新正则表达式

示例：
~~~javascript
RegExp.new("\\w+").test("abc") // => true
~~~

## NewRegExp
> NewRegExp(expr string) RegExp

构造新正则表达式

同 RegExp.New(expr)

示例：
~~~javascript
NewRegExp("\\w+").test("abc") // => true
~~~

## JSON
> JSON

JSON相关操作

### JSON.parse
> JSON.parse(text string) any

解析JSON文本

示例：
~~~javascript
let v = JSON.parse('{ "a": 1, "b": 2 }');
v.a // => 1
~~~

### JSON.stringify
> JSON.stringify(value any) string

将数据序列化为JSON文本

示例：
~~~javascript
JSON.stringify({"a":1, "b": 2} // => "{"a":1,"b":2}"
~~~

## NewURL
> NewURL(url string) URL

构造新URL对象

示例：
~~~javascript
let u = NewURL("https://user:pass@example.com:8080/docs?nav=1#link");
u.port // => 8080
u.hash // => "link"
u.host // => "example.com:8080"
u.query // => "nav=1"
u.scheme // => "https"
u.path // => "/docs"
u.opaque // => ""
u.user // => {"password":"pass", "username":"user"}
~~~

## Base64
> Base64

Base64相关操作

### Base64.encode
> Base64.encode(s string) string

Base64编码

示例：
~~~javascript
Base64.encode("Hello, World!") // => "SGVsbG8sIFdvcmxkIQ=="
~~~

### Base64.decode
> Base64.decode(s string) string

Base64 解码

示例：
~~~javascript
Base64.decode("SGVsbG8sIFdvcmxkIQ==") // => "Hello, World!"
~~~

## NetIP
> NetIP

IP地址相关操作

### NetIP.isValid
> NetIP.isValid(ip string) bool

判断IP地址是否有效

示例：
~~~javascript
NetIP.isValid("127.0.0.1") // => true
NetIP.isValid("::1") // => true
NetIP.isValid("127.0.0.1.1") // => false
NetIP.isValid("::1.1") // => false
NetIP.isValid("") // => false
NetIP.isValid("127.0.0.256") // => false
~~~

### NetIP.isIPv4
> NetIP.isIPv4(ip string) bool

判断IP地址是否为IPv4地址

示例：
~~~javascript
NetIP.isIPv4("127.0.0.1") // => true
NetIP.isIPv4("::1") // => false
~~~

### NetIP.isIPv6
> NetIP.isIPv6(ip string) bool

判断IP地址是否为IPv6地址

示例：
~~~javascript
NetIP.isIPv6("::1") // => true
NetIP.isIPv6("127.0.0.1") // => false
~~~

### NetIP.isBetween
> NetIP.isBetween(ip string, startIP string, endIP string) bool

判断IP地址是否在指定的范围内

示例：
~~~javascript
NetIP.isBetween("127.0.0.1", "127.0.0.1", "127.0.0.2") // => true
NetIP.isBetween("127.0.0.1", "127.0.0.2", "127.0.0.3") // => false
NetIP.isBetween("127.0.1.2", "127.0.0.2", "127.0.2.3") // => true
~~~

### NetIP.isInCIDR
> NetIP.isInCIDR(ip string, cidr string) bool

判断IP地址是否在指定的CIDR范围内

示例：
~~~javascript
NetIP.isInCIDR("127.0.0.1", "127.0.0.0/24") // => true
NetIP.isInCIDR("127.0.0.1", "127.0.0.0/16") // => true
NetIP.isInCIDR("127.0.1.2", "127.0.0.0/8") // => true
NetIP.isInCIDR("127.0.1.1", "127.0.0.0/32") // => false
~~~

### NetIP.isInRanges
> NetIP.isInRanges(ip string, ranges \[]\[2]string) bool

判断IP地址是否在指定的范围内

示例：
~~~javascript
NetIP.isInRanges("127.0.0.1", [["127.0.0.0", "127.0.0.255"]]) // => true
NetIP.isInRanges("127.0.0.1", [["127.0.0.0", "127.0.0.127"]]) // => true
NetIP.isInRanges("127.0.1.2", [
		["127.0.0.0", "127.0.0.255"],
		["127.0.1.0", "127.0.1.255"],
		["127.0.2.0", "127.0.2.255"]
]) // => true
~~~

## Crypto
> Crypto

加密相关操作

### Crypto.NewHMAC
> Crypto.NewHMAC(algorithm string, key string) (CryptoHMACHash, error)

创建HMAC摘要器

algorithm 支持md5、sha1、sha256、sha512

示例：
~~~javascript
let h = Crypto.NewHMAC("sha1", "123456");
h.update("ABCDEFG");
h.sum() // => "02ce4f53c007cb7c90f39b953e1fe71a3c6a178d"
~~~

## console
> console

控制台相关操作

### console.log
> console.log(data ...any) bool

输出日志

示例：
~~~javascript
console.log("Hello", "World")
~~~

## Request对象
### [object].id
> [object].id() string

获取当前请求ID

### [object].serverInfo
> [object].serverInfo() RequestServerInfo

获取网站信息

### [object].nodeInfo
> [object].nodeInfo() RequestNodeInfo

获取节点信息

### [object].url
> [object].url() string

获取当前URL

### [object].path
> [object].path() string

获取当前URL路径

### [object].uri
> [object].uri() string

获取当前URI

一般为路径加参数

### [object].setURI
> [object].setURI(uri string) bool

设置当前回源的URI

### [object].query
> [object].query() URLQuery

读取当前查询参数对象

### [object].host
> [object].host() string

读取当前主机地址

如果正在访问的地址使用非标准端口号，那么包含端口

### [object].remoteAddr
> [object].remoteAddr() string

读取客户端地址

### [object].rawRemoteAddr
> [object].rawRemoteAddr() string

读取直接连接的客户端地址

如果用户使用了代理，那么将会读取到代理的地址

### [object].remotePort
> [object].remotePort() int

读取客户端端口

### [object].method
> [object].method() string

读取请求方法

返回如GET、POST、HEAD之类的请求方法

### [object].contentLength
> [object].contentLength() int64

请求内容长度

即客户端发送的请求内容长度，GET方法请求的内容通常长度为0

### [object].transferEncoding
> [object].transferEncoding() string

请求使用的编码

### [object].proto
> [object].proto() string

请求使用的协议

比如 HTTP/1.1

### [object].protoMajor
> [object].protoMajor() int

请求使用的协议主版本

比如 HTTP/1.0 中的 1

### [object].protoMinor
> [object].protoMinor() int

请求使用的协议小版本

比如 HTTP/1.0 中的 0

### [object].cookie
> [object].cookie(name string) string

读取请求发送的Cookie值

### [object].header
> [object].header() HTTPHeader

读取所有请求发送的报头

### [object].setHeader
> [object].setHeader(name string, values ...string) bool

设置请求报头

用来修改发送到源站的请求报头

### [object].deleteHeader
> [object].deleteHeader(name string) bool

删除请求报头

用来修改发送到源站的请求报头

### [object].setAttr
> [object].setAttr(name string, value string) bool

设置请求属性

可以记录到访问日志的额外信息

### [object].setVar
> [object].setVar(name string, value string) bool

设置自定义变量

以便可以在 $req.format(name) 函数中使用

### [object].format
> [object].format(format string) string

根据请求信息格式化内容

在FlexCDN中，可以使用请求变量，以下是几个例子：
~~~javascript
$req.format("${requestId}") // => 请求ID
$req.format("${geo.country.name}") // => 类似于“中国”
$req.format("${geo.province.name}") // => 类似于“湖北省”
~~~

### [object].done
> [object].done() bool

完成请求

不再继续执行

### [object].close
> [object].close() bool

关闭请求

### [object].allow
> [object].allow() bool

允许访问

即将当前请求设置为白名单直接通过WAF检查

## Response对象
### [object].setHeader
> [object].setHeader(name string, values ...string) bool

设置响应报头

### [object].deleteHeader
> [object].deleteHeader(name string) bool

删除响应报头

### [object].header
> [object].header() HTTPHeader

读取所有响应报头

### [object].send
> [object].send(status int, body string) bool

发送内容

- status 为状态码，常用200

### [object].sendResp
> [object].sendResp(resp *http.Response) (int64, error)

发送响应对象

### [object].redirect
> [object].redirect(status int, url string) bool

跳转到某个URL

- status 为跳转状态码，如307等

## CryptoHMACHash对象
### [object].update
> [object].update(data string) bool

更新数据

### [object].sum
> [object].sum() string

计算摘要

## Date对象
### [object].getDate
> [object].getDate() int

获取日期

1-31

### [object].getDay
> [object].getDay() int

获取一周中的天

0-6

### [object].getFullYear
> [object].getFullYear() int

获取年份

类似于 2006

### [object].getHours
> [object].getHours() int

获取24制小时数

类似于 1、11、16等

### [object].getMilliseconds
> [object].getMilliseconds() int

获取当前时间戳毫秒部分

比如 150、320 等

### [object].getMinutes
> [object].getMinutes() int

获取当前分钟数

类似于 1、15、45

### [object].getMonth
> [object].getMonth() int

获取当前月数

值为 0-11

### [object].getSeconds
> [object].getSeconds() int

获取当前秒数

类似于 1、10、35 等

### [object].getTime
> [object].getTime() int64

获取当前时间戳

含毫秒

类似于 1777369354150

## RegExp对象
### [object].test
> [object].test(s string) bool

测试某个字符串是否匹配当前正则表达式
示例：
~~~javascript
NewRegExp("\\w+").test("abc") // => true
~~~

### [object].exec
> [object].exec(s string) \[]string

执行当前正则表达式匹配
返回匹配的内容，圆括号视为匹配的子串
~~~javascript
NewRegExp("\\d{3}").Exec("123456") // => [123]
NewRegExp("\\d+").Exec("123456") // => [123456]
NewRegExp("\\d+").Exec("123|456") // => [123]
NewRegExp("(\\d)(\\d)").Exec("123|456") // => [12, 1, 2]
~~~

### [object].split
> [object].split(s string) \[]string

使用当前正则表达式分割字符串
~~~javascript
NewRegExp("A").Split("123A456A789")  // => ["123", "456", "789"]
NewRegExp("A").Split("123456789") // => ["123456789"]
NewRegExp("\\|").Split("123|456|789") // => ["123", "456", "789"]
~~~

## HTTPHeader对象
### [object].add
> [object].add(key string, value string)

添加报头

### [object].set
> [object].set(key, value string)

设置报头

### [object].get
> [object].get(key string) string

读取报头值

示例：
~~~javascript
$resp.header().get("User-Agent") // => Mozilla/5.0 ... Safari/537.36
~~~

### [object].values
> [object].values(key string) \[]string

读取报头所有值

### [object].delete
> [object].delete(key string)

删除报头

### [object].has
> [object].has(key string) bool

判断是否包含某个报头

### [object].toJSON
> [object].toJSON() string

将报头转换为JSON

## RequestNodeInfo对象
### [object].id
> [object].id

节点ID

@expr id

## RequestServerInfo对象
### [object].id
> [object].id

网站ID

@expr id

## URL对象
### [object].host
> [object].host

主机名

如果有端口号的话，此值也包含端口号

### [object].query
> [object].query

查询参数

不包括开始的问号

### [object].port
> [object].port

端口号

### [object].path
> [object].path

路径

从正斜杠开始

### [object].hash
> [object].hash

锚点

不包括井号（#）部分

### [object].scheme
> [object].scheme

协议

### [object].opaque
> [object].opaque

非透明数据

比如 `mailto:user@example.com` 中的 `opaque` 为 `user@example.com`

### [object].user
> [object].user

用户信息内容

## URLQuery对象
### [object].get
> [object].get(name string) string

Get 获取某个参数值

### [object].values
> [object].values(name string) \[]string

Values 获取某个参数的所有值

在一个参数有很多值的时候很有用

示例：
对于：
~~~
https://example.com?name=lily&name=lucy&name=jim
~~~
来说：
~~~javascript
$req.query().values("name") // => ["lily", "lucy", "jim"]
~~~

### [object].has
> [object].has(name string) bool

Has 判断某个参数值是否存在

### [object].set
> [object].set(name string, value string) bool

Set 设置参数值

### [object].delete
> [object].delete(name string) bool

Delete 删除某个参数值

### [object].encode
> [object].encode() string

Encode 将所有参数值编码为一个字符串

### [object].toJSON
> [object].toJSON() string

ToJSON 将所有参数值转为为JSON

其中每个参数值对应一个字符串数组

### [object].toKV
> [object].toKV() map\[string]string

ToKV 将所有参数值转换为键值对

### [object].ToKVJSON
> [object].toKVJSON() string

ToKVJSON 将所有参数值转换为键值对JSON

其中每个参数值对应一个字符串

### [object].toMap
> [object].toMap() map\[string]\[]string

ToMap 将所有参数值转换为键值对Map

