package main

import (
	"algo/algo"
	"fmt"
	"sync"
)


func main() {
  var once sync.Once
  once.Do(func() {
    fmt.Println("Initializing")
  })

  var mutex sync.Mutex
  var wg sync.WaitGroup
  // example with negative and positive numbers from cmu link
  arr := []int{8, 14, -8, -9, 5, -9, -3, 0, 17, 19}
  fmt.Println("Unsorted Arr: ", arr)
  wg.Add(1)
  sorted := algo.InsertSortConcurrent(arr, &mutex, &wg)
  wg.Wait()

  fmt.Println("Sorted Arr: ", sorted)

}


