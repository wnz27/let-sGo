/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2022-01-26 11:59:00
 * @LastEditTime: 2022-01-26 12:01:41
 * @FilePath: /let-sGo/prac_code_content/webFramwork/jkwf/jkframe/group.go
 * @description: type some description
 */

package jkframe

// IGroup 代表前缀分组
type IGroup interface {
	Get(string, ControllerHandler)
	Post(string, ControllerHandler)
	Put(string, ControllerHandler)
	Delete(string, ControllerHandler)
}

// Group struct 实现了IGroup
type Group struct {
	core   *Core
	prefix string
}

// 初始化Group
func NewGroup(core *Core, prefix string) *Group {
	return &Group{core: core, prefix: prefix}
}

// 实现 Get 方法
func (g *Group) Get(uri string, handler ControllerHandler) {
	uri = g.prefix + uri
	g.core.Get(uri, handler)
}

// 实现 Post 方法
func (g *Group) Post(uri string, handler ControllerHandler) {
	uri = g.prefix + uri
	g.core.Post(uri, handler)
}

// 实现 Put 方法
func (g *Group) Put(uri string, handler ControllerHandler) {
	uri = g.prefix + uri
	g.core.Put(uri, handler)
}

// 实现 Delete 方法
func (g *Group) Delete(uri string, handler ControllerHandler) {
	uri = g.prefix + uri
	g.core.Delete(uri, handler)
}

// 从core中初始化这个Group
func (c *Core) Group(prefix string) IGroup {
	return NewGroup(c, prefix)
}
