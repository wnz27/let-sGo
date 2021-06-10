/**
 * @project prac_grpc
 * @Author 27
 * @Description //TODO
 * @Date 2021/5/23 12:57 5月
 **/
package main

import (
	hello_service "fzkprac/prac_grpc/server/service"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r*http.Request) {
	// string 子串可以用 [start:end] 来 获得
	//fmt.Fprintf(w, "It works !! %s %s", r.URL.Path, r.URL.Path[1:])
	// post for reflect
	data, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("input: %s \n", string(data))
	input := &hello_service.Input{}
	_ = json.Unmarshal(data, input)
	output, _ := json.Marshal(&hello_service.Output{
		Msg: "Hello, " + input.Name,
	})
	fmt.Fprintf(w, "%s", string(output))
}
// 不是所有error 都有堆栈信息
func main() {
	http.HandleFunc("/", index)

	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
