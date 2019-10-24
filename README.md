
<p align="center">
	<img src="https://github.com/shzy2012/static/blob/master/toolbox.png?raw=true" width="120" height="120">
</p>

<h1 align="center">http-server</h1>

用途: 为静态资源提供http-server功能

优点: 简单

使用: 进入website目录,执行http-server即可
```bash
http-server -p=8080

参数解释：
-p: 端口,默认8080
```

安装
```bash
# brew 安装
brew install shzy2012/tap/http-server
```


示例
```bash
➜ http-server
Starting up http-server, serving ./
Available on:
        http://127.0.0.1:8080
        http://192.168.10.49:8080
        http://192.168.195.71:8080
Hit CTRL-C to stop the server
2019/09/06 15:15:17 GET	/static/img/avatars/avatar-2.jpg	127.0.0.1:49895	Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.132 Safari/537.36	154.351µs
2019/09/06 15:15:30 GET	/static/img/avatars/avatar-1.jpg	127.0.0.1:49895	Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.132 Safari/537.36	79.361µs
```
