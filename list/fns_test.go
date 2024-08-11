package list_test

import (
	"github.com/zjom/fn"
	"github.com/zjom/fn/list"
	"testing"
)

func TestFold(t *testing.T) {
	xs := fn.NewList(1, 2, 3, 4, 5)
	sum := list.Fold(xs, 0, func(a, b int) int {
		return a + b
	})

	if sum != 15 {
		t.Errorf("Expected 15, got %d", sum)
	}
}

func TestMap(t *testing.T) {
	xs := fn.NewList(1, 2, 3, 4, 5)
	squared := list.Map(xs, func(a int) int {
		return a * a
	})

	if squared.String() != "(1 4 9 16 25)" {
		t.Errorf("Expected (1 4 9 16 25), got %s", squared.String())
	}
}

func TestFilter(t *testing.T) {
	xs := fn.NewList(1, 2, 3, 4, 5)
	evens := list.Filter(xs, func(a int) bool {
		return a%2 == 0
	})

	if evens.String() != "(2 4)" {
		t.Errorf("Expected (2 4), got %s", evens.String())
	}
}

func TestReduce(t *testing.T) {
	xs := fn.NewList(1, 2, 3, 4, 5)
	sum := list.Reduce(xs, func(acc int, item int) int {
		return acc + item
	})
	if sum != 15 {
		t.Errorf("Expected 15, got %d", sum)
	}
}

func TestZip(t *testing.T) {
	xs1 := fn.NewList(1, 2, 3)
	xs2 := fn.NewList(4, 5, 6)
	zipped := list.Zip(xs1, xs2)

	if zipped.String() != "((1 4) (2 5) (3 6))" {
		t.Errorf("Expected ((1 4) (2 5) (3 6)), got %s", zipped.String())
	}
}
