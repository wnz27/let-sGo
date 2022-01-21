/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2022-01-21 16:10:06
 * @LastEditTime: 2022-01-21 16:10:06
 * @FilePath: /let-sGo/prac_code_content/webFramwork/selfwf/gee/gee_test.go
 * @description: type some description
 */
package gee

import "testing"

func TestNestedGroup(t *testing.T) {
	r := New()
	v1 := r.Group("/v1")
	v2 := v1.Group("/v2")
	v3 := v2.Group("/v3")
	if v2.prefix != "/v1/v2" {
		t.Fatal("v2 prefix should be /v1/v2")
	}
	if v3.prefix != "/v1/v2/v3" {
		t.Fatal("v2 prefix should be /v1/v2")
	}
}
