package queue

const capacity = 1000

type Queue struct {
	size int
	data [capacity]int
	front int
	back int // Current assignable slot
}

func (q *Queue) Add(value int) bool {
	if q.size >= capacity {
		return false
	}

	q.size++
	q.data[q.back] = value
	q.back = (q.back+1) % capacity
	return true
}
