/**
 * @project let-sGo
 * @Author 27
 * @Description
 * @Date 2021/8/16 02:14 8月
 **/
package main

import (
	"bytes"
	"fmt"
	"image"
	_ "image/png"
	"io/ioutil"
	"reflect"
)

func main() {
	source := "prac_code_content/go-image/image1/1.png" 		//输入图片
	//target := "./image/result.png" //输出图片

	ff, err := ioutil.ReadFile(source) //读取文件
	if err != nil {
		panic(err)
	}

	bbb := bytes.NewBuffer(ff)
	m, formatName, err2 := image.Decode(bbb)
	if err2 != nil {
		panic(err2)
	}
	fmt.Printf("type: %s, formatName: %s\n", reflect.TypeOf(m), formatName)

	fmt.Println(m.ColorModel())
	bounds := m.Bounds()
	fmt.Println(bounds)
	//dx := bounds.Dx()
	//dy := bounds.Dy()

}
