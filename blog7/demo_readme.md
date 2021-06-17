#  题外话 go 本身结构
`tree -L 1 /usr/local/go` 输出：
```go
/usr/local/go
├── AUTHORS
├── CONTRIBUTING.md
├── CONTRIBUTORS
├── LICENSE
├── PATENTS
├── README.md
├── SECURITY.md
├── VERSION
├── api
├── bin
├── doc
├── favicon.ico
├── lib
├── misc
├── pkg
├── robots.txt
├── src
└── test
```
- api：用于存放依照 Go 版本顺序的 API 增量列表文件。这里所说的 API 包含公开的变量、常量、函数等。这些 API 增量列表文件用于 Go 语言 API 检查
- bin：用于存放主要的标准命令文件（可执行文件），包含go、godoc、gofmt
- blog：用于存放官方博客中的所有文章(现在可能没有)
- doc：用于存放标准库的 HTML 格式的程序文档。我们可以通过godoc命令启动一个 Web 程序展示这些文档
- lib：用于存放一些特殊的库文件
- misc：用于存放一些辅助类的说明和工具
- pkg：用于存放安装Go标准库后的所有归档文件（以.a结尾的文件）。注意，你会发现其中有名称为linux_amd64的文件夹，我们称为平台相关目录。这类文件夹的名称由对应的操作系统和计算架构的名称组合而成。通过go install命令，Go程序会被编译成平台相关的归档文件存放到其中
- src：用于存放 Go自身、Go 标准工具以及标准库的所有源码文件
- test：存放用来测试和验证Go本身的所有相关文件


## 初始化行为
在前面我们已经了解到 Go 依赖包管理的历史情况，首先你需要有一个你喜欢的目录，
例如：`$ mkdir ~/go-application && cd ~/go-application`，然后执行如下命令：
```go
$ mkdir go-gin-example && cd go-gin-example

$ go env -w GO111MODULE=on

$ go env -w GOPROXY=https://goproxy.cn,direct

$ go mod init github.com/EDDYCJY/go-gin-example
go: creating new go.mod: module github.com/EDDYCJY/go-gin-example
$ ls
go.mod
```
- `go env -w GO111MODULE=on`：打开 Go modules 开关（目前在 Go1.13 中默认值为 auto）。
- `go env -w GOPROXY=...`：设置 GOPROXY 代理，这里主要涉及到两个值，第一个是 https://goproxy.cn，它是由七牛云背书的一个强大稳定的 Go 模块代理，可以有效地解决你的外网问题；第二个是 direct，它是一个特殊的 fallback 选项，它的作用是用于指示 Go 在拉取模块时遇到错误会回源到模块版本的源地址去抓取（比如 GitHub 等）。
- `go mod init [MODULE_PATH]`：初始化 Go modules，它将会生成 go.mod 文件，需要注意的是 MODULE_PATH 填写的是模块引入路径，你可以根据自己的情况修改路径。

## 开启gin
回到刚刚创建的 go-gin-example 目录下，在命令行下执行如下命令：
```go
go get -u github.com/gin-gonic/gin
```

## 思考
程序的文本配置写在代码中，好吗？

API 的错误码硬编码在程序中，合适吗？

db 句柄谁都去Open，没有统一管理，好吗？

获取分页等公共参数，谁都自己写一套逻辑，好吗？

肯定不可以。

本demo选用 [go-ini/ini](https://github.com/go-ini/ini) ，它的 [中文文档](https://ini.unknwon.io/) 。

## 初始化项目目录
我们初始化了一个 go-gin-example 项目，接下来我们需要继续新增如下目录结构：
```go
go-gin-example/
├── conf
├── go.mod
├── go.sum
├── middleware
├── models
├── pkg
│   └── setting
├── routers
├── runtime
└── test.go
```
- conf：用于存储配置文件
- middleware：应用中间件
- models：应用数据库模型
- pkg：第三方包
- routers 路由逻辑处理
- runtime：应用运行时数据

## 添加 Go Modules Replace
打开 go.mod 文件，新增 replace 配置项，如下：
```go
replace (
    github.com/EDDYCJY/go-gin-example/conf => ~/fzk27/learn/let-sGo/blog7/go-gin-example/conf v1.7.2
    github.com/EDDYCJY/go-gin-example/middleware => ~/fzk27/learn/let-sGo/blog7/go-gin-example/middleware v1.7.2
    github.com/EDDYCJY/go-gin-example/models => ~/fzk27/learn/let-sGo/blog7/go-gin-example/models v1.7.2
    github.com/EDDYCJY/go-gin-example/pkg/setting => ~/fzk27/learn/let-sGo/blog7/go-gin-example/pkg/setting v1.7.2
    github.com/EDDYCJY/go-gin-example/routers => ~/fzk27/learn/let-sGo/blog7/go-gin-example/routers v1.7.2
)
```
可能你会不理解为什么要特意跑来加 replace 配置项，
首先你要看到我们使用的是完整的外部模块引用路径（github.com/EDDYCJY/go-gin-example/xxx），
而这个模块还没推送到远程，是没有办法下载下来的，因此需要用 replace 将其指定读取本地的模块路径，
这样子就可以解决本地模块读取的问题。

> 注：后续每新增一个本地应用目录，你都需要主动去 go.mod 文件里新增一条 replace（我不会提醒你），如果你漏了，那么编译时会出现报错，
找不到那个模块。

## 初始项目数据库
新建 blog 数据库，编码为utf8_general_ci，在 blog 数据库下，新建以下表
1、 标签表
```mysql
CREATE TABLE `blog_tag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT '' COMMENT '标签名称',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) unsigned DEFAULT '0',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章标签管理';
```
2、 文章表
```mysql
CREATE TABLE `blog_article` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `tag_id` int(10) unsigned DEFAULT '0' COMMENT '标签ID',
  `title` varchar(100) DEFAULT '' COMMENT '文章标题',
  `desc` varchar(255) DEFAULT '' COMMENT '简述',
  `content` text,
  `created_on` int(11) DEFAULT NULL,
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(255) DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) unsigned DEFAULT '0',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用1为启用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章管理';
```
3、 认证表
```mysql
CREATE TABLE `blog_auth` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) DEFAULT '' COMMENT '账号',
  `password` varchar(50) DEFAULT '' COMMENT '密码',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `blog`.`blog_auth` (`id`, `username`, `password`) VALUES (null, 'test', 'test123456');
```
## 编写项目配置包
在 go-gin-example 应用目录下，拉取 go-ini/ini 的依赖包，如下：
```shell
$ go get -u github.com/go-ini/ini
```

接下来我们需要编写基础的应用配置文件，
在 go-gin-example 的conf目录下新建app.ini文件，写入内容：
```go
#debug or release
RUN_MODE = debug

[app]
PAGE_SIZE = 10
JWT_SECRET = 23347$040412

[server]
HTTP_PORT = 8000
READ_TIMEOUT = 60
WRITE_TIMEOUT = 60

[database]
TYPE = mysql
USER = 数据库账号
PASSWORD = 数据库密码
#127.0.0.1:3306
HOST = 数据库IP:数据库端口号
NAME = blog
TABLE_PREFIX = blog_
```
建立调用配置的setting模块，在go-gin-example的pkg目录下新建setting目录（注意新增 replace 配置），
新建 setting.go 文件，写入内容：
```go
package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

var (
	Cfg *ini.File

	RunMode string

	HTTPPort int
	ReadTimeout time.Duration
	WriteTimeout time.Duration

	PageSize int
	JwtSecret string
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	LoadBase()
	LoadServer()
	LoadApp()
}

func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout =  time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}

	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}
```
当前的目录结构：
```go
go-gin-example/
├── conf
│   └── app.ini
├── go.mod
├── go.sum
├── middleware
├── models
├── pkg
│   └── setting
├── routers
├── runtime
└── test.go
```
## 编写 API 错误码包
建立错误码的e模块，在go-gin-example的pkg目录下新建e目录（注意新增 replace 配置），
新建code.go和msg.go文件，写入内容
1、 code.go：
```go
/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/6/16 5:42 下午
* Description:
 */
package e

const (
	SUCCESS = 200
	ERROR = 500
	INVALID_PARAMS = 400

	ERROR_EXIST_TAG = 10001
	ERROR_NOT_EXIST_TAG = 10002
	ERROR_NOT_EXIST_ARTICLE = 10003

	ERROR_AUTH_CHECK_TOKEN_FAIL = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN = 20003
	ERROR_AUTH = 20004
)
```
2、 msg.go：
```go
package e

var MsgFlags = map[int]string {
	SUCCESS : "ok",
	ERROR : "fail",
	INVALID_PARAMS : "请求参数错误",
	ERROR_EXIST_TAG : "已存在该标签名称",
	ERROR_NOT_EXIST_TAG : "该标签不存在",
	ERROR_NOT_EXIST_ARTICLE : "该文章不存在",
	ERROR_AUTH_CHECK_TOKEN_FAIL : "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT : "Token已超时",
	ERROR_AUTH_TOKEN : "Token生成失败",
	ERROR_AUTH : "Token错误",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
```
## 编写工具包
在go-gin-example的pkg目录下新建util目录（注意新增 replace 配置），并拉取com的依赖包，如下：
```shell
$ go get -u github.com/unknwon/com
```
新增replace
```go
github.com/EDDYCJY/go-gin-example/pkg/util => ~/your-app-path/go-gin-example/pkg/util v1.7.2
```
编写分页页码的获取方法
在util目录下新建pagination.go，写入内容：
```go
package util

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"

	"github.com/EDDYCJY/go-gin-example/pkg/setting"
)

func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
    if page > 0 {
        result = (page - 1) * setting.PageSize
    }

    return result
}
```
## 编写 models init
拉取gorm的依赖包，如下：
```shell
$ go get -u github.com/jinzhu/gorm
```
拉取mysql驱动的依赖包，如下：
```shell
$ go get -u github.com/go-sql-driver/mysql
```
完成后，在go-gin-example的models目录下新建models.go，用于models的初始化使用
```go
package models

import (
	"log"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/EDDYCJY/go-gin-example/pkg/setting"
)

var db *gorm.DB

type Model struct {
	ID int `gorm:"primary_key" json:"id"`
	CreatedOn int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

func init() {
	var (
		err error
		dbType, dbName, user, password, host, tablePrefix string
	)

	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
	    return tablePrefix + defaultTableName;
	}

	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer db.Close()
}
```
## 编写项目启动、路由文件

编写 Demo
在go-gin-example下建立main.go作为启动文件（也就是main包），我们先写个Demo，帮助大家理解，写入文件内容：
```go
package main

import (
    "fmt"
    "net/http"

    "github.com/gin-gonic/gin"

    "github.com/EDDYCJY/go-gin-example/pkg/setting"
)

func main() {
    router := gin.Default()
    router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
```
执行`go run main.go`，查看命令行是否显示
```shell
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /test                     --> main.main.func1 (3 handlers)
```
在本机执行curl 127.0.0.1:8000/test，检查是否返回{"message":"test"}

## 知识点

### 标准库
- [fmt](https://golang.org/pkg/fmt/) ：实现了类似 C 语言 printf 和 scanf 的格式化 I/O。格式化动作（‘verb’）源自 C 语言但更简单
- [net/http](https://golang.org/pkg/net/http/) ：提供了 HTTP 客户端和服务端的实现
- Gin
    - [gin.Default()](https://gowalker.org/github.com/gin-gonic/gin#Default) ：返回 Gin 的type Engine struct{...}，里面包含RouterGroup，相当于创建一个路由Handlers，可以后期绑定各类的路由规则和函数、中间件等
    - [router.GET(…){…}](https://gowalker.org/github.com/gin-gonic/gin#IRoutes) ：创建不同的 HTTP 方法绑定到Handlers中，也支持 POST、PUT、DELETE、PATCH、OPTIONS、HEAD 等常用的 Restful 方法
    - [gin.H{…}](https://gowalker.org/github.com/gin-gonic/gin#H) ：就是一个map[string]interface{}
    - [gin.Context](https://gowalker.org/github.com/gin-gonic/gin#Context) ：Context是gin中的上下文，它允许我们在中间件之间传递变量、管理流、验证 JSON 请求、响应 JSON 请求等，在gin中包含大量Context的方法，例如我们常用的DefaultQuery、Query、DefaultPostForm、PostForm等等
    
### &http.Server 和 ListenAndServe？
#### 1、http.Server：
```go
type Server struct {
    Addr    string
    Handler Handler
    TLSConfig *tls.Config
    ReadTimeout time.Duration
    ReadHeaderTimeout time.Duration
    WriteTimeout time.Duration
    IdleTimeout time.Duration
    MaxHeaderBytes int
    ConnState func(net.Conn, ConnState)
    ErrorLog *log.Logger
}
```
- Addr：监听的 TCP 地址，格式为:8000
- Handler：http 句柄，实质为ServeHTTP，用于处理程序响应 HTTP 请求
- TLSConfig：安全传输层协议（TLS）的配置
- ReadTimeout：允许读取的最大时间
- ReadHeaderTimeout：允许读取请求头的最大时间
- WriteTimeout：允许写入的最大时间
- IdleTimeout：等待的最大时间
- MaxHeaderBytes：请求头的最大字节数
- ConnState：指定一个可选的回调函数，当客户端连接发生变化时调用
- ErrorLog：指定一个可选的日志记录器，用于接收程序的意外行为和底层系统错误；如果未设置或为nil则默认以日志包的标准日志记录器完成（也就是在控制台输出） 
  
#### 2、 ListenAndServe：
```go
func (srv *Server) ListenAndServe() error {
    addr := srv.Addr
    if addr == "" {
        addr = ":http"
    }
    ln, err := net.Listen("tcp", addr)
    if err != nil {
        return err
    }
    return srv.Serve(tcpKeepAliveListener{ln.(*net.TCPListener)})
}
```
开始监听服务，监听 TCP 网络地址，Addr 和调用应用程序处理连接上的请求。

我们在源码中看到Addr是调用我们在&http.Server中设置的参数，因此我们在设置时要用&，
我们要改变参数的值，因为我们ListenAndServe和其他一些方法需要用到&http.Server中的参数，他们是相互影响的。
#### 3、 http.ListenAndServe和 连载一 的r.Run()有区别吗？
我们看看r.Run的实现：
```go
func (engine *Engine) Run(addr ...string) (err error) {
    defer func() { debugPrintError(err) }()

    address := resolveAddress(addr)
    debugPrint("Listening and serving HTTP on %s\n", address)
    err = http.ListenAndServe(address, engine)
    return
}
```
通过分析源码，得知本质上没有区别，同时也得知了启动gin时的监听 debug 信息在这里输出。

#### 4、 为什么 Demo 里会有WARNING？
首先我们可以看下Default()的实现
```go
// Default returns an Engine instance with the Logger and Recovery middleware already attached.
func Default() *Engine {
	debugPrintWARNINGDefault()
	engine := New()
	engine.Use(Logger(), Recovery())
	return engine
}
```
大家可以看到默认情况下，已经附加了日志、恢复中间件的引擎实例。
并且在开头调用了debugPrintWARNINGDefault()，而它的实现就是输出该行日志
```go
func debugPrintWARNINGDefault() {
	debugPrint(`[WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.
`)
}
```
而另外一个Running in "debug" mode. Switch to "release" mode in production.，是运行模式原因，并不难理解，
已在配置文件的管控下 :-)，运维人员随时就可以修改它的配置。

#### 5、 Demo 的router.GET等路由规则可以不写在main包中吗？
我们发现router.GET等路由规则，在 Demo 中被编写在了main包中，感觉很奇怪，我们去抽离这部分逻辑！

在go-gin-example下routers目录新建router.go文件，写入内容：
```go
package routers

import (
    "github.com/gin-gonic/gin"

    "github.com/EDDYCJY/go-gin-example/pkg/setting"
)

func InitRouter() *gin.Engine {
    r := gin.New()

    r.Use(gin.Logger())

    r.Use(gin.Recovery())

    gin.SetMode(setting.RunMode)

    r.GET("/test", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "test",
        })
    })

    return r
}
```
修改main.go的文件内容：
```go
package main

import (
	"fmt"
	"net/http"

	"github.com/EDDYCJY/go-gin-example/routers"
	"github.com/EDDYCJY/go-gin-example/pkg/setting"
)

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
```
当前目录结构
```shell
go-gin-example
├── conf
│   └── app.ini
├── go.mod
├── go.sum
├── main.go
├── middleware
├── models
│   └── models.go
├── pkg
│   ├── e
│   │   ├── code.go
│   │   └── msg.go
│   ├── setting
│   └── util
|       └── pagination.go
├── routers
│   └── router.go
├── runtime
└── test.go
```
## 接下来涉及
- Gin：Golang 的一个微框架，性能极佳。
- [beego-validation](https://github.com/astaxie/beego/tree/master/validation) ：本节采用的 beego 的表单验证库，[中文文档](https://beego.me/docs/mvc/controller/validation.md) 。
- [gorm](https://github.com/jinzhu/gorm) ，对开发人员友好的 ORM 框架，[英文文档](https://gorm.io/docs/)
- [com](https://github.com/Unknwon/com) ，一个小而美的工具包。

## 定义接口
一般接口为增删改查是基础的

- 获取标签列表：GET("/tags”)
- 新建标签：POST("/tags”)
- 更新指定标签：PUT("/tags/:id”)
- 删除指定标签：DELETE("/tags/:id”)

编写路由空壳
开始编写路由文件逻辑，在routers下新建api目录，我们当前是第一个 API 大版本，因此在api下新建v1目录，再新建tag.go文件，写入内容：
```go
package v1

import (
	"github.com/gin-gonic/gin"
)

//获取多个文章标签
func GetTags(c *gin.Context) {
}

//新增文章标签
func AddTag(c *gin.Context) {
}

//修改文章标签
func EditTag(c *gin.Context) {
}

//删除文章标签
func DeleteTag(c *gin.Context) {
}
```
## 注册路由
我们打开routers下的router.go文件，修改文件内容为：
```go
package routers

import (
	v1 "github.com/EDDYCJY/go-gin-example/routers/api/v1"
	"github.com/gin-gonic/gin"

	"github.com/EDDYCJY/go-gin-example/pkg/setting"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	apiv1 := r.Group("/api/v1")
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
	}

	return r
}
```
当前目录结构：
```go
go-gin-example
├── conf
│   └── app.ini
├── go.mod
├── go.sum
├── main.go
├── middleware
├── models
│   └── models.go
├── pkg
│     ├── e
│     │   ├── code.go
│     │   └── msg.go
│     ├── setting
│     │   └── setting.go
│     └── util
│         └── pagination.go
├── routers
│   ├── api
│   │   └── v1
│   │       └── tag.go
│   └── router.go
├── runtime
└── test.go
```
### 检验路由是否注册成功
回到命令行，执行go run main.go，检查路由规则是否注册成功。
```shell
$ go run main.go
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /api/v1/tags              --> gin-blog/routers/api/v1.GetTags (3 handlers)
[GIN-debug] POST   /api/v1/tags              --> gin-blog/routers/api/v1.AddTag (3 handlers)
[GIN-debug] PUT    /api/v1/tags/:id          --> gin-blog/routers/api/v1.EditTag (3 handlers)
[GIN-debug] DELETE /api/v1/tags/:id          --> gin-blog/routers/api/v1.DeleteTag (3 handlers)
```
## 下载依赖包
首先我们要拉取validation的依赖包，在后面的接口里会使用到表单验证
```shell
$ go get -u github.com/astaxie/beego/validation
```
## 编写标签列表的 models 逻辑
创建models目录下的tag.go，写入文件内容：
```go
package models

type Tag struct {
    Model

    Name string `json:"name"`
    CreatedBy string `json:"created_by"`
    ModifiedBy string `json:"modified_by"`
    State int `json:"state"`
}

func GetTags(pageNum int, pageSize int, maps interface {}) (tags []Tag) {
    db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)

    return
}

func GetTagTotal(maps interface {}) (count int){
    db.Model(&Tag{}).Where(maps).Count(&count)

    return
}
```
1. 我们创建了一个Tag struct{}，用于Gorm的使用。并给予了附属属性json，这样子在c.JSON的时候就会自动转换格式，非常的便利
2. 可能会有的初学者看到return，而后面没有跟着变量，会不理解；其实你可以看到在函数末端，我们已经显示声明了返回值，这个变量在函数体内也可以直接使用，因为他在一开始就被声明了
3. 有人会疑惑db是哪里来的；因为在同个models包下，因此db *gorm.DB是可以直接使用的

## 编写标签列表的路由逻辑
打开routers目录下 v1 版本的tag.go，第一我们先编写获取标签列表的接口
修改文件内容：
```go
package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"

	"github.com/EDDYCJY/go-gin-example/models"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/EDDYCJY/go-gin-example/pkg/setting"
	"github.com/EDDYCJY/go-gin-example/pkg/util"
)

//获取多个文章标签
func GetTags(c *gin.Context) {
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := e.SUCCESS

	data["lists"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : data,
	})
}

//新增文章标签
func AddTag(c *gin.Context) {
}

//修改文章标签
func EditTag(c *gin.Context) {
}

//删除文章标签
func DeleteTag(c *gin.Context) {
}
```
1. c.Query可用于获取?name=test&state=1这类 URL 参数，而c.DefaultQuery则支持设置一个默认值
2. code变量使用了e模块的错误编码，这正是先前规划好的错误码，方便排错和识别记录
3. util.GetPage保证了各接口的page处理是一致的
4. c *gin.Context是Gin很重要的组成部分，可以理解为上下文，它允许我们在中间件之间传递变量、管理流、验证请求的 JSON 和呈现 JSON 响应

在本机执行curl 127.0.0.1:8000/api/v1/tags，正确的返回值为{"code":200,"data":{"lists":[],"total":0},"msg":"ok"}，若存在问题请结合 gin 结果进行拍错。

在获取标签列表接口中，我们可以根据name、state、page来筛选查询条件，分页的步长可通过app.ini进行配置，以lists、total的组合返回达到分页效果。

## 编写新增标签的 models 逻辑
接下来我们编写新增标签的接口
打开models目录下的tag.go，修改文件（增加 2 个方法）：
```go
func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name = ?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}

	return false
}

func AddTag(name string, state int, createdBy string) bool{
	db.Create(&Tag {
		Name : name,
		State : state,
		CreatedBy : createdBy,
	})

	return true
}

```
编写新增标签的路由逻辑
打开routers目录下的tag.go，修改文件（变动 AddTag 方法）：
```go
package v1

import (
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/astaxie/beego/validation"
    "github.com/Unknwon/com"

    "gin-blog/pkg/e"
    "gin-blog/models"
    "gin-blog/pkg/util"
    "gin-blog/pkg/setting"
)

...

//新增文章标签
func AddTag(c *gin.Context) {
    name := c.Query("name")
    state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
    createdBy := c.Query("created_by")

    valid := validation.Validation{}
    valid.Required(name, "name").Message("名称不能为空")
    valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
    valid.Required(createdBy, "created_by").Message("创建人不能为空")
    valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100字符")
    valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

    code := e.INVALID_PARAMS
    if ! valid.HasErrors() {
        if ! models.ExistTagByName(name) {
            code = e.SUCCESS
            models.AddTag(name, state, createdBy)
        } else {
            code = e.ERROR_EXIST_TAG
        }
    }

    c.JSON(http.StatusOK, gin.H{
        "code" : code,
        "msg" : e.GetMsg(code),
        "data" : make(map[string]string),
    })
}
...
```
用Postman用 POST 访问http://127.0.0.1:8000/api/v1/tags?name=1&state=1&created_by=test，查看code是否返回200及blog_tag表中是否有值，有值则正确。

## 编写 models callbacks
但是这个时候大家会发现，我明明新增了标签，但created_on居然没有值，那做修改标签的时候modified_on会不会也存在这个问题？

为了解决这个问题，我们需要打开models目录下的tag.go文件，修改文件内容（修改包引用和增加 2 个方法）：
```go
package models

import (
    "time"

    "github.com/jinzhu/gorm"
)

...

func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
    scope.SetColumn("CreatedOn", time.Now().Unix())

    return nil
}

func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
    scope.SetColumn("ModifiedOn", time.Now().Unix())

    return nil
}
```
重启服务，再在用Postman用 POST 访问http://127.0.0.1:8000/api/v1/tags?name=2&state=1&created_by=test，发现created_on已经有值了！

在这几段代码中，涉及到知识点：

> 这属于gorm的Callbacks，可以将回调方法定义为模型结构的指针，在创建、更新、查询、删除时将被调用，如果任何回调返回错误，gorm 将停止未来操作并回滚所有更改。

gorm所支持的回调方法：
- 创建：BeforeSave、BeforeCreate、AfterCreate、AfterSave
- 更新：BeforeSave、BeforeUpdate、AfterUpdate、AfterSave
- 删除：BeforeDelete、AfterDelete
- 查询：AfterFind

## 编写其余接口的路由逻辑
接下来，把剩下两个接口（EditTag、DeleteTag）完成
打开routers目录下 v1 版本的tag.go文件，修改内容：
```go
//修改文章标签
func EditTag(c *gin.Context) {
    id := com.StrTo(c.Param("id")).MustInt()
    name := c.Query("name")
    modifiedBy := c.Query("modified_by")

    valid := validation.Validation{}

    var state int = -1
    if arg := c.Query("state"); arg != "" {
        state = com.StrTo(arg).MustInt()
        valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
    }

    valid.Required(id, "id").Message("ID不能为空")
    valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
    valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")
    valid.MaxSize(name, 100, "name").Message("名称最长为100字符")

    code := e.INVALID_PARAMS
    if ! valid.HasErrors() {
        code = e.SUCCESS
        if models.ExistTagByID(id) {
            data := make(map[string]interface{})
            data["modified_by"] = modifiedBy
            if name != "" {
                data["name"] = name
            }
            if state != -1 {
                data["state"] = state
            }

            models.EditTag(id, data)
        } else {
            code = e.ERROR_NOT_EXIST_TAG
        }
    }

    c.JSON(http.StatusOK, gin.H{
        "code" : code,
        "msg" : e.GetMsg(code),
        "data" : make(map[string]string),
    })
}

//删除文章标签
func DeleteTag(c *gin.Context) {
    id := com.StrTo(c.Param("id")).MustInt()

    valid := validation.Validation{}
    valid.Min(id, 1, "id").Message("ID必须大于0")

    code := e.INVALID_PARAMS
    if ! valid.HasErrors() {
        code = e.SUCCESS
        if models.ExistTagByID(id) {
            models.DeleteTag(id)
        } else {
            code = e.ERROR_NOT_EXIST_TAG
        }
    }

    c.JSON(http.StatusOK, gin.H{
        "code" : code,
        "msg" : e.GetMsg(code),
        "data" : make(map[string]string),
    })
}
```
## 编写其余接口的 models 逻辑
打开models下的tag.go，修改文件内容：
```go
func ExistTagByID(id int) bool {
    var tag Tag
    db.Select("id").Where("id = ?", id).First(&tag)
    if tag.ID > 0 {
        return true
    }

    return false
}

func DeleteTag(id int) bool {
    db.Where("id = ?", id).Delete(&Tag{})

    return true
}

func EditTag(id int, data interface {}) bool {
    db.Model(&Tag{}).Where("id = ?", id).Updates(data)

    return true
}
```
### 验证功能
重启服务，用 Postman
PUT 访问 http://127.0.0.1:8000/api/v1/tags/1?name=edit1&state=0&modified_by=edit1 ，查看 code 是否返回 200
DELETE 访问 http://127.0.0.1:8000/api/v1/tags/1 ，查看 code 是否返回 200



