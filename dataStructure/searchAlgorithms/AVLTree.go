package searchAlgorithms

/* 二叉平衡树 */

// AVL 树
type AVLTree struct {
	Root 	*AVLTreeNode
}

/* AVL树节点 */
type AVLTreeNode struct {
	Value 		int			// 值
	Times 		int			// 值出现的次数
	Height 		int			// 该节点作为树根节点，树的高度，方便计算平衡因子
	Left		*AVLTreeNode // 左子树
	Right		*AVLTreeNode // 右子树
}

/* 更新树的高度 */
func (node *AVLTreeNode) UpdateHeight()  {
	if node == nil {
		return
	}

	var leftHeight, rightHeight int = 0, 0

	if node.Left != nil {
		leftHeight = node.Left.Height
	}

	if node.Right != nil {
		rightHeight = node.Right.Height
	}

	// 哪棵树高算哪棵树的
	maxHeight := leftHeight
	if maxHeight < rightHeight {
		maxHeight = rightHeight
	}

	// 高度加上自己的那一层
	node.Height = maxHeight + 1
}

/* 计算书的平衡因子 */
/* 也就是左右子树的高度差 */
func (node *AVLTreeNode) BalanceFactor() int {
	var leftHeight, rightHeight int = 0, 0
	if node.Left != nil {
		leftHeight = node.Left.Height
	}

	if node.Right != nil {
		rightHeight = node.Right.Height
	}

	return leftHeight - rightHeight
}

/* AVL 树添加元素 */

// 单右旋操作
func RightRotation(Root *AVLTreeNode) *AVLTreeNode {
	// 只有Pivot和B，root位置变了
	Pivot := Root.Left
	B := Pivot.Right

	//Root.Left = Pivot.Right
	//temp := Root
	//Root = Pivot
	//Pivot.Right = temp

	Pivot.Right = Root
	Root.Left = B

	// 只有Pivot和root 变化了高度
	Root.UpdateHeight()
	Pivot.UpdateHeight()
	return Pivot
}

/* 右子树插入右儿子，左单旋 */
/* 右右，左单 */
func LeftRotation(Root *AVLTreeNode) *AVLTreeNode {
	// 只有Pivot和B，Root位置变了
	Pivot := Root.Left
	B := Pivot.Left

	Pivot.Right = Root
	Root.Left = B

	// 只有Root和Pivot变化了高度
	Root.UpdateHeight()
	Pivot.UpdateHeight()
	return Pivot

}

/* z左子树插有节点 */