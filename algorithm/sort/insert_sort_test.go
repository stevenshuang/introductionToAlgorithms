package sort

import (
	"sort"
	"testing"
)

func TestInsertSort(t *testing.T) {
	for _, testCase := range [][]int{
		{6, 3, 7, 9, 1, 0, 23},
		{1, 2, 3, 4, 5},
		{1000, 1324, 878, 111, 2, 8734, 113247389},
	} {
		InsertSort(testCase)
		if !sort.IntsAreSorted(testCase) {
			t.Fatalf("%v are not sorted", testCase)
		}
	}
}
