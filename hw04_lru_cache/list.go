package hw04lrucache

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type list struct {
	len   int
	front *ListItem
	back  *ListItem
}

func (l list) Len() int {
	return l.len
}

func (l list) Front() *ListItem {
	return l.front
}

func (l list) Back() *ListItem {
	return l.back
}

func (l *list) PushFront(v interface{}) *ListItem {
	node := new(ListItem)
	node.Value = v

	if l.front == nil {
		l.front = node
		l.back = node
	} else {
		node.Next = l.front
		l.front.Prev = node
		l.front = node
	}

	l.len++

	return node
}

func (l *list) PushBack(v interface{}) *ListItem {
	node := new(ListItem)
	node.Value = v

	if l.front == nil {
		l.front = node
		l.back = node
	} else {
		l.back.Next = node
		node.Prev = l.back
		l.back = node
	}

	l.len++

	return node
}

func (l *list) Remove(i *ListItem) {
	if i == nil {
		return
	}

	switch {
	case i.Prev != nil && i.Next != nil:
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
	case i.Prev == nil && i.Next != nil:
		l.front = i.Next.Prev
	case i.Prev != nil && i.Next == nil:
		l.back = i.Prev
		l.back.Next = nil
	case i.Prev == nil && i.Next == nil:
		l.front = nil
		l.back = nil
	}
	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	l.PushFront(i.Value)
	l.Remove(i)
}

func NewList() List {
	return new(list)
}
