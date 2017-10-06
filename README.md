# Wamp-Server
Wamp服务器实现。关于WAMP的相关项目和介绍详情参考 [https://crossbar.io/](https://crossbar.io/)

## 使用说明

该服务器程序基于 [turnpike](https://github.com/jcelliott/turnpike) 实现.

* 安装golang环境，参考 [https://golang.org/](https://golang.org/)
* 安装turnpike库： `go get -u -v gopkg.in/jcelliott/turnpike.v2`
* 编译: `go build -o wamp-server main.go`
* 运行：`./wamp-server -debug true -port 8888 -realm yourapp.domain`
* 具体说明，查看帮助：`./wamp-server -help`

该版本修复了不同浏览器出现 CheckOrigin 失败导致的403错误。