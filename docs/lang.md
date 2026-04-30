# 语言参考

> FlexLang基于`ExprLang`实现，所以可以使用`ExprLang`的绝大部分特性。
>
> 这篇文档大幅参考了 ExprLang：https://expr-lang.org/docs/language-definition

## 数据类型

### 布尔类型（Boolean）

`true` 或 `false`。

### 整数类型（Integer）

支持十进制、二进制、八进制、十六进制等，比如 `42`， `0x2A`， `0o52`， `0b101010`。

### 浮点数类型（Float）

比如 `3.1415926`、`0.5`、`.5`。

### 字符串类型（String）

支持单双引号，比如 `"foo"`、`'bar'`。

多行字符串，可以使用反引号（`）：

~~~
`Hello
World`
~~~

比如在函数调用中使用反引号：

~~~
$resp.send(200, `Hello, 
  World!`)
~~~

### 数组（Array）

比如 `[1, 2, 3]`。

### 映射（Map）

比如 `{a: 1, b: 2, c: 3}`。

### 空值（Nil）

只有一个空值 `nil`。

### 字节序列（Bytes）

比如`b"hello"`， `b'\xff\x00'`。

更多示例：

~~~
b"abc"      // => []byte{97, 98, 99}
b"ÿ"        // => []byte{195, 191} - UTF-8 encoding of ÿ
b"\xff"     // => []byte{255}
b"\x00\x01" // => []byte{0, 1}
b"\101"     // => []byte{65} - octal for 'A'
~~~

## 运算符

<table>
    <tr>
        <td><strong>算数运算符（Arithmetic）</strong></td>
        <td>
            <code>+</code>, <code>-</code>, <code>*</code>, <code>/</code>, <code>%</code> (取余), <code>^</code> 或 <code>**</code> (指数)
        </td>
    </tr>
    <tr>
        <td><strong>比较运算符（Comparison）</strong></td>
        <td>
            <code>==</code>, <code>!=</code>, <code>&lt;</code>, <code>&gt;</code>, <code>&lt;=</code>, <code>&gt;=</code>
        </td>
    </tr>
    <tr>
        <td><strong>逻辑运算符（Logical）</strong></td>
        <td>
            <code>not</code> 或 <code>!</code>, <code>and</code> 或 <code>&amp;&amp;</code>, <code>or</code> 或 <code>||</code>
        </td>
    </tr>
    <tr>
        <td><strong>条件运算符（Conditional）</strong></td>
        <td>
            <code>?:</code> (三元运算符), <code>??</code> (空值合并运算符), <code>if {} else {}</code> (多行)
        </td>
    </tr>
    <tr>
        <td><strong>成员运算符（Membership）</strong></td>
        <td>
            <code>[]</code>, <code>.</code>, <code>?.</code>, <code>in</code>
        </td>
    </tr>
    <tr>
        <td><strong>字符串操作（String）</strong></td>
        <td>
            <code>+</code> (字符串拼接运算符), <code>contains</code>, <code>startsWith</code>, <code>endsWith</code>
        </td>
    </tr>
    <tr>
        <td><strong>正则表达式（Regex）</strong></td>
        <td>
            <code>matches</code>
        </td>
    </tr>
    <tr>
        <td><strong>范围运算符（Range）</strong></td>
        <td>
            <code>..</code>
        </td>
    </tr>
    <tr>
        <td><strong>切片运算符（Slice）</strong></td>
        <td>
            <code>[:]</code>
        </td>
    </tr>
    <tr>
        <td><strong>管道运算符（Pipe）</strong></td>
        <td>
            <code>|</code>
        </td>
    </tr>
</table>

### 成员运算符

可以使用点（`.`）或 `[]`来访问一个对象的成员：

~~~javascript
user.Name
user["Name"]
~~~

可以使用`[]`运算符来访问数组和切片的元素：

~~~javascript
array[0] // 第一个元素
~~~

也支持负索引（实际访问的索引是`数组长度+索引`）：

~~~javascript
array[-1] // 最后一个元素
~~~

`in`操作符可以用来检查一个元素是否在数组或映射中：

~~~javascript
"John" in ["John", "Jane"]
"name" in {"name": "John", "age": 30}
~~~

#### 可选调用链

可以使用 `?.` 操作符访问一个可能为空值的数据，如果数据不存在，则直接返回`nil`，不报告错误：

~~~javascript
author.User?.Name
~~~

等价于：

~~~javascript
author.User != nil ? author.User.Name : nil
~~~

#### 空值合并运算符

`??`运算符可以判断数据是否不为空，如果不为空，则返回左侧数据，否则返回右侧数据：

~~~javascript
author.User?.Name ?? "Anonymous"
~~~

等价于：

~~~javascript
author.User != nil ? author.User.Name : "Anonymous"
~~~

### 切片运算符

切片运算符`[:]`可以用来访问数组的某个片段。

以数组 `[1, 2, 3, 4, 5]` 为例：

~~~expr
array[1:4] == [2, 3, 4]
array[1:-1] == [2, 3, 4]
array[:3] == [1, 2, 3]
array[3:] == [4, 5]
array[:] == array
~~~

### 管道运算符

孤雁到运算符可以将左侧表达式结果作为第一个参数传递给右侧表达式：

~~~javascript
user.Name | lower() | split(" ")
~~~

等价于：

~~~javascript
split(lower(user.Name), " ")
~~~

### 范围运算符

范围运算符`..`可以用于生成一定范围的整数数组：

~~~expr
1..3 == [1, 2, 3]
~~~

## 变量

可以使用 `let` 关键词定义变量，变量名可以包含字母、数字和下划线：

~~~javascript
let x = 42;
x * 2
~~~

多个变量可以使用分号分割：

~~~javascript
let x = 42;
let y = 2;
x * y
~~~

变量中也可以使用管道运算符（`|`）：

~~~javascript
let name = user.Name | lower() | split(" ");
"Hello, " + name[0] + "!"
~~~

## 断言

断言（`Predicate`）是一个表达式。断言可用于 `filter`（过滤）、`all`（全匹配）`any`、`one`、`none`等函数中。

下面例子是创建一个从0到9的数组，然后过滤偶数：

~~~expr
filter(0..9, {# % 2 == 0})
~~~

如果数组内容是映射，可以忽略`#`符号（`#.Value`变成`.Value`）：

~~~expr
filter(tweets, {len(.Content) > 240})
~~~

花括号（`{`、`}`）也可以省略：

~~~expr
filter(tweets, len(.Content) > 240)
~~~

对于嵌套的断言，可以使用变量代替`#`，避免混淆：

~~~expr
filter(posts, {
    let post = #; 
    any(.Comments, .Author == post.Author)
})
~~~

## 注释

可以使用 `/* */` 或 `//`：

~~~javascript
/**
 * 这是注释1
 */
1 + 1 // 这是注释2
~~~

## 日期（Date）函数

Expr 内置支持 Go 语言中的 [time.Time]((https://pkg.go.dev/time))。你可以通过两个日期相减来获取它们之间的时间间隔（Duration）：

```javascript
createdAt - now()
```

可以将一段时间间隔（Duration）加到日期上：

```javascript
createdAt + duration("1h")
```

并且可以对日期进行比较：

```javascript
createdAt > now() - duration("1h")
```

### now()

返回一个Go语言中的 [time.Time](https://pkg.go.dev/time#Time) 值。

```javascript
now().Year() == 2024
```

### duration(str)

根据给定的一个字符串 `str` 返回一个Go语言中的 [time.Duration](https://pkg.go.dev/time#Duration).

有效的时间单位为 "ns", "us" (or "µs"), "ms", "s", "m", "h"。

```javascript
duration("1h").Seconds() == 3600
```

### date(str[, format[, timezone]])

将给定的字符串`str`转换为日期对象形式。

如果未提供 `format` 参数，则参数 v 必须符合以下格式之一：

- 2006-01-02
- 15:04:05
- 2006-01-02 15:04:05
- RFC3339
- RFC822,
- RFC850,
- RFC1123,

```javascript
date("2023-08-14")
date("15:04:05")
date("2023-08-14T00:00:00Z")
date("2023-08-14 00:00:00", "2006-01-02 15:04:05", "Europe/Zurich")
```

可用的方法如下：

- `Year()` - 返回年份
- `Month()` - 返回月份（从1开始）
- `Day()` - 返回一个月中的日期
- `Hour()` - 返回小时（24小时制）
- `Minute()` - 返回分钟
- `Second()` - 返回秒数
- `Weekday()` - 返回一周中的天（0-6）
- `YearDay()` - 返回一年中的天
- [更多](https://pkg.go.dev/time#Time).

```javascript
date("2023-08-14").Year() == 2023
```

### timezone(str)

根据给定的字符串 `str`。

返回时区对象。可用的时区可以 [查看这里](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones).

```javascript
timezone("Europe/Zurich")
timezone("UTC")
```

为了转换日期到不同的时区，可以使用 [`In()`](https://pkg.go.dev/time#Time.In)方法:

```javascript
date("2023-08-14 00:00:00").In(timezone("Europe/Zurich"))
```

## 数值类型函数

### max(n1, n2)

返回`n1`、`n2`两者之间的最大值。

```javascript
max(5, 7) == 7
```

### min(n1, n2)

返回`n1`、`n2`两者之间的最小值。

```javascript
min(5, 7) == 5
```

### abs(n)

获取`n`的绝对值。

```javascript
abs(-5) == 5
```

### ceil(n)

返回大于或等于`n`的最小整数：

```javascript
ceil(1.5) == 2.0
```

### floor(n)

返回小于或等于`n`的最大整数：

```javascript
floor(1.5) == 1.0
```

### round(n)

返回`n`四舍五入后的整数：

```expr
round(1.5) == 2.0
```

## 数组函数（Array Functions）

### all(array, predicate)

如果所有的元素都满足 [断言](#断言)，则返回 **true**。

如果数组是空数组，也会返回 **true**。

```expr
all(tweets, {.Size < 280})
```

### any(array, predicate)

如果任一元素满足 [断言](#断言)，则返回 **true**。

如果元素为空，则返回 **false**。

```expr
any(tweets, {.Size > 280})
```

### one(array, predicate)

如果有且仅有一个元素满足 [断言](#断言)，则返回 **true**。

如果元素为空，则返回 **false**。

```expr
one(participants, {.Winner})
```

### none(array, predicate)

如果数组中所有元素都不满足 [断言](#断言)，则返回 **true**。

如果元素为空， 也返回 **true**。

```expr
none(tweets, {.Size > 280})
```

### map(array, predicate)

数组中每一个元素都应用 [断言](#断言)，返回一个新数组。

```expr
map(tweets, {.Size})
```

### filter(array, predicate)

使用 [断言](#断言) 筛选数组的元素，返回一个新数组。

```expr
filter(users, .Name startsWith "J")
```

### find(array, predicate)

查找数组中第一个满足 [断言](#断言) 的元素。

```expr
find([1, 2, 3, 4], # > 2) == 3
```

### findIndex(array, predicate)

查找数组中第一个满足 [断言](#断言) 的元素索引。

```expr
findIndex([1, 2, 3, 4], # > 2) == 2
```

### findLast(array, predicate)

查找数组中最后一个满足 [断言](#断言) 的元素。

```expr
findLast([1, 2, 3, 4], # > 2) == 4
```

### findLastIndex(array, predicate)

查找数组中最后一个满足 [断言](#断言) 的元素索引。

```expr
findLastIndex([1, 2, 3, 4], # > 2) == 3
```

### groupBy(array, predicate)

使用 [断言](#断言)对数组元素进行分组。

```expr
groupBy(users, .Age)
```

### count(array[, predicate])

返回满足 [断言](#断言) 的元素数量。

```expr
count(users, .Age > 18)
```

等价于：

```expr
len(filter(users, .Age > 18))
```

如果没有传递断言，则返回数组中所有 `true` 元素的个数。

```javascript
count([true, false, true]) == 2
```

### concat(array1, array2[, ...])

拼接两个或更多数组。

```javascript
concat([1, 2], [3, 4]) == [1, 2, 3, 4]
```

### flatten(array)

扁平化给定数组为一维数组。

```javascript
flatten([1, 2, [3, 4]]) == [1, 2, 3, 4]
```

### uniq(array)

从数组中移除重复的元素。

```javascript
uniq([1, 2, 3, 2, 1]) == [1, 2, 3]
```

### join(array[, delimiter])

使用分隔符拼接数组中的字符串为一个新的字符串。

如果没有传递分隔符，则默认分隔符为空字符串。

```javascript
join(["apple", "orange", "grape"], ",") == "apple,orange,grape"
join(["apple", "orange", "grape"]) == "appleorangegrape"
```

### reduce(array, predicate[, initialValue])

对数组中的每个元素都使用断言，以便于归约（reduce）为一个单独的值。

可选项 `initialValue` 参数用于指定累加器（accumulator）的初始值。

如果 `initialValue` 没有指定，则数组第一个元素将会作为初始值。

可以在断言中使用以下变量：

- `#` - 当前元素
- `#acc` - 累加器（accumulator）
- `#index` - 当前元素索引

```expr
reduce(1..9, #acc + #)
reduce(1..9, #acc + #, 0)
```

### sum(array[, predicate])

返回数组中所有数值元素之和。

```javascript
sum([1, 2, 3]) == 6
```

如果提供了可选的 `predicate`（断言）参数，则在求和之前，该断言将应用于数组的每个元素。

```expr
sum(accounts, .Balance)
```

等价于:

```expr
reduce(accounts, #acc + .Balance, 0)
// 或
sum(map(accounts, .Balance))
```

### mean(array)

返回数组中所有数值的平均值。

```javascript
mean([1, 2, 3]) == 2.0
```

### median(array)

返回数组中所有数值的中位数。

```javascript
median([1, 2, 3]) == 2.0
```

### first(array)

返回数组中第一个元素。如果数组是空的，则返回 `nil`。

```javascript
first([1, 2, 3]) == 1
```

### last(array)

返回数组中最后一个元素。如果数组是空的，则返回 `nil`。

```javascript
last([1, 2, 3]) == 3
```

### take(array, n)

返回数组的前 `n` 个元素。如果数组中的元素少于 `n` 个，则返回整个数组。

```javascript
take([1, 2, 3, 4], 2) == [1, 2]
```

### reverse(array)

返回数组的一个新的反转副本。

```javascript
reverse([3, 1, 4]) == [4, 1, 3]
reverse(reverse([3, 1, 4])) == [3, 1, 4]
```

### sort(array[, order])

将数组按升序排序。可以使用可选参数 `order` 来指定排序方式：`asc`（升序）或 `desc`（降序）。

```javascript
sort([3, 1, 4]) == [1, 3, 4]
sort([3, 1, 4], "desc") == [4, 3, 1]
```

### sortBy(array[, predicate, order])

根据[断言](#断言)的结果对数组进行排序。可以使用可选参数 `order` 来指定排序方式：`asc`（升序）或 `desc`（降序）。

```expr
sortBy(users, .Age)
sortBy(users, .Age, "desc")
```

## 映射函数（Map Functions）

### keys(map)

返回一个包含该映射（Map）所有键（Keys）的数组。

```javascript
keys({"name": "John", "age": 30}) == ["name", "age"]
```

### values(map)

返回一个包含该映射（Map）所有值（Values）的数组。

```javascript
values({"name": "John", "age": 30}) == ["John", 30]
```

## 类型转换函数

### type(v)

返回给定值 `v` 的类型。

返回以下类型之一：

- `nil`
- `bool`
- `int`
- `uint`
- `float`
- `string`
- `array`
- `map`.

对于命名类型（Named types）和结构体（Structs），将返回其类型名称。

```javascript
type(42) == "int"
type("hello") == "string"
type(now()) == "time.Time"
```

### int(v)

返回数字或字符串的整数值。

```javascript
int("123") == 123
```

### float(v)

返回数字或字符串的浮点数值。

```javascript
float("123.45") == 123.45
```

### string(v)

将给定值 `v` 转换为字符串表示形式。

```javascript
string(123) == "123"
```

### toJSON(v)

将给定值 `v` 转换为其 JSON 字符串表示形式。

```javascript
toJSON({"name": "John", "age": 30})
```

### fromJSON(v)

解析给定的 JSON 字符串 `v` 并返回其对应的数据。

```javascript
fromJSON('{"name": "John", "age": 30}')
```

### toBase64(v)

将字符串 `v` 编码为 Base64 格式。

```javascript
toBase64("Hello World") == "SGVsbG8gV29ybGQ="
```

### fromBase64(v)

将 Base64 编码的字符串 `v` 解码回原始形式。

```javascript
fromBase64("SGVsbG8gV29ybGQ=") == "Hello World"
```

### toPairs(map)

将映射（Map）转换为键值对（Key-value pairs）数组。

```javascript
toPairs({"name": "John", "age": 30}) == [["name", "John"], ["age", 30]]
```

### fromPairs(array)

将键值对数组转换为映射（Map）。

```expr
fromPairs([["name", "John"], ["age", 30]]) == {"name": "John", "age": 30}
```

## 杂项函数

### len(v)

返回数组、映射（Map）或字符串的长度。

```javascript
len([1, 2, 3]) == 3
len({"name": "John", "age": 30}) == 2
len("Hello") == 5
```

### get(v, index)

从数组或映射 `v` 中检索指定索引或键的元素。如果索引超出范围或键不存在，则返回 `nil`。

```javascript
get([1, 2, 3], 1) == 2
get({"name": "John", "age": 30}, "name") == "John"
```

## 位运算函数（Bitwise Functions）

### bitand(int, int)

返回按位与（AND）运算的结果值。

```javascript
bitand(0b1010, 0b1100) == 0b1000
```

### bitor(int, int)

返回按位或（OR）运算的结果值。

```javascript
bitor(0b1010, 0b1100) == 0b1110
```

### bitxor(int, int)

返回按位异或（XOR）运算的结果值。

```javascript
bitxor(0b1010, 0b1100) == 0b110
```

### bitnand(int, int)

返回按位与非（AND NOT）运算的结果值。

```javascript
bitnand(0b1010, 0b1100) == 0b10
```

### bitnot(int)

返回按位取反（NOT）运算的结果值。

```javascript
bitnot(0b1010) == -0b1011
```

### bitshl(int, int)

返回左移（Left Shift）运算的结果值。

```javascript
bitshl(0b101101, 2) == 0b10110100
```

### bitshr(int, int)

返回右移（Right Shift）运算的结果值。

```javascript
bitshr(0b101101, 2) == 0b1011
```

### bitushr(int, int)

返回无符号右移（Unsigned Right Shift）运算的结果值。

```javascript
bitushr(-0b101, 2) == 4611686018427387902
```



