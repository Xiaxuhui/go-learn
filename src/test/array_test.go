package try_test

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [3]int{1, 2, 3}
	arr2 := [...]int{3, 4, 5, 6}
	t.Log(arr[1], arr[2]);
	for i := 0; i < len(arr2); i++ {
		t.Log(arr2[i]);
	}
	for _, item := range arr1 {
		t.Log(item);
	}
}

func TestArraySection(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4, 5}
	arr2 := [5]int{1, 2, 3, 4, 5}
	// arr3_sec := arr3[:]
	t.Log(arr2 == arr3);
}