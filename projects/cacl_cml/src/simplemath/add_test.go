package simplemath

import (
	"testing"
)

func TestAdd1(t *testing.T) {
	r := add(1, 2)
	if r != 3 {

		t.Errorf("add(1,2) failed. Got %d ,expected 3.", r)
	}
}
