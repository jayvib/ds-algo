package queue

type Queue struct {
	data []string
}

func (q *Queue) Push(s string) {
	q.data = append(q.data, s)
}
func (q *Queue) Front() (f string) {
	f, q.data = q.data[0], q.data[1:]
	return
}
func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}
