package utils

import "testing"

func TestRandShuffle(t *testing.T) {
	list := []interface{}{1, 2, 3, 4, 5}

	for i := 0; i < 10; i++ {
		RandShuffle(list)
	}

	t.Log(list)
}
