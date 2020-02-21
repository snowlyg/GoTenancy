package excel

import "testing"

func TestToAxis(t *testing.T) {
	results := map[string][]int{
		"A1":  []int{0, 1},
		"B5":  []int{1, 5},
		"Z1":  []int{25, 1},
		"Z10": []int{25, 10},
		"AA1": []int{26, 1},
		"AZ9": []int{51, 9},
		"BA8": []int{52, 8},
	}

	for key, value := range results {
		if axis := toAxis(value[0], value[1]); axis != key {
			t.Errorf("%v's axis should be %v, but got %v", value, key, axis)
		}
	}
}
