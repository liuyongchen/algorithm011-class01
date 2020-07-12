package Week_03

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//递归
func buildTree(preorder []int, inorder []int) *TreeNode {

	for k, v := range inorder {
		if inorder[k] == preorder[0] {
			return &TreeNode{
				Val:   v,
				Left:  buildTree(preorder[1:k+1], inorder[:k]),
				Right: buildTree(preorder[k+1:], inorder[k+1:]),
			}
		}
	}
	return nil
}

//哈希 执行时间更短，为优化执行效率，减少程序的运算时间
func buildTree0(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	m := make(map[int]int)
	if len(preorder) != len(inorder) {
		return nil
	}
	for k, v := range inorder {
		m[v] = k
	}

	root := treebuild(preorder, inorder, 0, len(preorder)-1, 0, len(inorder)-1, m)
	return root
}

func treebuild(preorder, inorder []int, pl, pr, il, ir int, m map[int]int) *TreeNode {
	if pl > pr {
		return nil
	}
	root := &TreeNode{}
	root.Val = preorder[pl]
	idx := m[preorder[pl]]
	pltr := idx - il + pl
	root.Left = treebuild(preorder, inorder, pl+1, pltr, il, idx-1, m)
	root.Right = treebuild(preorder, inorder, pltr+1, pr, idx+1, ir, m)
	return root
}
