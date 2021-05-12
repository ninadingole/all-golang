package datastructures

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

func (n *TreeNode) insert(val int) {
	if val <= n.value {
		if n.left == nil {
			n.left = &TreeNode{value: val}
		} else {
			n.left.insert(val)
		}
	} else {
		if n.right == nil {
			n.right = &TreeNode{value: val}
		} else {
			n.right.insert(val)
		}
	}
}

func (n *TreeNode) Exists(val int) bool {
	if n == nil {
		return false
	}

	if val == n.value {
		return true
	}

	if val <= n.value {
		return n.left.Exists(val)
	} else {
		return n.right.Exists(val)
	}
}

type Tree struct {
	node *TreeNode
}

func (t *Tree) Insert(val int) *Tree {
	if t.node == nil {
		t.node = &TreeNode{value: val}
	} else {
		t.node.insert(val)
	}

	return t
}

func (t *Tree) Exists(val int) bool {
	return t.node.Exists(val)
}

func printNode(node *TreeNode) {
	if node == nil {
		return
	}

	println(node.value)
	printNode(node.left)
	printNode(node.right)
}

func PrintTree(tree *Tree) {
	printNode(tree.node)
}
