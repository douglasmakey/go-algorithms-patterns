package utils

import (
	"math/rand"
	"time"
	"log"
)

func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func CreateArraySeq(size int) []int {
	var arr []int
	for i := size; i > 0; i-- {
		arr = append(arr, i)
	}
	return arr
}

func CreateArrayRand(size int) []int {
	var arr []int
	for i := 0; i <= size; i++ {
		var intRand int = rand.Intn(1000)
		arr = append(arr, intRand)
	}
	return arr
}

func Swap(a []int, indexA, indexB int) {
	var tmp int = a[indexA]
	a[indexA] = a[indexB]
	a[indexB] = tmp
}