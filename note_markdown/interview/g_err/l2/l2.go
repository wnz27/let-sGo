/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/7/21 4:46 下午
* Description:
 */
package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"

)

func main() {
	g := new(errgroup.Group)
	var urls = []string{
		"http://www.golang.org/",
		"https://golang2.eddycjy.com/",
		"https://eddycjy.com/",
	}
	for _, url := range urls {
		url := url
		g.Go(func() error {
			resp, err := http.Get(url)
			if err == nil {
				resp.Body.Close()
			}
			return err
		})
	}
	if err := g.Wait(); err == nil {
		fmt.Println("Successfully fetched all URLs.")
	} else {
		fmt.Printf("Errors: %+v", err)
	}
}
