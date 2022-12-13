package queue

import (
	"container/list"
	"fmt"
)

type Queue[T any] struct {
	list *list.List
}

func New[T any]() Queue[T] {
	return Queue[T]{
		list: list.New(),
	}
}

func (q Queue[T]) Pop() (T, error) {
	val := (q).list.Front()
	q.list.Remove(val)

	switch t := val.Value.(type) {
	case T:
		return val.Value.(T), nil
	default:
		return *new(T), fmt.Errorf("unknown type in queue: %v", t)
	}
}

func (q Queue[T]) Append(element T) {
	q.list.PushBack(element)
}

func (q Queue[T]) Size() int {
	return q.list.Len()
}
