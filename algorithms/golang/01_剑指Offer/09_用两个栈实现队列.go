package main

import "fmt"

/*
 方法一：双栈
*/

type CQueue struct {
	inStack  []int
	outStack []int
}

func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.inStack = append(this.inStack, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.outStack) == 0 {
		if len(this.inStack) == 0 {
			return -1
		}
		for len(this.inStack) > 0 {
			this.outStack = append(this.outStack, this.inStack[len(this.inStack)-1])
			this.inStack = this.inStack[:len(this.inStack)-1]
		}
	}
	value := this.outStack[len(this.outStack)-1]
	this.outStack = this.outStack[:len(this.outStack)-1]
	return value
}

func main() {
	obj := Constructor()
	fmt.Println(obj.DeleteHead() == -1)
	obj.AppendTail(5)
	obj.AppendTail(2)
	fmt.Println(obj.DeleteHead() == 5)
	fmt.Println(obj.DeleteHead() == 2)
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
