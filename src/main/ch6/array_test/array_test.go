package array_test

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	t.Log(arr[0], arr[1])

	arr1 := [4]int{1, 2, 3}
	t.Log(arr1[0], arr1[1])

	arr2 := [...]int{1, 2, 3, 4}
	arr2[1] = 0
	t.Log(arr2[0], arr2[1])
}

func TestArrayTravel(t *testing.T) {
	arr2 := [...]int{1, 2, 3, 4}
	for i := 0; i < len(arr2); i++ {
		t.Log(arr2[i])
	}
	// idx 下标 _ 是占位
	for _, e := range arr2 {
		t.Log(e)
	}
}

func TestArraySection(t *testing.T) {
	arr2 := [...]int{1, 2, 3, 4}
	arr2_sec := arr2[:3]
	t.Log(arr2_sec)
}
