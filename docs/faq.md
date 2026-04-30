# 常见问题

## 为什么使用FlexLang替代先前的Javascript

因为Javascript是弱类型语言，即数据类型只能在运行时推导和转换，即使使用了V8
Javascript引擎，在内存耗用和运行速度上仍然达不到高性能的标准。而新引进的FlexLang，是强类型表达式语言，够用且够快。而且我们增加了很多在Javascript中常用的常量和函数，以便让用户更方便地迁移到新的语言上来。

## FlexLang怎么写多行的表达式

使用分号（`;`）可以分割多行表达式，最后一行不用写分号：

~~~
表达式1;
表达式2;
最后一个表达式
~~~

例如：

~~~javascript
let year = string(d.getFullYear());
let month = String.padStart(string(d.getMonth() + 1), 2, '0');
let day = String.padStart(string(d.getDate()), 2, '0');
year + "-" + month + "-" + day
~~~

## FlexLang中怎么做条件判断

单行中，可以使用三元运算符（`?:`）：

~~~javascript
let a = (1 == 1);
a ? 'TRUE' : 'FALSE'
~~~

在多行中，可以使用`if/else if/else`：

~~~javascript
let a = Math.randN(3);
if (a < 1) {
	"小于1"
} else if (a == 1) {
	"等于1"
} else {
	"大于1"
}
~~~

这里注意：`else`子句一定要出现，不能只有`if`没有`else`；而且每个子句中一定要有内容，不能为空。

## 我怎么知道当前正在访问哪个URL

一般通过路径来判断当前用户正在访问的URL，进而做下一步的操作：

~~~javascript
if ($req.path() == "/index.html") {
	// 这里做一些操作 ...
} else {
	true // 通常一定会有else，如果没有操作可以简单地返回一个true
}
~~~

## 怎么快速地返回自定义内容

如果我们想快速地返回一个内容，而不调用缓存和源站的话，可以使用 `$resp.send`：

~~~javascript
if ($req.path() == "/hello") {
	$resp.setHeader("Content-Type", "text/html; charset=utf-8");
	$resp.send(200, "Hello, World!")
} else {
	true
}
~~~

## 怎么方便地传递多行内容
有时候我们想返回一个HTML内容，但是如果用引号拼接字符串的话非常麻烦，针对这个问题，在FlexLang中可以使用反引号（`）来包含多行内容：
~~~javascript
if ($req.path() == "/hello") {
  $resp.setHeader("Content-Type", "text/html; charset=utf-8");
  $resp.send(200, `<html>
  <title>多行内容测试</title>
  <body>
  <h1>测试FlexLang！</h1>
  <p>亲爱的用户，你访问的这个页面为FlexLang生成。</p>
  </body>
  </html>`)
} else {
  true
}
~~~

## 怎么查看 `$req` 和 `$resp`变量内容

可以在 [变量和函数](./references.md) 中查看所有可以在FlexLang中使用的变量和函数。

## 访问用户可以看到脚本代码吗？
不可以，因为脚本在服务器端执行，不会暴露给浏览器等客户端。