学习笔记

小结提纲



1、动态规划复习；附带递归、分治



2、多种情况的动态规划的状态转移方程串讲



3、进阶版动态规划的复习







递归复习：



\```go

//模板 

func recur() {

  //terminator

 if a = b {

  return

 }

 //process current logic

 process(level,param)

 

 //drill down

 recur(level+1,newParam)

​	

 //restore current status

}

\```



分治（分而治之）



用递归（经常用的习惯）



\```go

//模板

func divide_conquer(problem) {

  //recursion terminator

  if true {

​    return

  }

  //prepare data

  data := prepare_data(problem)

  subproblems := split_problem(prolem)

   

  //conquer subproblems

  subresult1 := divide_conquer(subproblems)

  subresult2 := divide_conquer(subproblems)

​		//process and generate the final result

  result = process_result(subresult1,subrresult2)

  //revert the current level states

}

\```



感触回顾



1、人肉递归低效、很累



2、找到最近最简的方法，将其拆解成可重复解决的问题



3、数学归纳法思维



递归状态树



![image-20200816235950912](/Users/liuyongchen/go/src/algorithm011-class01/Week_09/test1.png)







分治和动态规划的区别：每个环节寻找最优解，淘汰次优解—>变成动态规划。



动态规划



模板



\```go

//模板

func DP() {

 dp := [][]int{} //结构定义

 for i:=0...M {

  for i:=0...N {

   dp[i][j] = func[i][j] //重心状态转移方程

  }

 }

 return dp[M][N]

}

\```



关键点



拥有共性，找到重复子问题



差异性，最优子结构、中途可以淘汰次优解







例题：爬楼梯问题



状态转移公式：



res[n] = dfs(step1) + dfs(step2)



dp[i] = dp[i-1]+dp[i-2]



y, x = x+y, y



例题：不同路径



f(x,y) = f(x-1,y)+ f(x,y-1)



****dp[i][j]= dp[i-1][j] + dp[i][j-1]  >> dp[y][x]****



例题：打家劫舍



![image-20200817130004991](/Users/liuyongchen/go/src/algorithm011-class01/Week_09/image-20200817130004991.png)



例题：最小路径和



dp[i][j] = min(dp[i-1][j],dp[i][j-1]) + A [i][j]



例题：股票买卖



天数，是否拥有股票，最多交易次数



作业：![image-20200817131330479](/Users/liuyongchen/go/src/algorithm011-class01/Week_09/image-20200817131330479.png)







高阶DP



复杂度来源：状态拥有更多维度、状态方程更加复杂



内功，多练提炼状态定义，状态转移方程。



****例题****：爬楼梯问题



三级台阶？切片选项？不能重复走？"最小爬楼梯步数？"



****例题****：编辑距离



DFS（复杂度过高，很难满足最少操作数）、BFS（最好双端）



优化：字符串长度在m - n 之间。



DP：最后一个字符相同就都打掉，最后一个字符不同，打掉i 或打掉j 或都打掉的最小者