package Week_01

//链表实现，简洁清楚
type MyCircularDeque struct {
	head  *Node
	tail  *Node
	len   int
	count int
}

type Node struct {
	Next *Node
	Pre  *Node
	Val  int
}

func Constructor(k int) MyCircularDeque {
	head := Node{
		Next: nil,
		Pre:  nil,
		Val:  -1,
	}
	tail := Node{
		Next: nil,
		Pre:  nil,
		Val:  -1,
	}
	head.Next = &tail
	tail.Pre = &head
	deque := MyCircularDeque{
		head:  &head,
		tail:  &tail,
		len:   k,
		count: 0,
	}
	return deque
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	temp := this.head.Next
	tempNode := Node{
		Next: temp,
		Pre:  this.head,
		Val:  value,
	}
	this.head.Next = &tempNode
	temp.Pre = &tempNode
	this.count++
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	temp := this.tail.Pre
	tempNode := Node{
		Next: this.tail,
		Pre:  temp,
		Val:  value,
	}
	this.tail.Pre = &tempNode
	temp.Next = &tempNode
	this.count++
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	deleteTemp := this.head.Next
	this.head.Next = deleteTemp.Next
	deleteTemp.Next.Pre = this.head
	deleteTemp.Next, deleteTemp.Pre = nil, nil
	this.count--
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	deleteTemp := this.tail.Pre
	this.tail.Pre = deleteTemp.Pre
	deleteTemp.Pre.Next = this.tail
	deleteTemp.Next, deleteTemp.Pre = nil, nil
	this.count--
	return true
}

func (this *MyCircularDeque) GetFront() int {
	return this.head.Next.Val
}

func (this *MyCircularDeque) GetRear() int {
	return this.tail.Pre.Val
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.head.Next == this.tail
}

func (this *MyCircularDeque) IsFull() bool {
	return this.count == this.len
}
