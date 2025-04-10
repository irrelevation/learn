package linkedlist

type LinkedList[T any] interface {
	len() int
	add(val T, i int)
	remove(i int) T
	append(i T)
	prepend(i T)
	get(i int) T
}

type node[T any] struct {
	val        T
	prev, next *node[T]
}

type linkedList[T any] struct {
	length     int
	head, tail *node[T]
}

func (l *linkedList[T]) len() int {
	return l.length
}

func (l *linkedList[T]) getNode(i int) *node[T] {
	cursor := l.head
	for ; i > 0; i-- {
		cursor = cursor.next
	}
	return cursor
}

func (l *linkedList[T]) get(i int) T {
	return l.getNode(i).val
}

func (l *linkedList[T]) add(val T, i int) {
	if i == 0 {
		l.prepend(val)
	} else if i == l.length {
		l.append(val)
	} else {
		prev := l.getNode(i)
		n := &node[T]{val, prev, prev.next}
		prev.next = n
		n.next.prev = n
		l.length++
	}
}

func (l *linkedList[T]) pop() T {
	n := l.tail
	l.tail = n.prev
	l.tail.next = nil
	n.prev = nil
	l.length--
	return n.val
}

func (l *linkedList[T]) shift() T {
	n := l.head
	l.head = n.next
	l.head.prev = nil
	n.next = nil
	l.length--
	return n.val
}

func (l *linkedList[T]) remove(i int) T {
	if i == 0 {
		return l.shift()
	}
	if i == l.length {
		return l.pop()
	}
	n := l.getNode(i)
	n.prev.next = n.next
	n.next.prev = n.prev
	n.next = nil
	n.prev = nil
	l.length--
	return n.val
}

func (l *linkedList[T]) append(val T) {
	n := &node[T]{val, l.tail, nil}
	l.tail.next = n
	l.tail = n
	l.length++
}

func (l *linkedList[T]) prepend(val T) {
	n := &node[T]{val, nil, l.head}
	l.head.prev = n
	l.head = n
	l.length++
}

func New[T any]() linkedList[T] {
	return linkedList[T]{length: 0}
}
