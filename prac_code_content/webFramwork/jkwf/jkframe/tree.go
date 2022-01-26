/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2022-01-26 12:05:57
 * @LastEditTime: 2022-01-26 12:07:10
 * @FilePath: /let-sGo/prac_code_content/webFramwork/jkwf/jkframe/tree.go
 * @description: type some description
 */

package jkframe

/*
- 定义树和节点的数据结构
- 编写函数：增加路由规则
- 编写函数：查找路由
- 将 "增加路由规则" 和 "查找路由" 添加到框架中
*/

// 代表树结构
type Tree struct {
	root *node // 根节点
}

// 代表节点
type node struct {
	isLast  bool              // 代表这个节点是否可以成为最终的路由规则。该节点是否能成为一个独立的uri, 是否自身就是一个终极节点
	segment string            // uri中的字符串，代表这个节点表示的路由中某个段的字符串
	handler ControllerHandler // 代表这个节点中包含的控制器，用于最终加载调用
	childs  []*node           // 代表这个节点下的子节点
}
