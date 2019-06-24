package main

import "fmt"

type MyCircularQueue struct {
	head *node
	tail *node
	size int
}

type node struct {
	next  *node
	value interface{}
}

//初始化队列
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		size: k,
	}
}

//入队
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.size == 0 {
		return false
	} else {
		n := &node{
			value: value,
		}
		if this.tail == nil {
			this.head = n
			this.tail = n
			this.size--
		} else {
			this.tail.next = n
			this.tail = n
			this.size--
		}
		return true
	}
}

//出队
func (this *MyCircularQueue) DeQueue() bool {
	if this.head == nil {
		return false
	} else {
		if this.head == this.tail {
			this.head, this.tail = nil, nil
			this.size++
		} else {
			this.head = this.head.next
			this.size++
		}
		return true
	}
}

//取队头
func (this *MyCircularQueue) Front() int {
	if this.head == nil {
		return -1
	}
	if v, ok := this.head.value.(int); ok {
		return v
	} else {
		return -1
	}
}

//取队尾
func (this *MyCircularQueue) Rear() int {
	if this.tail == nil {
		return -1
	}
	if v, ok := this.tail.value.(int); ok {
		return v
	} else {
		return -1
	}
}

//判断队列空
func (this *MyCircularQueue) IsEmpty() bool {
	if (this.head == nil || this.tail == nil) && this.size != 0 {
		return true
	}
	return false
}

//判断队列满
func (this *MyCircularQueue) IsFull() bool {
	if this.head != nil && this.tail != nil && this.size == 0 {
		return true
	}
	return false
}

func main() {
	var k int
	fmt.Println("输入队列长度：")
	fmt.Scanf("%d\n", &k)
	obj := Constructor(k)
	var flag int
	for {
		fmt.Println("输入对应的操作：")
		fmt.Scanf("%d\n", &flag)
		if flag == 1 {
			var value int
			fmt.Scanf("%d\n", &value)
			param_1 := obj.EnQueue(value)
			fmt.Println(param_1)
		} else if flag == 2 {
			param_2 := obj.DeQueue()
			fmt.Println(param_2)
		} else if flag == 3 {
			param_3 := obj.Front()
			fmt.Println(param_3)
		} else if flag == 4 {
			param_4 := obj.Rear()
			fmt.Println(param_4)
		} else if flag == 5 {
			param_5 := obj.IsEmpty()
			fmt.Println(param_5)
		} else if flag == 6 {
			param_6 := obj.IsFull()
			fmt.Println(param_6)
		} else {
			break
		}
	}
}
