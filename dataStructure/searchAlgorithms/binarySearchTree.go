package searchAlgorithms

import "fmt"

/* 二叉查找书 */
type BinarySearchTree struct {
	Root 	*BinarySearchTreeNode	// 树根节点
}

/* 二叉查找树节点 */
type BinarySearchTreeNode struct {
	Value 	int						// 值
	Times 	int						// 值出现的次数
	Left 	*BinarySearchTreeNode 	// 左子树
	Right 	*BinarySearchTreeNode	// 右子树
}

/* 初始化 */
func NewBinarySearchTree() *BinarySearchTree {
	return new(BinarySearchTree)
}

/* 添加元素 */
func (tree *BinarySearchTree) Add(value int)  {
	// 如果没有树根，证明是棵空树，添加树根后返回
	if tree.Root == nil {
		tree.Root = &BinarySearchTreeNode{
			Value: value,
		}
	}

	// 将值添加进去
	tree.Root.Add(value)

}

func (node *BinarySearchTreeNode) Add(value int)  {
	if value < node.Value {
		// 如果插入的值，比节点的值小，俺么插入该节点的左子树中
		if node.Left == nil {
			// 如果左子树为空，则直接插入
			node.Left = &BinarySearchTreeNode{Value:value}
		} else {
			// 否则，递归
			node.Left.Add(value)
		}
	} else if value > node.Value {
		// 如果插入的值，比节点的值打，那么插入该节点的右子树中
		if node.Right == nil {
			// 如果右子树为空，则直接插入
			node.Right = &BinarySearchTreeNode{Value:value}
		} else {
			// 否则，递归
			node.Right.Add(value)
		}
	} else {
		// 插入的值等于节点的值, 不需要添加，值出现的次数+1
		node.Times = node.Times + 1
	}
}


/* 查找最大或者最小值的元素 */
/* 最小值：左子树一直往下找；最大值：右子树一直往下找 */

// 找最小值
func (tree *BinarySearchTree) FindMinValue() *BinarySearchTreeNode {
	if tree.Root == nil {
		return nil
	}

	return tree.Root.FindMinValue()
}

func (node *BinarySearchTreeNode) FindMinValue() *BinarySearchTreeNode {
	if node.Left == nil {
		return node
	}

	// 一直是左子树递归
	return node.Left.FindMinValue()
}

// 找最大值
func (tree *BinarySearchTree) FindMaxValue() *BinarySearchTreeNode {
	if tree.Root == nil {
		return nil
	}

	return tree.Root.FindMaxValue()
}


func (node *BinarySearchTreeNode) FindMaxValue() *BinarySearchTreeNode {
	if node.Right == nil {
		return node
	}

	// 一直右子树递归
	return node.Right.FindMaxValue()
}

/* 查找指定元素 */
/* 二分查找 */
func (tree *BinarySearchTree) Find(value int) *BinarySearchTreeNode {
	if tree.Root == nil {
		return nil
	}

	return tree.Root.Find(value)
}

func (node *BinarySearchTreeNode) Find(value int) *BinarySearchTreeNode {
	if node.Value == value {
		return node
	} else if value > node.Value {
		// 如果查找值比节点值大，则查找右子树
		if node.Right == nil {
			return nil
		}
		return node.Right.Find(value)
	} else {
		// 如果查找值比节点值小，则查找左子树
		if node.Left == nil {
			return nil
		}
		return node.Left.Find(value)
	}
}

/* 查找指定元素的父亲节点 */
func (tree *BinarySearchTree) FindParent(value int) *BinarySearchTreeNode {
	if tree.Root == nil {
		// 如果是空树，直接返回
		return nil
	}

	// 如果根节点等于该值，根节点没有父节点，返回nil
	if tree.Root.Value == value {
		return nil
	}

	return tree.Root.FindParent(value)
}

func (node *BinarySearchTreeNode) FindParent(value int) *BinarySearchTreeNode {
	// 外层没有值相等的判断，因为内层已经判断完毕，并返回父亲节点

	if value < node.Value {
		// 如果查询的值，小于节点值，从节点的左子树开始找
		leftTree := node.Left
		if leftTree == nil {
			// 左子树为空，表示找不到该值
			return nil
		}

		// 如果左子树根节点的值，刚好等于该值，那么父亲节点就是现在的 node
		if leftTree.Value == value {
			return node
		} else {
			return leftTree.FindParent(value)
		}

	} else {
		// 如果查找的值大于节点值，从节点的右子树开始找
		rightTree := node.Right
		if rightTree == nil {
			// 右子树为空，表示找不到该值
			return nil
		}

		// 右子树的节点大于该节点值，那么父节点就是当前节点，
		if rightTree.Value == value {
			return node
		} else {
			return rightTree.FindParent(value)
		}
	}
}

/* 删除元素 */
func (tree *BinarySearchTree) Delete(value int)  {
	if tree.Root == nil {
		// 如果是空树，直接返回
		return
	}

	// 查找该值是否存在
	node := tree.Root.Find(value)
	if node == nil {
		// 不存在，直接返回
		return
	}

	// 查找该值的父亲节点
	parent := tree.Root.FindParent(value)

	// 第一种情况，删除的是根节点，切根节点没有儿子
	if parent == nil && node.Left == nil && node.Right == nil {
		// 置空后返回
		tree.Root = nil
		return
	} else if node.Left == nil && node.Right == nil {
		// 第二种情况，删除节点有父亲节点，但没有子树
		// 也就是叶子节点，直接删除

		// 判断删除节点是父节点的左子树还是右子树
		if parent.Left != nil && parent.Left.Value == value {
			parent.Left = nil
		} else {
			parent.Right = nil
		}
		return

	} else if node.Left != nil && node.Right != nil {
		// 第三种情况，删除节点下有两个子树
		// 因为右子树的值都比左子树大，那么用右子树的最小元素来替换删除节点

		// 找出右子树中最小的值
		minNode := node.Right
		for minNode.Left != nil  {
			minNode = minNode.Left
		}

		// 把最小的节点删去
		tree.Delete(minNode.Value)

		// 把最小值的节点替换被删除节点
		node.Value = minNode.Value
		node.Times = minNode.Times
	} else {
		// 第四种情况，只有一个子树，那么该子树直接替换被删除的节点即可

		// 父亲为空，表示删除的是根节点，替换树根
		if parent == nil {
			if node.Left != nil {
				tree.Root = node.Left
			} else {
				tree.Root = node.Right
			}
			return
		}
		// 左子树不为空
		if node.Left != nil {
			// 如果删除的是节点是父亲的左儿子，让删除的节点的左子树接班
			if parent.Left != nil && value == parent.Left.Value {
				parent.Left = node.Left
			} else {
				parent.Right = node.Left
			}
		} else {
			// 如果删除的是节点是父亲的左儿子，让删除的节点的右子树接班
			if parent.Left != nil && value == parent.Left.Value {
				parent.Left = node.Right
			} else {
				parent.Right = node.Right
			}
		}
	}

}


/* 中序遍历（实现排序） */
/* 左子树 - 根节点 - 右子树 */
func (tree *BinarySearchTree) MidOrder()  {
	tree.Root.MidOrder()
}

func (node *BinarySearchTreeNode) MidOrder()  {
	if node == nil {
		return
	}

	// 先打印左子树
	node.Left.MidOrder()

	// 按照次数打印根节点
	for i:=0; i <= int(node.Times); i++ {
		fmt.Println(node.Value)
	}

	// 打印右子树
	node.Right.MidOrder()
}




/* TODO COPY */
// 删除指定的元素
func (tree *BinarySearchTree) Delete1(value int) {
	if tree.Root == nil {
		// 如果是空树，直接返回
		return
	}

	// 查找该值是否存在
	node := tree.Root.Find(value)
	if node == nil {
		// 不存在该值，直接返回
		return
	}

	// 查找该值的父亲节点
	parent := tree.Root.FindParent(value)

	// 第一种情况，删除的是根节点，且根节点没有儿子
	if parent == nil && node.Left == nil && node.Right == nil {
		// 置空后直接返回
		tree.Root = nil
		return
	} else if node.Left == nil && node.Right == nil {
		// 第二种情况，删除的节点有父亲节点，但没有子树

		// 如果删除的是节点是父亲的左儿子，直接将该值删除即可
		if parent.Left != nil && value == parent.Left.Value {
			parent.Left = nil
		} else {
			// 删除的原来是父亲的右儿子，直接将该值删除即可
			parent.Right = nil
		}
		return
	} else if node.Left != nil && node.Right != nil {
		// 第三种情况，删除的节点下有两个子树，因为右子树的值都比左子树大，那么用右子树中的最小元素来替换删除的节点。
		// 右子树的最小元素，只要一直往右子树的左边一直找一直找就可以找到，替换后二叉查找树的性质又满足了。

		// 找右子树中最小的值，一直往右子树的左边找
		minNode := node.Right
		for minNode.Left != nil {
			minNode = minNode.Left
		}
		// 把最小的节点删掉
		tree.Delete(minNode.Value)

		// 最小值的节点替换被删除节点
		node.Value = minNode.Value
		node.Times = minNode.Times
	} else {
		// 第四种情况，只有一个子树，那么该子树直接替换被删除的节点即可

		// 父亲为空，表示删除的是根节点，替换树根
		if parent == nil {
			if node.Left != nil {
				tree.Root = node.Left
			} else {
				tree.Root = node.Right
			}
			return
		}
		// 左子树不为空
		if node.Left != nil {
			// 如果删除的是节点是父亲的左儿子，让删除的节点的左子树接班
			if parent.Left != nil && value == parent.Left.Value {
				parent.Left = node.Left
			} else {
				parent.Right = node.Left
			}
		} else {
			// 如果删除的是节点是父亲的左儿子，让删除的节点的右子树接班
			if parent.Left != nil && value == parent.Left.Value {
				parent.Left = node.Right
			} else {
				parent.Right = node.Right
			}
		}
	}
}














