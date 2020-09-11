package main

import (
	"fmt"
	list "github.com/jayvib/ds-algo/books/common-sense-guide-to-ds-algo/data-structure/single_linked_list"
	"regexp"
)

func NewEmailList() *EmailList {
	return &EmailList{}
}

type EmailList struct {
	list.LinkedList
}

func (e *EmailList) Filter(re *regexp.Regexp) *EmailList {
	newEmailList := NewEmailList()
	e.Traverse(func(i int, n *list.Node)bool{
		if re.MatchString(n.Value().(string)) {
			newEmailList.PushBack(n.Value())
		}
		return false
	})
	return newEmailList
}

type Iterator interface {
	Next() (email string, done bool)
	Close()
}

func newIter(el *EmailList) *iter {
	i := &iter{
		e:         el,
		emailChan: make(chan string),
	}
	i.run()
	return i
}

type iter struct {
	e *EmailList
	emailChan chan string
	isClose bool
}

func (i *iter) run() {
	go func() {
		defer close(i.emailChan)
		i.e.Traverse(func(index int, n *list.Node)bool{
			if i.isClose {
				return true
			}
			i.emailChan <- n.Value().(string)
			return false
		})
	}()
}

func (i *iter) Next() (email string, done bool) {
	v, ok := <-i.emailChan
	if !ok {
		return "", true
	}
	return v, false
}

func (i *iter) Close() {
	i.isClose = true
}

func (e *EmailList) Iterator() Iterator {
	return newIter(e)
}

func (e *EmailList) Print() {
	e.Traverse(func(i int, n *list.Node)bool{
		fmt.Println(n.Value())
		return false
	})
}

func fillInEmails(l *EmailList, emails []string) {
	for _, email := range emails {
		l.PushBack(email)
	}
}

func main() {
	el := NewEmailList()
	fillInEmails(el, emails)
	startsWithHEmailList := el.
		Filter(regexp.MustCompile("^h")).
		Filter(regexp.MustCompile("net$"))

	itr := startsWithHEmailList.Iterator()

	fmt.Println(itr.Next())
	fmt.Println(itr.Next())
	itr.Close()
}

