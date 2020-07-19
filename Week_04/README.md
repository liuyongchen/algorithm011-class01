学习笔记

深度优先搜索和广度优先搜索

遍历所有的点

搜索遍历

每个节点访问一次且仅访问一次

深度优先 depth first search

示例代码：递归模板（二叉树，多叉树）

```go
//多叉树递归
visited = set()
func dfs(node,visited) {
  visited = append(visited,node)
  #process current node here.
  ...
  for next_node in node.children() {
    if not next_node in visited {
      dfs(next node,visited)
    }
  }
}
//DFS 非递归模板
func dfs(root *TreeNode) {
  if root == nil {
    return nil
  }
  visited, stack := []*TreeNode{}, []*TreeNode{root}
  for stack != nil {
    node := stack[:len(stack)-1]
    visited = append(visited,node)
    process(node)
    nodes = generate_related_nodes(node)
    stack = append(stack,nodes)
  }
}
```



广度优先 breadth first search

示例代码：

```go
type TreeNode struct {
  Val int
  Left *TreeNode
  Right *TreeNode
}

func bfs(graph,start,end)
queue := []*TreeNode{start}
visited := []*TreeNode{start}
for queue != nil {
  node := queue[0]
  visited = append(visited,node)
  process(node)
  nodes = generate_related_nodes(node)
  queue = append(queue,nodes...)
}
  
// other processing work
}
```



优先级优先（现实场景（启发式搜索（深度学习）））

相关面试题：
1、二叉树的层序遍历
2、最小基因变化
3、括号生成
4、在每个树行中找最大值





**贪心算法**

每一步选择中都采取当前状态下最好或最优（最有利）的选择，从而希望导致结果是全局最好或最优的算法

每个子问题的解决方案都做出选择，不能回退

动态规划：会保存以前的运算结果，并根据以前的结果对当前进行选择，有回退功能

**对比**

贪心：当下做局部最优判断

回溯：能够回退

动态规划：最优判断+回退

解决问题：最优化问题：最小生成树、求哈夫曼编码；最优子结构

高效，结果接近最优结果。

**相关问题：**

零钱兑换、433. 分发并给饼干、最佳买卖股票的时机、跳跃游戏、模拟行走机器人

**二分查找**

1、目标函数单调性（单调递增或者单调递减）

2、存在上下界（bounded）

3、能够通过索引访问（index access）

**代码模板：**

```go
left，right := 0,lenth-1
for left <= right {
	mid = left+(right-left)/2
	if array[mid] == target {
		//find the target!!
		break or return result
	}else if array[mid] < target{
		left = mid + 1
	}else {
		right = mid - 1
	}
}
```



相关面试题：

sqrt X的算数平方根、有效的完全平方数、搜索旋转排序数组、搜索二维矩阵

