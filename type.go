package stack

import (
	"github.com/yms2772/stack/comparator"
	"golang.org/x/exp/constraints"
)

type stack[T constraints.Ordered] struct {
	data []T
}

type Stack[T constraints.Ordered] interface {
	Clean()
	Data() []T
	Empty() (empty bool)
	Index(index int) (data T, err error)
	Peek() (data T)
	Pop() (data T)
	Push(data T)
	Remove(index int)
	Search(data T) (index int)
	Size() (length int)
	Sort(c ...comparator.Comparator)
}
