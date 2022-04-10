/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2022-01-26 15:06:41
 * @LastEditTime: 2022-01-26 16:34:22
 * @FilePath: /let-sGo/prac_code_content/webFramwork/jkwf/subject_controller.go
 * @description: type some description
 */

package main

import "jkframe"

func SubjectAddController(c *jkframe.Context) error {
	c.SetOkStatus().Json("ok, SubjectAddController")
	return nil
}

func SubjectListController(c *jkframe.Context) error {
	c.SetOkStatus().Json("ok, SubjectListController")
	return nil
}

func SubjectDelController(c *jkframe.Context) error {
	c.SetOkStatus().Json("ok, SubjectDelController")
	return nil
}

func SubjectUpdateController(c *jkframe.Context) error {
	c.SetOkStatus().Json("ok, SubjectUpdateController")
	return nil
}

func SubjectGetController(c *jkframe.Context) error {
	c.SetOkStatus().Json("ok, SubjectGetController")
	return nil
}

func SubjectNameController(c *jkframe.Context) error {
	c.SetOkStatus().Json("ok, SubjectNameController")
	return nil
}
