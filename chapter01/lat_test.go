package chapter01

import "testing"

func TestLatNormal(t *testing.T) {
	s := []any{"a", "sdxc", "x"}

	r := lat(s)

	if r != true {
		t.Error("wrong")
	}
}

func TestLatSliceWithNumber(t *testing.T) {
	s := []any{"a", 1, "x"}

	r := lat(s)

	if r != false {
		t.Error("wrong")
	}
}

func TestLatSliceWithSlice(t *testing.T) {
	s := []any{"a", []any{1, 2, 3}, "x"}

	r := lat(s)

	if r != false {
		t.Error("wrong")
	}
}
