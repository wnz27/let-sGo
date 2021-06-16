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
│ └── setting
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
```go
go get -u github.com/go-ini/ini
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

```



