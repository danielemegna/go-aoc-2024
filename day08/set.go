package day08

import (
	"maps"
	"slices"
)

type Set[T comparable] struct {
	items map[T]bool
}

func NewSet[T comparable]() Set[T] {
	return Set[T]{
		items: make(map[T]bool),
	}
}

func (s *Set[T]) Add(item T) {
	s.items[item] = true
}

func (s *Set[T]) ToSlice() []T {
	return slices.Collect(maps.Keys(s.items))
}
