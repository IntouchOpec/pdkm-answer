package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	// texts := []string{"ABCDEZXY", "abcdezxy", "AbCdEzXy", "A1b2C3d4/+*"}
	// for _, text := range texts {
	// 	fmt.Println(AnswerSix(text))
	// }
	fmt.Println(AnswerSeven("dddd"))
	fmt.Println(AnswerSeven("world"))
	fmt.Println(AnswerSeven("cycle"))
	fmt.Println(AnswerSeven("abba"))
	// result, _ := OptionalOne(6)
	// fmt.Println(result)
	// fmt.Println(AnswerSeven("abacdec"))
}

// result AnswerSix function
const InResult string = "Invalid Input"
const CapitalResult string = "All Capital Letter"
const SmallResult string = "All Small Letter"
const MixResult string = "Mix"

const Da byte = 97
const Dz byte = 122
const DA byte = 65
const DZ byte = 90

func AnswerSix(text string) string {
	var totalChar int
	var isCapital bool
	var isSmall bool
	var isMix bool
	totalChar = len(text)
	if totalChar >= 10000 {
		return InResult
	}
	for i := 0; i < totalChar; i++ {
		if isMix {
			isCapital = IsUpper(text[i]) // DA <= text[i] && DZ >= text[i]
			isSmall = IsLetter(text[i])  // Da <= text[i] && Dz >= text[i]
			if !isSmall && !isCapital {
				return InResult
			}
		} else if isSmall {
			isSmall = IsLetter(text[i]) // Da <= text[i] && Dz >= text[i]
			isMix = IsUpper(text[i])    // DA <= text[i] && DZ >= text[i]
			if !isMix && !isSmall {
				return InResult
			}
		} else if isCapital {
			isCapital = IsUpper(text[i]) // DA <= text[i] && DZ >= text[i]
			isMix = IsLetter(text[i])    // Da <= text[i] && Dz >= text[i]
			if !isMix && !isCapital {
				return InResult
			}
		} else {
			isCapital = IsUpper(text[i]) // DA <= text[i] && DZ >= text[i]
			isSmall = IsLetter(text[i])  // Da <= text[i] && Dz >= text[i]
			if !isCapital && !isSmall {
				return InResult
			}
		}
	}
	if isMix {
		return MixResult
	} else if isSmall {
		return SmallResult
	} else if isCapital {
		return CapitalResult
	}
	return InResult
}

func IsUpper(as byte) bool {
	return DA <= as && DZ >= as
}

func IsLetter(as byte) bool {
	return Da <= as && Dz >= as
}

func AnswerSeven(text string) []string {
	var result []string
	var m string
	var isContinue bool
	var isAppend bool
	var chrmini [27]int
	for i := 0; i < len(text); i++ {
		var distCount int = 0
		for x := i; x < len(text); x++ {
			if chrmini[text[x]-Da] == 0 {
				m += string(text[x])
				distCount++
			} else {
				if text[x] == text[i] && x != i {
					isAppend = true
					break
				} else if len(text)-1 == i {
					isAppend = true
				}
			}
			chrmini[text[x]-Da]++
			if distCount == len(text) {
				isContinue = true
			}
		}
		if isContinue {
			result = append(result, m)
			break
		}
		if isAppend {
			if "a" == m {
				result = append(result, fmt.Sprintf("%sa", string(text[i])))
				break
			} else {
				result = append(result, m)
				m = string(text[i])
				isAppend = false
			}
		}
	}
	return result
}

func OptionalOne(n int) (int, error) {
	if 0 <= n && n >= int(math.Pow(10, 9)) {
		return 0, errors.New("")
	}
	return n + (n / 3), nil
}
