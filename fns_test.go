package fn_test

import (
	"github.com/zjom/fn"
	"testing"
)

func TestFold(t *testing.T) {
	list := fn.NewList(1, 2, 3, 4, 5)
	sum := fn.Fold(list, 0, func(a, b int) int {
		return a + b
	})

	if sum != 15 {
		t.Errorf("Expected 15, got %d", sum)
	}
}

func TestMap(t *testing.T) {
	list := fn.NewList(1, 2, 3, 4, 5)
	squared := fn.Map(list, func(a int) int {
		return a * a
	})

	if squared.String() != "(1 4 9 16 25)" {
		t.Errorf("Expected (1 4 9 16 25), got %s", squared.String())
	}
}

func TestFilter(t *testing.T) {
	list := fn.NewList(1, 2, 3, 4, 5)
	evens := fn.Filter(list, func(a int) bool {
		return a%2 == 0
	})

	if evens.String() != "(2 4)" {
		t.Errorf("Expected (2 4), got %s", evens.String())
	}
}

func TestReduce(t *testing.T) {
	list := fn.NewList(1, 2, 3, 4, 5)
	sum := fn.Reduce(list, func(acc int, item int) int {
		return acc + item
	})
	if sum != 15 {
		t.Errorf("Expected 15, got %d", sum)
	}
}

func TestZip(t *testing.T) {
	list1 := fn.NewList(1, 2, 3)
	list2 := fn.NewList(4, 5, 6)
	zipped := fn.Zip(list1, list2)

	if zipped.String() != "((1 4) (2 5) (3 6))" {
		t.Errorf("Expected ((1 4) (2 5) (3 6)), got %s", zipped.String())
	}
}
