这一次我们引入第三方依赖包，https://github.com/urfave/cli

这个项目取名netcmd，推荐用下面的目录结构，这种目录结构是目前推荐的一种做命令行工具的结构：

netcli/
- pkg/
- cmd/cli/
- vendor/
- README.md
- ...

首先下载第三方包：go get github.com/urfave/cli

然后先给大家一个小例子理解一下。
```go
// cmd/cli/cli.go
import (
"log"
"os"

"github.com/urfave/cli"
)

func main() {
err := cli.NewApp().Run(os.Args)
if err != nil {
log.Fatal(err)
}
}
```

大家在自己的终端运行一下命令行：
```shell
➜  go run cmd/my-cli/cli.go
NAME:
cli - A new cli application

USAGE:
cli [global options] command [command options] [arguments...]

VERSION:
0.0.0

COMMANDS:
help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
--help, -h     show help
--version, -v  print the version
```
是不是很帅，实现了自己第一个终端命令行工具。
接下来我们要去实现我们更多的命令参数，这个作业我们的目标是实现三个命令行工具
- ns - 根据host获取 name servers
- cname - 根据host获取 CNAME
- ip - 根据host获取IP地址

那么三个都怎么去获取的呢？给大家一些提示，这些实际上go内置的net库都已经实现了。

- ns 使用函数 ns, err := net.LookupNS(host)
- cname 使用函数 cname, err := net.LookupCNAME(host)
- ip 使用函数         ip, err := net.LookupIP(host)

所以首先这个命令行要接收两个参数，cmd和host

最终的效果如下：
```shell
./cli help
NAME:
Website Lookup CLI - Let's you query IPs, CNAMEs and Name Servers!

USAGE:
cli [global options] command [command options] [arguments...]

VERSION:
0.0.0

COMMANDS:
ns       Looks Up the NameServers for a Particular Host
cname    Looks up the CNAME for a particular host
ip       Looks up the IP addresses for a particular host
help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
--help, -h     show help
--version, -v  print the version
```
