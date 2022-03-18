/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2022-03-18 14:37:00
 * @LastEditTime: 2022-03-18 14:39:19
 * @FilePath: /let-sGo/db-learn/obj-storage/main.go
 * @description: type some description
 */

package main

import (
	"log"
	"net/http"
	"os"

	"obj-learn/objects"
)

func main() {
	http.HandleFunc("/objects/", objects.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
