# 示例

## 小试牛刀“Hello, World!”

~~~javascript
$resp.send(200, "Hello, World!")
~~~

## 拦截符合条件的请求

### 拦截路径/hello

~~~javascript
if ($req.path() == "/hello") {
	$resp.setHeader("Content-Type", "text/html; charset=utf-8");
	$resp.send(200, "Hello, World!")
} else {
	true
}
~~~

### 拦截User-Agent为Firefox的请求

~~~javascript
if (Strings.indexOf($req.header().get("User-Agent"), "Firefox") >= 0) {
	$resp.setHeader("Content-Type", "text/html; charset=utf-8");
	$resp.send(404, "Firefox Denied!")
} else {
	true
}
~~~

## 只允许用户访问特定的文件

下面例子中，除了`/hello.html`，其他页面均不允许访问

~~~javascript
if ($req.path() != "/hello.html") {
	$resp.setHeader("Content-Type", "text/html; charset=utf-8");
	$resp.send(404, "404 not found")
} else {
	true
}
~~~

这样只有访问 `/hello.html`才会出现正确的内容，否则提示`404`。

## 有点复杂的访问鉴权

对于仿阿里云TypeA `md5(path+"-"+timestamp+"-"+randomString+"-"+uid+"-"+secret)` 算法的URL：

~~~javascript
let pieces = String.split($req.query().get("auth_key"), "-");
if (len(pieces) != 4) {
	$resp.send(403, "403 Forbidden")
} else {
	let timestamp = pieces[0];
	let randomString = pieces[1];
	let uid = pieces[2];
	let hash = pieces[3];
	let path = $req.path();
	let secret = "aliyuncdnexp1234";

	if (parseInt(timestamp) < NewDate().getTime() / 1000 - 600/* 10分钟有效 */) {
		$resp.send(403, "403 Forbidden")
	} else {
		let realHash = md5(path + "-" + timestamp + "-" + randomString + "-" + uid + "-" + secret);
		if (realHash != hash) {
			$resp.send(403, "403 Forbidden")
		} else {
			true
		}
	}
}
~~~

输入类似于
`https://example.com/images/test.jpg?auth_key=1777465050-477b3bbc253f467b8def6711128c7bec-0-a7a0a7aea4b54dd5a1b7959bda221631`
可以正常访问（需要重新计算这个URL的时间戳、Hash值等数据）。

## 限制某个国家的访问
比如我们要限制`越南`国内的访问：
~~~javascript
if ($req.format("${geo.country.name}") == "越南") {
	$resp.send(403, "403 Forbidden")
} else {
	true
}
~~~

其中可以使用的变量可以参考官网文档。

## 只允许某个中国省份访问
比如我们只允许`北京`访问：
~~~javascript
if (String.indexOf($req.format("${geo.province.name}"), "北京") < 0) {
	$resp.send(403, "403 Forbidden")
} else {
	true
}
~~~

## 只允许某个IP范围访问
比如我们只允许`192.168.1.100 - 192.168.1.255`、`192.168.2.100 - 192.168.2.255`之间的IP访问：
~~~javascript
if (!NetIP.isInRanges($req.remoteAddr(), [
	["192.168.1.100", "192.168.1.255"],
	["192.168.2.100", "192.168.2.255"]
])) {
	$resp.send(403, "403 Forbidden")
} else {
	true
}
~~~

## 根据条件修改响应Header
可以使用`$resp.setHeader()`方法来修改响应Header：
~~~javascript
// 查找匹配 /webhook 的请求路径
if ($req.path() == "/webhook") {
	// 设置响应Header
	$resp.setHeader("Hello", "World")
} else {
  true
}
~~~

## 重写URL

当用户访问`/hello`时我们希望访问转移到`/world`，而且不需要跳转：

~~~javascript
if ($req.path() == "/hello") {
	$req.setURI("/world")
} else {
	true
}
~~~

## 根据参数跳转到不同的地址

我们可以根据`u`参数跳转到不同的地址：

~~~javascript
if ($req.path() == "/redirect") {
	let u = $req.query().get("u");
	if (u == "1") {
		$resp.redirect(307, "https://example.com/1")
	} else if (u == "2") {
		$resp.redirect(307, "https://example.com/2")
	} else {
		$resp.redirect(307, "https://example.com/3")
	}
} else {
	true
}
~~~

这样访问 `/redirect?u=1`、`/redirect?u=2`、`/redirect?u=3`将会跳转到不同的地址。  

## 输出调试日志
可以使用`console.log(msg)`输出调试信息：
~~~javascript
if ($req.path() == "/hello") {
	console.log("Hello, World")
} else {
	true
}
~~~

有用户访问后，可以在`边缘脚本`菜单 -- `调试日志`中查看。

控制台输出不影响网站和当前页面正常访问，且不会出现在访问用户的浏览器上。