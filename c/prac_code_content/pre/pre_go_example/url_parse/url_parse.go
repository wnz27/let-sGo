/*
* Author:  a27
* Version: 1.0.0
* Date:    2021/9/22 10:52 上午
* Description:
 */
package main

import (
	"fmt"
	"net"
	"net/url"
	"reflect"
)

func main() {
	// We’ll parse this example URL, which includes a scheme, authentication info,
	// host, port, path, query params, and query fragment.
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	// Parse the URL and ensure there are no errors.
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	// Accessing the scheme is straightforward.
	fmt.Println(u.Scheme)

	// User contains all authentication info;
	// call Username and Password on this for individual values.
	fmt.Println("user -> ", u.User)
	fmt.Println("username -> ", u.User.Username())
	p, _ := u.User.Password()
	fmt.Println("password -> ", p)

	// The Host contains both the hostname and the port, if present. Use SplitHostPort to extract them.
	fmt.Println("Host -> ", u.Host)
	host, post, _ := net.SplitHostPort(u.Host)
	fmt.Println("host -> ", host)
	fmt.Println("post -> ", post)

	// Here we extract the path and the fragment after the #.
	fmt.Println("Path -> ", u.Path)
	fmt.Println("Fragment -> ", u.Fragment)

	// To get query params in a string of k=v format, use RawQuery.
	// You can also parse query params into a map.
	// The parsed query param maps are from strings to slices of strings,
	// so index into [0] if you only want the first value.
	fmt.Println("raw query -> ", u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println("query map -> ", m)
	fmt.Println(m["k"][0], reflect.TypeOf(m["k"]))

}

