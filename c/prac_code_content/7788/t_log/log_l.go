/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/10/10 18:25 10月
 **/
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	temp, err := ioutil.TempFile("", "zap-prod-config-test")
	if err != nil {
		log.Fatal("xxxx")
	}
	fmt.Println(temp.Name())
	defer os.Remove(temp.Name())
}


