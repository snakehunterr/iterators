package iterators

import "testing"

func Test_map_works_with_integers(t *testing.T) {
	nums := make([]int, 100)

	for i := 0; i < len(nums); i++ {
		nums[i] = i
	}

	gen := Map(nums, func(elem int) int { return elem * elem })

	for _, num := range nums {
		res, ok := gen.Next()
		doubled := num * num

		if doubled != res {
			t.Errorf("res != test (%d != %d)", res, doubled)
		}
		if ok != true {
			t.Error("generator closed unexpectedly")
		}
	}

	if !gen.IsClosed {
		t.Error("generator not closed after walking through all values")

		for !gen.IsClosed {
			res, _ := gen.Next()
			t.Error("unexpected:", res)
		}
	}
}
