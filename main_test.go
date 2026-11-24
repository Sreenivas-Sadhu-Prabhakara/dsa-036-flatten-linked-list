package main

import (
	"reflect"
	"testing"
)

func TestFlatten(t *testing.T) {
	a := column(5, 7, 8, 30)
	b := column(10, 20)
	c := column(19, 22, 50)
	a.Next, b.Next = b, c
	got := flatToSlice(Flatten(a))
	want := []int{5, 7, 8, 10, 19, 20, 22, 30, 50}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Flatten = %v, want %v", got, want)
	}
}

func TestFlattenEdges(t *testing.T) {
	if Flatten(nil) != nil {
		t.Error("nil flattens to nil")
	}
	one := column(1, 2, 3)
	if got := flatToSlice(Flatten(one)); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("single column unchanged, got %v", got)
	}
	// Columns of length 1 behave like a plain sorted merge.
	x, y := column(4), column(2)
	x.Next = y
	if got := flatToSlice(Flatten(x)); !reflect.DeepEqual(got, []int{2, 4}) {
		t.Errorf("got %v, want [2 4]", got)
	}
}
