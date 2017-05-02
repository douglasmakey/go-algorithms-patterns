package main

import (
	"fmt"
	"math/rand"
	"time"
	"log"
)

func main()  {
	fmt.Println("bubbleSort Algorithms")

	arr1 := createArraySeq(15)
	fmt.Println("Unsorted arraySeq: ", arr1 )
	bubbleSort(arr1)
	fmt.Println("bubleSort arraySeq: ", arr1)
	fmt.Println("---------------------------------")
	arr2 := createArrayRand(15)
	fmt.Println("Unsorted arrayRand: ", arr2)
	bubbleSort(arr2)
	fmt.Println("bubleSort arrayRand: ", arr2)

}

func createArraySeq(size int) []int {
	var arr []int
	for i := size; i > 0 ; i--  {
		arr = append(arr, i)
	}
	return arr
}

func createArrayRand(size int) []int {
	var arr []int
	for i := 0; i <= size ; i++  {
		var intRand int = rand.Intn(1000)
		arr = append(arr, intRand)
	}
	return arr
}

func swap(a []int, indexA, indexB int)  {
	var tmp int = a[indexA]
	a[indexA] = a[indexB]
	a[indexB] = tmp
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func bubbleSort(a []int){
	defer timeTrack(time.Now(), "bubbleSort")
	for i := 0; i < len(a) ; i++  {
		for j := 0; j < len(a) -1; j++  {
			if a[j] > a[j + 1] {
				swap(a, j, j+1)
			}
		}
	}
}
