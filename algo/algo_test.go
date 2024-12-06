package algo_test

import (

  "algo/algo"
  "testing"
  "sync"
  "reflect"
)

// Unit Test for InsertionSortConcurrent
func TestInsertionSortConcurrent(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"Already sorted", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"Reverse order", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"With duplicates", []int{4, 2, 2, 8, 3}, []int{2, 2, 3, 4, 8}},
		{"Negative numbers", []int{-1, -3, -2, -4, 0}, []int{-4, -3, -2, -1, 0}},
		{"Mixed numbers", []int{0, -1, 3, -2, 1}, []int{-2, -1, 0, 1, 3}},
		{"Single element", []int{42}, []int{42}},
		{"Empty array", []int{}, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var mutex sync.Mutex
			var wg sync.WaitGroup
			wg.Add(1)

			result := algo.InsertSortConcurrent(tt.input, &mutex, &wg)
			wg.Wait()

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("insertionSortConcurrent(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}
