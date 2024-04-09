package main

import "fmt"

type MyQueue[T any] struct {
	queue []T
}

func Constructor[T any](q []T) MyQueue[T] {
	return MyQueue[T]{queue: q}
}

func (this *MyQueue[T]) Push(x T) {
	if this == nil || this.queue == nil {
		return
	}

	this.queue = append(this.queue, x)
}

func (this *MyQueue[T]) Pop() T {
	val := this.queue[0]

	this.queue = this.queue[1:]
	return val
}

func (this *MyQueue[T]) Peek() T {
	return this.queue[0]
}

func (this *MyQueue[T]) Empty() bool {
	if this == nil {
		return false
	}

	return this.Length() == 0
}

func (this *MyQueue[T]) Length() int {
	return len(this.queue)
}

func countStudents(students []int, sandwiches []int) int {
	studQ := Constructor[int](students)
	sandwQ := Constructor[int](sandwiches)

	for !studQ.Empty() {
		somoOneEat := false
		l := studQ.Length()

		for range l {
			topStudent := studQ.Pop()
			if topStudent == sandwQ.Peek() {
				sandwQ.Pop()
				somoOneEat = true
			} else {
				studQ.Push(topStudent)
			}
		}

		if !somoOneEat {
			return studQ.Length()
		}
	}

	return 0
}

func main() {
	students := []int{1, 1, 0, 0}
	sandwiches := []int{0, 1, 0, 1}
	fmt.Println(countStudents(students, sandwiches))

	students2 := []int{1, 1, 1, 0, 0, 1}
	sandwiches2 := []int{1, 0, 0, 0, 1, 1}
	fmt.Println(countStudents(students2, sandwiches2))

	students3 := []int{1, 1, 0, 0, 1, 1}
	sandwiches3 := []int{0, 1, 0, 1, 0, 0}
	fmt.Println(countStudents(students3, sandwiches3))

	students4 := []int{1, 1, 1, 0, 0, 1, 1}
	sandwiches4 := []int{1, 0, 1, 0, 1, 0, 0}
	fmt.Println(countStudents(students4, sandwiches4))

}
