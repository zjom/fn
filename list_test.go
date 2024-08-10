package fn_test

import (
	"github.com/zjom/fn"
	"testing"
)

func TestList(t *testing.T) {
	l := fn.NewList(1, 2, 3)
	if l.Head() != 1 {
		t.Errorf("Expected 1, got %v", l.Head())
	}

	if l.Rest().Head() != 2 {
		t.Errorf("Expected 2, got %v", l.Rest().Head())
	}

	if l.Rest().Rest().Head() != 3 {
		t.Errorf("Expected 3, got %v", l.Rest().Rest().Head())
	}

	if l.Rest().Rest().Rest() != nil {
		t.Errorf("Expected nil, got %v", l.Rest().Rest().Rest())
	}

	l = l.Cons(0)
	if l.Head() != 0 {
		t.Errorf("Expected 0, got %v", l.Head())
	}

	if l.Rest().Head() != 1 {
		t.Errorf("Expected 1, got %v", l.Rest().Head())
	}

	slice := l.ToSlice()
	if len(slice) != 4 {
		t.Errorf("Expected 4, got %v", len(slice))
	}

	if slice[0] != 0 {
		t.Errorf("Expected 0, got %v", slice[0])
	}
}
