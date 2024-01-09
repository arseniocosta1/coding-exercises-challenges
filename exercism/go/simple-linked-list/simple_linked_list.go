package linkedlist

import "fmt"

type List struct {
	head *Element
	size int
}

type Element struct {
	data int
	next *Element
}

func New(elements []int) *List {
	list := &List{}

	for _, element := range elements {
		list.Push(element)
	}

	return list
}

func (l *List) Size() int {
	if l == nil {
		return 0
	}
	return l.size
}

func (l *List) Push(element int) {
	l.size++
	l.head = &Element{data: element, next: l.head}
}

func (l *List) Pop() (int, error) {
	if l == nil || l.head == nil {
		return 0, fmt.Errorf("List is empty")
	}

	element := l.head.data
	l.size--
	l.head = l.head.next

	return element, nil
}

func (l *List) Array() []int {
	ret := make([]int, l.Size(), l.Size())

	for i, cur := l.size-1, l.head; i >= 0; i, cur = i-1, cur.next {
		ret[i] = cur.data
	}

	return ret
}

func (l *List) Reverse() *List {
	ret := &List{}

	for i, cur := l.size-1, l.head; i >= 0; i, cur = i-1, cur.next {
		ret.Push(cur.data)
	}

	return ret
}
