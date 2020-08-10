学习笔记

1、字典树

结点本身不存储完整的单词

从根结点到某一结点，路径上经过的字符连接起来，为该结点对应的字符串

每个结点的所有子结点路径代表的字符都不相同

2、高级搜索

2.1初级搜索

朴素搜索

优化方式：不重复（fibonacci） 剪枝（生成括号问题，作业剪枝）

搜索方向：

DFS：depth first search 深度优先搜索

BFS：breadth first search 广度优先搜索

双向搜索、启发式搜索

生成括号问题，八皇后问题，（剪枝）、数独问题引入（优先级优先搜索）双向bfs 模板



```python
A* search
def AstarSearch(graph,start,end):
	pq = collections.priority_queue() #优先级->估价函数
  pq.append([start])
  visited.add(start)
  
  while pq:
    node = pq.pop() #can we add more intelligence here?
    visited.add(node)
    
    process(node)
    nodes = generate_related_nodes(node)
    unvisited = [node for node in nodes if node not in visited]
    pq.push(unvisited)
```

#估价函数

启发式函数：h(n)，它用来评价哪些结点最有希望的是一个我们要找的结点，h(n)会返回一个非负实数，也可以认为是从结点n的目标结点路径的估计成本。

启发式函数是一种告知搜索方向的方法。它提供了一种明智的方法来猜测哪个邻居结点会导向一个目标。

最短路径A* 解法，"曼哈顿距离"

红黑树和AVL树



```go
//前序遍历 根左右
func preorder(root *TreeNode) {
  if root != nil {
    traverse_path = append(traverse_path,root.val)
    preorder(root.left)
    preorder(root.right)
  }
}

//中序遍历 左根右
func inorder(root *TreeNode) {
  if root != nil {
    inorder(root.left)
    traverse_path = append(traverse_path,root.val)
    inorder(root.right)
  }
}

//后序遍历 左友根
func postorder(root *TreeNode) {
  if root != nil {
    postorder(root.left)
    postorder(root.right)
    traverse_path = append(traverse_path,root.val)
  }
}


```

AVL树

balance Factor(平衡因子)

是它的左子树的高度减去他的右子树的高度（有时相反）。

balance factor = {-1,0,1}

通过旋转操作来进行平衡（四种）

左旋：左子树挂成左结点的右子树

右旋：右子树挂成右节点的左子树

左右旋、右左旋

AVL总结：平衡的二叉搜索树

​				每个结点存balance factor = {-1,0,1}

​				四种旋转操作

不足：结点需要存储额外信息，且调整次数频繁



近似平衡二叉树

红黑树是一种近似平衡的二叉搜索树（Binary Search Tree），它能够确保任何一个结点的左右子树的**高度差小于两倍**。

红黑树特点：

每个结点要么是红色，要么是黑色

根结点是黑色

每个叶子结点（nil结点，空结点）是黑色的。

不能有相邻接的两个红色结点

从任一结点到其每个叶子的所有路径都包含相同数目的黑色结点。



· AVL trees provide faster lookups than Red-Black Tree because they are more strictly balanced。

· Red-Black Trees provide faster insertion and removal operations than AVL trees as fewer rotations are done due to relatively relaxed balancing.

· AVL trees store balance factors or heights with each node，thus requires storage for an integer per node whereas Red-Black Trees requires only 1bit of information per node.

· Red-Black Trees are used in most of the language libraries like map,multimap,multisetin C++ where as AVL trees are used in databases where faster retrievals are required.