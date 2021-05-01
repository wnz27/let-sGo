## 2021-04-13
安装 kratos
[kratos-github](https://github.com/go-kratos/kratos)
[kratos-官网](https://go-kratos.dev)

按照官网顺序安装之后

命令行工具 发现 ：`command not found: kratos`

在go env 里是这样
```
GOPATH="/Users/a27/go"
GOROOT="/usr/local/go"
```
第一个里面有kratos
第二个里面没有kratos
cd 到第一个里面也执行不了kratos

echo $PATH 是含有GOROOT的

## 解决
zsh 的终端，.zshrc 的配置问题，重新配一下配置一下路径。



