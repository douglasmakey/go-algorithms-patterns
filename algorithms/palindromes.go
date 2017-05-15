package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"github.com/douglasmakey/golang-algorithms-/utils"
	"time"
	"math"
)

func isPalindrome(s string) bool {
	defer utils.TimeTrack(time.Now(), "isPalindrome")

	// Remove space of sentence
	s = fmt.Sprint(strings.Join((strings.Split(s, " ")), ""))

	// Init len and mid
	l := len(s)
	m := int(math.Floor(float64(l / 2)))
	for i := 0; i < m; i++ {
		if s[i] != s[l-1-i] {
			return false
		}
	}

	return true
}


func main() {
	// Read
	t := bufio.NewScanner(os.Stdin)
	fmt.Println("Ingres the sentence")
	t.Scan()
	s := t.Text()

	// Evaluate
	if isPalindrome(s) {
		fmt.Println("Sentence is Palindrome")
	} else {
		fmt.Println("Sentence is not palindrome")
	}
}
