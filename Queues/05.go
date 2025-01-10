package queues

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

func (q *circularQueue) Enqueue() {}

func (q *circularQueue) Dequeue() {}

func (q *circularQueue) Front() {}

func (q *circularQueue) Rear() {}

func (q *circularQueue) IsEmpty() bool {
	return q.count == 0
}

func (q *circularQueue) IsFull() bool {
	return q.count == q.size
}
