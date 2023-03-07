package stack

import (
	"errors"
	"sort"

	"github.com/yms2772/stack/comparator"
	"golang.org/x/exp/constraints"
)

// New stack
func New[T constraints.Ordered]() Stack[T] {
	return &stack[T]{
		data: []T{},
	}
}

// Clean stack
func (s *stack[T]) Clean() {
	s.data = s.data[:0]
}

// Data is print all data
func (s *stack[T]) Data() []T {
	return s.data
}

// Empty stack
func (s *stack[T]) Empty() (empty bool) {
	return len(s.data) == 0
}

// Index stack
func (s *stack[T]) Index(index int) (data T, err error) {
	if index >= s.Size() {
		return data, errors.New("overflow")
	}

	return s.data[index-1], nil
}

// Pop from stack
func (s *stack[T]) Pop() (data T) {
	if !s.Empty() {
		data = s.data[len(s.data)-1]
		s.data = s.data[:len(s.data)-1]

		return data
	}

	return data
}

// Peek from stack
func (s *stack[T]) Peek() (data T) {
	if !s.Empty() {
		return (s.data)[len(s.data)-1]
	}

	return data
}

// Push to stack
func (s *stack[T]) Push(data T) {
	s.data = append(s.data, data)
}

// Remove element of stack
func (s *stack[T]) Remove(index int) {
	s.data = append(s.data[:index], s.data[index+1:]...)
}

// Search stack
func (s *stack[T]) Search(data T) (index int) {
	if s.Empty() {
		return 0
	}

	for i, item := range s.data {
		if item == data {
			return i
		}
	}

	return 0
}

// Size stack
func (s *stack[T]) Size() (length int) {
	return len(s.data)
}

// Sort stack by string
// Default: Ascending
func (s *stack[T]) Sort(c ...comparator.Comparator) {
	sort.Slice(s.data, func(i, j int) bool {
		if len(c) != 0 {
			switch c[0] {
			case comparator.ASC:
				return s.data[i] < s.data[j]
			case comparator.DESC:
				return s.data[i] > s.data[j]
			}
		}

		return s.data[i] < s.data[j]
	})
}
