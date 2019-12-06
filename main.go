package main

import (
	"fmt"
	"math"
)

func Less (word1, word2 string, alphaDict map[byte] int) bool {

	minLength := int(math.Min(float64(len(word1)), float64(len(word2))))
	i := 0
	for ; i < minLength; i++ {
		char1Order, _ := alphaDict[word1[i]]
		char2Order, _ := alphaDict[word2[i]]

		fmt.Printf("char1Order %d\n", char1Order)
		fmt.Printf("char2Order %d\n", char2Order)

		if char1Order < char2Order {
			return false
		}
	}

	if len(word1) > i {
		return false
	}

	return true

}

func isAlienSorted(words []string, order string) bool {
	alphaDict := make(map[byte] int)
	for i := 0; i < len(order); i++ {
		alphaDict[order[i]] = i
	}

	for i := 1; i < len(words); i++ {
		if !Less(words[i-1], words[i], alphaDict) {
			return false
		}
	}
	return true
}

func main() {
 fmt.Printf("Words %v are sorted %v", []string{"hello","leetcode"}, isAlienSorted([]string{"hello","leetcode"}, "hlabcdefgijkmnopqrstuvwxyz"))
}
