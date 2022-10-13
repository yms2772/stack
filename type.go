package stack

import (
	"github.com/yms2772/stack/comparator"
	"golang.org/x/exp/constraints"
)

type stack[T constraints.Ordered] struct {
	data []T
}

type Stack[T constraints.Ordered] interface {
	Pop() (data T)
	Push(data T)
	Clean()
	Peek() (data T)
	Empty() (empty bool)
	Search(data T) (index int)
	Size() (length int)
	Sort(c ...comparator.Comparator)
	Index(index int) (data T, err error)
	Remove(index int)
}
