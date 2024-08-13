/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2022-04-09 10:11:25
 * @LastEditTime: 2022-04-09 11:13:41
 * @FilePath: /let-sGo/go-high-perfomence/algorithm-explore/graph_none_dir/g_n_d.go
 * @description: type some description
 */

package graph_none_dir

// NoneDirectionGraph 无向图
type NoneDirectionGraph interface {
	Graph(V int)
	// GraphFrom(in In) // 从标准输入溜 in 读入一幅图
	V() int               // 顶点数
	E() int               // 边数
	addEdge(v int, w int) // 向图中添加一条边 v-w
	adj(v int) []int      // 和 v 相邻的所有定点
	toString() string     // 对象的字符串表示
}

// 计算 图 G 中顶点 v 的度数
func degree(G NoneDirectionGraph, v int) int {
	allvs := G.adj(v)
	return len(allvs)
}

// 计算所有顶点的最大度数
func maxDegree(G NoneDirectionGraph) int {
	var max int
	for v := 0; v < G.V(); v++ {
		currVDegree := degree(G, v)
		if currVDegree > max {
			max = currVDegree
		}
	}
	return max
}

// 所有顶点的平均度数
func avgDegree(G NoneDirectionGraph) float64 {
	return float64(2.0*G.E()) / float64(G.V())
}

// 计算自环的度数
func numberOfSelfLoop(G NoneDirectionGraph) int {
	return 0
}

// 图的邻接表的字符串表示（Graph 的实例方法）java
/*
public String toString()
{
   String s = V + " vertices, " + E + " edges\n";
   for (int v = 0; v < V; v++)
   {
      s += v + ": ";
      for (int w : this.adj(v))
         s += w + " ";
      s += "\n";
}
return s; }
*/