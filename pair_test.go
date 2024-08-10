package fn_test

import (
	"github.com/zjom/fn"
	"testing"
)

func TestPair(t *testing.T) {
	p := fn.NewPair(1, 2)
	if p.First() != 1 {
		t.Errorf("Expected %d, got %d", 1, p.First())
	}
	if p.Second() != 2 {
		t.Errorf("Expected %d, got %d", 2, p.Second())
	}
}
