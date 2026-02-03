package iteration

import "strings"

// const repeatCount = 5

func Repeat(character string, repeatCount int) string {
	return strings.Repeat(character, repeatCount)
	// var repeated strings.Builder
	// for i := 0; i < repeatCount; i++ {

	// 	// repeated.WriteString(character)
	// 	// repeated += character
	// }
	// return repeated.String()
}
