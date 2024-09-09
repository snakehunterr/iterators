package iterators

import (
	"math/rand"
	"testing"
)

func Test_generator_puts_values(t *testing.T) {
	nums := make([]int, 10)

	for i := 0; i < len(nums); i++ {
		nums[i] = i
	}

	gen := Generator[int]{}.New(len(nums))

	for _, num := range nums {
		gen.Put(num)
	}

	for i := 0; i < len(nums); i++ {
		res, ok := gen.Next()

		if res != nums[i] {
			t.Errorf("res != num (%d != %d)", res, nums[i])
		}
		if ok != true {
			t.Error("generator closed unexpectedly!")
		}
	}
}

func Test_generator_generates_values(t *testing.T) {
	rnds := make([]int, 1000)

	for i := 0; i < len(rnds); i++ {
		rnds[i] = rand.Intn(10_000)
	}

	gen := Generator[int]{}.New(1)

	for _, num := range rnds {
		gen.Put(num)

		result, ok := gen.Next()

		if result != num {
			t.Errorf("result != num (%d != %d)", result, num)
		}
		if ok != true {
			t.Fatal("generator is closed unexpectedly!")
		}
	}
}

func Test_generator_closes(t *testing.T) {
	input := []int{1, 2, 3}

	gen := Generator[int]{}.New(1)

	for _, num := range input {
		gen.Put(num)

		_, ok := gen.Next()

		if ok != true {
			t.Fatal("generator is closed unexpectedly!")
		}
	}

	gen.Close()

	_, ok := gen.Next()

	if ok != false {
		t.Error("generator not closed properly")
	}

	if !gen.IsClosed {
		t.Error("generator IsClosed = false while generator is closed")
	}
}
