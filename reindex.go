package reindex

import "fmt"

type List[T any] struct {
	data []T
}

func (l *List[T]) Add(index int, to T) {
	if index < 0 {
		panic(fmt.Errorf("index is negative: %d", index))
	}
	if size := len(l.data); size <= index { // change len of slice
		l.data = append(l.data, make([]T, index-size+1)...)
	}
	l.data[index] = to
}

func (l List[T]) Get(index int) (_ T, err error) {
	if len(l.data) <= index {
		err = fmt.Errorf("out of range: %d <= %d", len(l.data), index)
		return
	}
	return l.data[index], nil
}
