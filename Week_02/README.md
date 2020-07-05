**学习笔记**

**一、哈希表、映射、集合**

用的做多的是map、set

应用例子： 电话号码，用户信息表、缓存、键值对

实现原理：数组+链表（出现哈希碰撞用链表处理（redis）拉链式解决冲突法）

复杂度分析：查询、添加、删除 O(1) ，最坏情况 O(n) 哈希碰撞较多（哈希函数选的不好，表的size太小了）

关联题：两数之和、三数之和、有效的字母异位词、字母异位词分组

**做题流程：**

```
//1、clarification 讨论需求细节
//2、possible solutions --> optional (time && space)
//3、code
//4、test cases
```

**二、树、二叉树、二叉搜索树的实现和特性**

思考题树的面试题解法一般都是递归的，为什么？

因为树没有清晰的线性结构，迭代不利于明白理解遍历的每个节点。

机械记忆：

**前序遍历**

```go
//根->左->右
var res []int

func preorderTraversal(root *TreeNode) []int {
	res = []int{}
	dfs(root)
	return res
}

func dfs(root *TreeNode) {
	if root != nil {
		res = append(res, root.Val)
		dfs(root.Left)
		dfs(root.Right)
	}
}
```

**中序遍历**

```go
//左->根->右
var res []int

func inorderTraversal(root *TreeNode) []int {
  res = []int{}
  dfs(root)
  return res
}

func dfs(root *TreeNode) {
  if root != nil {
    dfs(root.Left)
    res = append(res, root.Val)
    dfs(root.Right)
  }
```

}

**后序遍历**

```go
//迭代法 左->右->根
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	return postorderIterate(root)
}

func postorderIterate(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	stack, rest := Stack([]*TreeNode{root}), []int{}
	for len(stack) > 0 {
		cur := stack.Pop()
		if cur != nil {
			stack.Push(cur)
			stack.Push(nil) // 待读数标记
			if cur.Right != nil {
				stack.Push(cur.Right)
			}
			if cur.Left != nil {
				stack.Push(cur.Left)
			}
		} else {
			rest = append(rest, stack.Pop().Val)
		}
	}
	return rest
}

type Stack []*TreeNode

func (s *Stack) Push(node *TreeNode) {
	*s = append(*s, node)
}

func (s *Stack) Pop() *TreeNode {
	n := []*TreeNode(*s)[len(*s)-1]
	*s = []*TreeNode(*s)[:len(*s)-1]
	return n
}
```



**三、堆、二叉堆**

大顶堆 find O（1）

delete O（logN）

insert O（logN）-O（1）

二叉堆：

完全二叉树来实现，实现容易。不是二叉搜索树

性质：是一颗完全树

树中任意节点的值总是大于子节点的值



二叉堆一般通过数组实现

```
在数组中的索引
//左儿子 2i + 1
//右儿子 2i+2
```

插入节点：

```
新元素插到堆的尾部
从新维护堆的元素。
heapifyup（替换所有小于新元素的父亲节点）
```

删除元素：

```
堆尾元素替换堆顶元素（删除堆顶元素），重新维护堆的元素。
heapifydown
```

相关算法题：

最小的K个数、滑动窗口最大值、丑数、前K个高频元素、heapsort 自学

**图的属性和分类**

点 和 边

```
度: 点有多少个边
入度和出度
点与点之间是否连通
```

边 点与点的连线

```
有向和无向
权重 边长，损耗
```

**图的表示和分类**

无向无权图

有向无权图

无向有权图

有向有权图

**基于图相关的算法**

DFS-递归算法

<u>visited 集合</u>

BFS

<u>visited 集合</u>

题：连通图个数



