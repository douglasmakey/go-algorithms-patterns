package main

import (
	"fmt"
	"time"
	"github.com/douglasmakey/golang-algorithms-/utils"
)

func bubbleSort(a []int) {
	defer utils.TimeTrack(time.Now(), "bubbleSort")
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-1; j++ {
			if a[j] > a[j+1] {
				utils.Swap(a, j, j+1)
			}
		}
	}
}


func main() {
	fmt.Println("bubbleSort Algorithms")

	arr1 := utils.CreateArraySeq(40)
	fmt.Println("Unsorted arraySeq: ", arr1)
	bubbleSort(arr1)
	fmt.Println("bubleSort arraySeq: ", arr1)

	fmt.Println("---------------------------------")

	arr2 := utils.CreateArrayRand(40)
	fmt.Println("Unsorted arrayRand: ", arr2)
	bubbleSort(arr2)
	fmt.Println("bubleSort arrayRand: ", arr2)

}