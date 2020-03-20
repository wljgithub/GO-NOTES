package simplemath

import (
	"testing"
)

func testSqrt(t *testing.T) {
	r := Sqrt(16)
	if r != 5 {
		t.Errorf("sqrt(16) failed ,Got %v expect 4", r)
	}
}
