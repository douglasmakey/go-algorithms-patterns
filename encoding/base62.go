package main

import (
	"bytes"
	"math"
	"fmt"
	"github.com/douglasmakey/golang-algorithms-/utils"
	"time"
)

// All characters
const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Convert number to base62
func Encode(n int) string {
	defer utils.TimeTrack(time.Now(), "Encode")

	if n == 0 {
		return string(alphabet[0])
	}

	chars := make([]byte, 0)

	length := len(alphabet)

	for n > 0 {
		r := n / length
		remainder := n % length
		chars = append(chars, alphabet[remainder])
		n = r
	}

	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}

	return string(chars)
}

// convert base62 token to int
func Decode(token string) int {
	defer utils.TimeTrack(time.Now(), "Decode")

	n := 0
	idx := 0.0
	chars := []byte(alphabet)

	charsLength := float64(len(chars))
	tokenLength := float64(len(token))

	for _, c := range []byte(token) {
		power := tokenLength - (idx + 1)
		ind := bytes.IndexByte(chars, c)
		n += ind * int(math.Pow(charsLength, power))
		idx++
	}

	return n
}

func main()  {
	encode := Encode(1600000000)
	fmt.Println("Encode: ", encode)

	decode := Decode(encode)
	fmt.Println("Decode: ", decode)


}
