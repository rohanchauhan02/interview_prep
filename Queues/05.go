package queues

import "fmt"

type circularQueue struct {
	data        []int
	size        int
	front, rear int
	count       int
}

func NewCircularQueue(k int) *circularQueue {
	return &circularQueue{
		data:  make([]int, k),
		size:  k,
		front: 0,
		rear:  -1,
		count: 0,
	}
}

func (q *circularQueue) Enqueue(val int) bool {
	if q.IsFull() {
		fmt.Println("Queue is full")
		return false
	}
	q.rear = (q.rear + 1) % q.size
	q.data[q.rear] = val
	q.count++
	return true
}

func (q *circularQueue) Dequeue() bool {
	if q.IsEmpty() {
		fmt.Println("Queue is empty")
		return false
	}
	q.front = (q.front + 1) % q.size
	q.count--
	return true
}

func (q *circularQueue) Front() int {
	if q.IsEmpty() {
		fmt.Println("Queue is empty")
		return -1
	}
	return q.data[q.front]
}

func (q *circularQueue) Rear() int {
	if q.IsEmpty() {
		fmt.Println("Queue is empty")
		return -1
	}
	return q.data[q.rear]
}

func (q *circularQueue) IsEmpty() bool {
	return q.count == 0
}

func (q *circularQueue) IsFull() bool {
	return q.count == q.size
}
