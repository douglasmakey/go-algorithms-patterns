package main

import (
	"fmt"
	"github.com/douglasmakey/golang-algorithms-/utils"
	"time"
)

func insertionSort(a []int) {
	defer utils.TimeTrack(time.Now(), "insertionSort")

	for i := 1; i < len(a); i++ {
		for j := i; j > 0 && a[j] < a[j - 1]; j-- {
			utils.Swap(a, j, j - 1)
		}
	}

}

func main() {
	fmt.Println("insertionSort Algorithms")

	arr1 := utils.CreateArraySeq(40)
	fmt.Println("Unsorted arraySeq: ", arr1)
	insertionSort(arr1)
	fmt.Println("insertionSort arraySeq: ", arr1)

	fmt.Println("---------------------------------")

	arr2 := utils.CreateArrayRand(40)
	fmt.Println("Unsorted arrayRand: ", arr2)
	insertionSort(arr2)
	fmt.Println("insertionSort arrayRand: ", arr2)
}

