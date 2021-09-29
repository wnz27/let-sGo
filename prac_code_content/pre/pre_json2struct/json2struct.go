/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/9/11 4:16 下午
* Description:
 */
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"age"`
	Social social `json:"social"`
}

type UserList struct {
	Users []User `json:"users"`
}

//prac_code_content/pre/pre_json2struct/user.json
const jsonFile = "./prac_code_content/pre/pre_json2struct/user.json"

func (u *User) Print() string {
	return fmt.Sprintf(`User Type: %s
User Age: %d
User Name: %s
Facebook Url: %s`, u.Type, u.Age, u.Name, u.Social.Facebook)
}

func main() {
	//读取文件
	jsonBytes := readJsonFileAll(jsonFile)

	// 解析成结构体
	var users UserList
	err := json.Unmarshal(jsonBytes, &users)
	if err != nil {
		log.Println("parse user list err", err)
		return
	}

	for _, u := range users.Users {
		fmt.Println(u.Print())
	}
	fmt.Println(users)
	//log.Printf("%+v\n", users)

}

func readJsonFileAll(jsonFilePath string) []byte {
	bytes, err := ioutil.ReadFile(jsonFilePath)
	if err != nil {
		log.Println("read json file err", err)
		return nil
	}
	log.Println("Successfully opened users.json")
	return bytes
}
