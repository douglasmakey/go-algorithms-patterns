package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"github.com/douglasmakey/golang-algorithms-/utils"
	"time"
)

func isPalindrome(s string) (result bool) {
	defer utils.TimeTrack(time.Now(), "isPalindrome")

	// Remove space of sentence
	s = fmt.Sprint(strings.Join((strings.Split(s, " ")), ""))
	// Init and set position
	var i, l int
	l = (len(s) - 1)

	result = true
	for l/2 != 0 && result {
		if s[i] == s[l] {
			result = true
			i++
			l--
		} else {
			result = false
		}
	}

	return

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
