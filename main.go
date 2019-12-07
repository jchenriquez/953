package main

import (
	"fmt"
	"math"
)

func Less (word1, word2 string, alphaDict map[byte] int) bool {

	minLength := int(math.Min(float64(len(word1)), float64(len(word2))))
	i := 0
	word1CurrentOrder := 0
  	word2CurrentOrder := 0

  for ; i < minLength; i++ {
    char1Value,_ := alphaDict[word1[i]]
    char2Value,_ := alphaDict[word2[i]]


    if char1Value < char2Value && word1CurrentOrder == word2CurrentOrder {
      return true
    }

    if char1Value > char2Value && word1CurrentOrder == word2CurrentOrder {
    	return false
	}

  	word1CurrentOrder+=char1Value
  	word2CurrentOrder+=char2Value
  }

	if word1CurrentOrder == word2CurrentOrder && len(word1) > i {
    	return false
  	}

  return true
}

func isAlienSorted(words []string, order string) bool {
	alphaDict := make(map[byte] int, len(order))
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
 fmt.Printf("Words %v are sorted %v", []string{"word", "world", "row"}, isAlienSorted([]string{"word", "world", "row"}, "hlabcdefgijkmnopqrstuvxyz"))
}
