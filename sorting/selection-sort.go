package main

import (
	"fmt"
	"github.com/douglasmakey/golang-algorithms-/utils"
	"time"
)

func selectionSort(a []int) {
	defer utils.TimeTrack(time.Now(), "selectionSort")
	var min int
	for i := 0; i < len(a); i++ {
		min = i

		for j := i; j < len(a); j++ {
			if a[min] > a[j] {
				min = j
			}
		}
		if i != min {
			utils.Swap(a, i, min)
		}
	}
}

func main() {
	fmt.Println("selectionSort Algorithms")
	arr1 := utils.CreateArraySeq(40)
	fmt.Println("Unsorted arraySeq: ", arr1)
	selectionSort(arr1)
	fmt.Println("selectionSort arraySeq: ", arr1)

	fmt.Println("---------------------------------")

	arr2 := utils.CreateArrayRand(40)
	fmt.Println("Unsorted arrayRand: ", arr2)
	selectionSort(arr2)
	fmt.Println("selectionSort arrayRand: ", arr2)
}
