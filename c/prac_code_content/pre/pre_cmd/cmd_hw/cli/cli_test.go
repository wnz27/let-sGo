/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/10/10 10:23 10月
 **/
package main

import (
	"flag"
	"fmt"
	"strings"
	"testing"
)

var conf = flag.String("c", "./config.json",
	`当前仅支持json文件, 默认为当前目录下的config.json文件`)

var env = flag.String("env", "dev",
	`app运行环境, 当前仅支持: dev、test、prod`)
var mode = flag.String("m", "mmm", "app运行模式, 仅支持debug, release")

func mockArgs1(params []string) {
	flag.Set("c", params[0])
	flag.Set("env", params[1])
	flag.Set("m", params[2])
}

func parseArgs() {
	flag.Parse()
	confFile := strings.TrimSpace(*conf)
	parsedEnv := strings.TrimSpace(*env)
	appMode := strings.ToLower(strings.TrimSpace(*mode))
	fmt.Println("ttt: ")
	fmt.Println(confFile, parsedEnv, appMode)
	fmt.Println()
}

func TestParseFlag(t *testing.T) {
	mockArgs1([]string{"dev", "xxx", "release"})
	parseArgs()
}
