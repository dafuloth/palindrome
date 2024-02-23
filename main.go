// A solution written in Go sort a list of words into two lists.
//
// Input should be file containing a list of words, e.g. words.txt
//

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// A helper function to check whether a word is palindrome.
// Works by comparing characters. First with last, second with second-from-last, etc
func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	// Read from the file words.txt
	file, err := os.Open("words.txt")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var words []string

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error scanning file: %s", err)
	}

	// Declare two string slices that will be the two lists
	var normalWords, palindromes []string
	for _, word := range words {
		if isPalindrome(word) {
			palindromes = append(palindromes, word)
		}

		if !isPalindrome(word) {
			normalWords = append(normalWords, word)
		}
	}

	// Assuming words.txt contains: anna racecar normal
	fmt.Println("Palindromes: ", palindromes)  // Palindromes:  [anna racecar]
	fmt.Println("Normal Words: ", normalWords) // Normal Words:  [normal]
}
