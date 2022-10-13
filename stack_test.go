package stack_test

import (
	"reflect"
	"testing"

	"github.com/yms2772/stack"
	"github.com/yms2772/stack/comparator"
	"golang.org/x/exp/rand"
)

func Test_New(t *testing.T) {
	s := stack.New[int]()
	stackType := reflect.ValueOf(s).Type().String()

	if stackType != "*stack.stack[int]" {
		t.Errorf("stack.New() = %s; want %s", stackType, "*stack.stack[int]")
	}
}

func TestStack_Push(t *testing.T) {
	s := stack.New[int]()
	for i := 0; i < 100; i++ {
		s.Push(i)
	}

	if s.Size() != 100 {
		t.Errorf("stack.Push() = %d; want %d", s.Size(), 100)
	}
}

func BenchmarkStack_Push(b *testing.B) {
	s := stack.New[int]()
	for i := 0; i < b.N; i++ {
		for j := 0; j < b.N; j++ {
			s.Push(j)
		}
	}
}

func BenchmarkSlice_Append(b *testing.B) {
	var s []int
	for i := 0; i < b.N; i++ {
		for j := 0; j < b.N; j++ {
			s = append(s, j)
		}
	}
}

func TestStack_Size(t *testing.T) {
	s := stack.New[int]()

	if s.Size() != 0 {
		t.Errorf("stack.Push() = %d; want %d", s.Size(), 0)
	}
}

func TestStack_Search(t *testing.T) {
	s := stack.New[int]()
	s.Push(9)
	s.Push(8)
	s.Push(7)

	if s.Search(7) != 2 {
		t.Errorf("stack.Search() = %d; want %d", s.Search(7), 2)
	}
}

func TestStack_Sort(t *testing.T) {
	s := stack.New[int]()
	randSlice := make([]int, 100)
	for i := 0; i < 100; i++ {
		randSlice[i] = i
	}

	for i := range randSlice {
		j := rand.Intn(i + 1)
		randSlice[i], randSlice[j] = randSlice[j], randSlice[i]
	}

	for _, item := range randSlice {
		s.Push(item)
	}

	s.Sort(comparator.ASC)
	if s.Peek() != 99 {
		t.Errorf("stack.Sort() = %d; want %d", s.Peek(), 99)
	}

	s.Sort(comparator.DESC)
	if s.Peek() != 0 {
		t.Errorf("stack.Sort() = %d; want %d", s.Peek(), 0)
	}

	s.Sort()
	if s.Peek() != 99 {
		t.Errorf("stack.Sort() = %d; want %d", s.Peek(), 99)
	}
}

func TestStack_Clean(t *testing.T) {
	s := stack.New[int]()
	for i := 0; i < 100; i++ {
		s.Push(i)
	}

	s.Clean()
	if s.Size() != 0 {
		t.Errorf("stack.Size() = %d; want %d", s.Size(), 0)
	}
}

func TestStack_Index(t *testing.T) {
	s := stack.New[int]()
	for i := 0; i < 100; i++ {
		s.Push(i)
	}

	if data, err := s.Index(10); err != nil {
		t.Errorf("stack.Index() = %s", err.Error())
	} else {
		if data != 9 {
			t.Errorf("stack.Index() = %d; want %d", data, 9)
		}
	}
}

func TestStack_Remove(t *testing.T) {
	s := stack.New[int]()
	for i := 0; i < 100; i++ {
		s.Push(i)
	}

	s.Remove(0)
	s.Remove(98)

	if s.Size() != 98 {
		t.Errorf("stack.Size() = %d; want %d", s.Size(), 98)
	}
}
