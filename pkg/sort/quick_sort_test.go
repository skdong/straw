package sort

import (
	"testing"
)

func list_equal(t *testing.T, want []int, got []int) {
	if len(want) != len(got) {
		t.Errorf("sort was incorrect, got: %v , want: %v", got, want)
	}
	for i := 0; i < len(want); i++ {
		if want[i] != got[i] {
			t.Errorf("sort was incorrect, got: %v , want: %v", got, want)
		}
	}

}

func TestQuickSort(t *testing.T) {
	var want = []int{1, 2, 3, 4}
	got := QuickSort([]int{4, 3, 2, 1})
	list_equal(t, want, got)
	//if got != want {
	//t.Errorf("sort was incorrect, got: %v , want: %v", got, want)
	//}
}
