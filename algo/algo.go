package algo 

import (
  "sync"
)

// Algorithm taken from 
// https://www.cs.cmu.edu/~scandal/nesl/alg-sequence.html#insertsort
// function insertion_sort(a) =
// if #a < 2 then a
// else let
//     head = a[0];
//     rest = insertion_sort(drop(a,1));
//     i = count({x < head : x in rest});
// in take(rest,i)++[head]++drop(rest,i)
//
// InsertSortConcurrent runs in O(n) time complexity (said on the website)
func InsertSortConcurrent(arr []int, mutex *sync.Mutex, wg *sync.WaitGroup) []int {
  defer func() {
		if wg != nil {
			wg.Done()
		}
	}()
  // defer wg.Done()

  if len(arr) < 2 {
    return arr
  }

  head := arr[0]
  sorted := make([]int, len(arr))
  var innerWg sync.WaitGroup
  innerWg.Add(1)
  go func() {
    sorted = InsertSortConcurrent(arr[1:], mutex, &innerWg)
  }()
  innerWg.Wait()

  insertIndex := 0
  mutex.Lock()
  for _, v := range sorted {
    if v < head {
      insertIndex++
    } else {
      break
    }
  }
  mutex.Unlock()
  mutex.Lock()

  result := append(sorted[:insertIndex], append([]int{head}, sorted[insertIndex:]...)...)
  mutex.Unlock()

  return result
}
