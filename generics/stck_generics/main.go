package main

import "fmt"

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	deletedItem := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return deletedItem, true
}

func main() {
	intStack := Stack[int]{
		items: []int{},
	}
	intStack.Push(2)
	intStack.Push(3)
	fmt.Println(intStack.Pop())

	stringStack := Stack[string]{
		items: []string{},
	}
	stringStack.Push("sunny")
	stringStack.Push("Rahul")
	fmt.Println(stringStack.Pop())

}
