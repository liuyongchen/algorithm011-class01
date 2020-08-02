学习笔记

递归代码模板

```go
//terminator
//current process logic
//drill down
//restore current status
```

分治

```go
//大问题分解成若干小问题
//terminatro
//prepare date (拆分子问题)
//conquer subproblem （递归解决子问题）
//process and generate the final result （合并结果）
//revert the current level status （恢复当前层）
```



动态规划步骤
1、找到重复子问题
2、存储中间状态opt[i]
3、提炼递推公式（状态转移方程）

感触
动态规划问题，找重复子问题和提炼递推公式比较抽象，需要多加练习；重复刷题，培养计算机思维。

对于提炼动态规划递推公式，最入门需要考虑升维，通过升维，找到解决问题的方法。