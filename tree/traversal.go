/**
 * @project let-sGo
 * @Author 27
 * @Description //TODO
 * @Date 2021/3/24 01:57 3月
 **/
package tree


func (node *Node) Traverse() {
	if node == nil {
		return
	}

	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}


