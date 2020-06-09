package main

import "fmt"

type MinStack struct {
	stack []item
}

type item struct {
	min, x int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

// 入栈
func (this *MinStack) Push(x int)  {
	min := x
	// 判断此时栈中是否有其他数据以及最小值是否小于当前值
	if len(this.stack) > 0 && this.GetMin() < x {
		min = this.GetMin()
	}
	this.stack = append(this.stack, item{min, x})
}

// 出栈
func (this *MinStack) Pop()  {
	this.stack = this.stack[:len(this.stack) - 1]
}

// 获取顶部
func (this *MinStack) Top() int {
	return this.stack[len(this.stack) - 1].x
}

// 获取最小值
func (this *MinStack) GetMin() int {
	return this.stack[len(this.stack) - 1].min
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
func main() {
	minStack := Constructor()
	minStack.Push(-2)
	fmt.Println(minStack.stack)

	minStack.Push(0)
	fmt.Println(minStack.stack)

	minStack.Push(-3)
	fmt.Println(minStack.stack)

	param_4 := minStack.GetMin()
	fmt.Println(param_4)

	minStack.Pop()
	fmt.Println(minStack.stack)

	param_3 := minStack.Top()
	fmt.Println(param_3)

	minStack.GetMin()
}
