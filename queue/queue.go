package queue

import (
	"fmt"

)

type node struct {
	val  int
	next *node
	prev *node
}

type linkedList struct {
	head *node
	size int
}

type Queue struct {
	list *linkedList
}

func newList() *linkedList {
	return new(linkedList)
}

func newNode() *node {
	return new(node)
}

func (l *linkedList) printList() {
	n := l.head
	for i := 0; i < l.size; i++ {
		fmt.Println(n.val)
		n = n.next

	}

}

func (l *linkedList) addToHead(value int) {
	if l.head == nil {
		l.head = newNode()
		l.head.next = l.head
		l.head.prev = l.head
		l.head.val = value
		l.size = 1

	} else {
		n := newNode()
		n.val = value
		n.next = l.head
		n.prev = l.head.prev
		l.head.prev.next = n
		l.head.prev = n
		l.head = n
		l.size++

	}

}

func (l *linkedList) addToTail(value int) {
	if l.head == nil {
		l.head = newNode()
		l.head.next = l.head
		l.head.prev = l.head
		l.head.val = value
		l.size = 1
	} else {
		n := newNode()
		n.val = value
		n.next = l.head
		l.head.prev.next = n
		n.prev = l.head.prev
		l.head.prev = n
		l.size++
	}

}

func (l *linkedList) delFromHead() int {
	var n *node
	if l.size == 1 {
		l.head.next = nil
		l.head.prev = nil
		value := l.head.val
		l.head = nil
		l.size--
		return value
	} else {

		l.head.prev.next = l.head.next
		l.head.next.prev = l.head.prev
		n = l.head
		l.head = l.head.next
		n.next = nil
		n.prev = nil
		l.size--
		return n.val

	}

}

func NewQueue() *Queue {
	q := new(Queue)
	q.list = newList()
	return q

}

func (q *Queue) Enqueue(val int) {
	q.list.addToTail(val)

}

func (q *Queue) Dequeue() int {
	if q.list.size != 0 {

		value := q.list.delFromHead()
		return value
	} else {
		fmt.Println("Queue Empty")
		return 0
	}

}

func (q *Queue) PrintQueue() {
	if q.list.size != 1 {
		q.list.printList()
	} else {
		fmt.Println("Queue empty")
	}

}

