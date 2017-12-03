package main

import "testing"

func TestModifyArray(t *testing.T) {
	t.Error("Oops 1")
	arr := [...]int{0}
	modifyArray(arr, -1)
	// Arrays are copied by value, func can't modify its value
	if arr[0] != 0 {
		t.Errorf("modifyArray: expected %d, got %d", 0, arr[0])
	}
}

func TestModifySlice(t *testing.T) {
	t.Error("Oops 2")
	s := []int{0}
	// slice has a pointer to array inside and modifing a slice modifies
	// underline array
	modifySlice(s, -1)
	if s[0] != -1 {
		t.Errorf("modifySlice: expect %d, got %d", -1, s[0])
	}
}

func TestSliceAppend(t *testing.T) {

}
