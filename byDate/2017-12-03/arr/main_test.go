package main

import "testing"

func TestModifyArray(t *testing.T) {
	arr := [...]int{0}
	modifyArray(arr, -1)
	if arr[0] != 0 {
		t.Errorf("modifyArray: expected %d, got %d", 0, arr[0])
	}
}

func TestModifySlice(t *testing.T) {
	s := []int{0}
	modifySlice(s, -1)
	if s[0] != -1 {
		t.Errorf("modifySlice: expect %d, got %d", -1, s[0])
	}
}

func TestSliceAppend(t *testing.T) {

}
