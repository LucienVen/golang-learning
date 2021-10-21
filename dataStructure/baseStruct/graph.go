// @Author : a1234
// @Time : 2021/10/21 11:53 上午 
// @Describe : 图

package baseStruct

import "fmt"

// 图顶点
type Node struct {
	Value interface{}
}

// 图结构，
type Graph struct {
	Nodes []*Node
	Edges map[Node][]*Node // 边
}

// 打印顶点值
func (n *Node) String() string {
	return fmt.Sprintf("%s", n.Value)
}

// 添加图顶点
func (g *Graph) AddNode(n *Node)  {
	g.Nodes = append(g.Nodes, n)
}

// 添加表
func (g *Graph) AddEdge(n1, n2 *Node)  {
	if g.Edges == nil {
		g.Edges = make(map[Node][]*Node)
	}

	// 无向图
	g.Edges[*n1] = append(g.Edges[*n1], n2) // 设定n1到n2的边
	g.Edges[*n2] = append(g.Edges[*n2], n1) // 设定n2到n1的边
}

// 打印输出图
func (g *Graph) String()  {
	s := ""

	for i := 0; i< len(g.Nodes); i++{
		s += g.Nodes[i].String() + "->"
		near := g.Edges[*g.Nodes[i]]

		for j :=0; j<len(near); j++{
			s += near[j].String() + " "
		}
		s += "\n"
	}

	fmt.Println(s)
}
