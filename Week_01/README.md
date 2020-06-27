学习笔记

#学号:G20200343080049

学习总结

**1、算法处理思想**

a.升维度

b.空间换时间

**2、三分看视频，七分练习**

难点反复看

摒弃旧习惯

​    a.不要死磕

​    b.五毒神掌（死记硬背优秀代码）

​    c.多看高手代码（国际版高票回答）

五步学习法（五毒神掌）

​    a.多过遍数，而不是每一次花很多时间，难点重点题多做几次，面试前再做一次。

​    b.多看高票和高质量题解，在学习工具，学习知识，不是在考试，不要耻于学习别人的 优秀题解。

​    c.职业选手反复练习同一动作，普通玩家才只打排位不训练细节。

五步：

第一步：5-15分钟，读题思考；

第二步：直接看题解，多解法比较优劣；

第三步：背诵默写好的解法，立刻自己写leatcode 提交

第四步：过一天后再重复做题，过一周后再重复做题

第五步：面试前恢复性训练



PS：
1、刷题心得
1两数之和、15三数之和
使用双指针、三指针解题时控制好边界，做好对边界的判断保护程序避免index溢出；
分支条件存在 && 并关系时，边界判断放在表达式前面，避免程序因index溢出引发异常；

2、视频作业
作业：queue 和 priority queue源码的分析。

Go语言 queue 和 priority queue源码分析

Go语言中 queue 和 priority实现了同样一组API，

queue源码分析
路径：go-datastructures/queue/queue.go 

Queue 结构体中包含

waiters （[]*sema） //封装了 get 和 put 两个方法，用于获取和弹出sema对象。

item  （[]interface{}）//用于保存数据，封装了get 和 getutil 两个方法，用于获取所需的数据

lock   sync.mutex //用于对并发操作可能产生冲突的数据上锁

disposed  bool  //指示队列是否调用了dispose，调用了dispose 的队列将不可用并发挥一个错误

API

func (q *Queue) Put(items ...interface{}) error

向队列中推入指定数据；
检查disposed字段是否为true，如果是返回错误，如果不是推数据进队列；

不理解：等待组的处理


{

    for {
    
    	sema := q.waiters.get()
    	
    	if sema == nil {
    	
    		break
    		
    	}
    	
    	sema.response.Add(1)
    	
    	sema.wg.Done()
    	
    	sema.response.Wait()
    	
    	if len(q.items) == 0 {
    	
    		break
    		
    	}
    	
    }

func (q *Queue) Get(number int64) ([]interface{}, error)

如果number < 1 返回空切片
如果disposed 为true 返回disposed错误
如果items的长度为0，将创建新的sema对象并推入waiters切片，等待数据传入
然后将所需number数量的数据传入items
如果items的长度不为0，直接将number传入items。
最终返回新的items 该方法不涉及循环操作，也不涉及新的额外空间的占用，
时间复杂度和空间复杂度均为O(1)

func (q *Queue) TakeUntil(checker func(item interface{}) bool) ([]interface{}, error)

通过checker 获取数据切片；
如果checker为空直接返回空；如果disposed 为true 返回disposed错误
调用items.getUntil()方法，获取匹配数据的切片并返回。
该方法的时间和空间复杂度取决于items.getUntil()方法，
由于该方法需要遍历整个切片，所以时间复杂度为O(n)
由于获取元素的过程，会涉及切片的扩容量；最差空间复杂度也是O(n)
该方法不会因为queue 中没有元素而等待。

func (q *Queue) Empty() bool

通过判断items字段的长度是否为0，判断queue是否为空；

func (q *Queue) Len() int64

通过判断items的长度，判断queue的长度。

func (q *Queue) Disposed() bool

判断disposed的值是否为true

func (q *Queue) Dispose()

修改disposed的值为true

func New(hint int64) *Queue

创建新的队列，队列长度由hint决定

func ExecuteInParallel(q *Queue, fn func(interface{}))

调用方法接走queue里的数据，当所有goroutine 被耗尽后，queue将被释放。


priority queue 源码分析
路径：go-datastructures/queue/priority_queue.go 

PriorityQueue 结构体中包含五个字段

waiters []*sema //与queue中类型一致

items priorityitems 实质是[]item 包含get 和 insert两个方法，用于获取和插入元素

lock  sync.Mutex

disposeLock sync.Mutex

disposed bool

API

func (pq *PriorityQueue) Put(items ...Item) error

将元素推入队列，如果添加元素为空，不做任何操作返回
循环遍历传入items，将元素添加进pq.items

func (pq *PriorityQueue) Get(number int) ([]Item, error)

从queue中获取number数量的数据，
优先级的划分在pq.items.insert()方法中完成了优先级的判断，
高优先级的数据会被插入到切片的最前面。

func (pq *PriorityQueue) Peek() Item

查看queue中第一个数据，当queue中为空的时候，返回nil

func (pq *PriorityQueue) Empty() bool

根据items的长度判断queue是否为空，当长度为0的时候，结果为true

func (pq *PriorityQueue) Len() int

根据items的长度获取对垒长度

func (pq *PriorityQueue) Disposed() bool

读取disposed 字段的值

func (pq *PriorityQueue) Dispose()

设置disposed 字段的值为true 重置所有切片为空触发垃圾回收，释放queue

func NewPriorityQueue(hint int) *PriorityQueue 

新建优先队列，初始化items 切片，设置切片长度为hint