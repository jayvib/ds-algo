package queue

import "container/list"

type Queue struct {
	l *list.List
}

func (q *Queue) Put(v interface{}) {
	q.l.PushBack(v)
}

func (q *Queue) Remove() (v interface{}) {
	e := q.l.Front()
	return q.l.Remove(e)
}
