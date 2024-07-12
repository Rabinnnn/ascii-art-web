package asciiart

import (
	"fmt"
	"strings"
)

// Fuction takes a string and map of ascii patterns
// Checks if it is a printable ascii character
// If yes, the it generates the pattern in a continous string format(while also respecting "\n")
// Otherwise, it notifies the user the existence of a unprintable character in the input
func Art(s string, mYmap map[int][]string) (builder strings.Builder) {
	// var builder strings.Builder
	var str string
	// s = Backspace(s)
	// s = Tab(s)
	// s = strings.ReplaceAll(s, "\n", "\\n")
	input := strings.Split(s, "\n")
	// fmt.Println(input)
	// generating ascii-art for a string
	for j, word := range input {
		if word == "" && j != len(input)-1 {
			str += fmt.Sprintln()
			continue
		} else if word != "" {
			for i := 0; i < 8; i++ {
				for _, cha := range word {
					// check if character is a printable chararcter
					if ok := (cha >= ' ' && cha <= rune(126)); ok {
						str += mYmap[int(cha)][i]
					} // else if !ok {
					// 	fmt.Println("Unprintable character within input")
					// 	return ""
					// 	// os.Exit(1)
					// }
				}
				str += fmt.Sprintln()
			}
		}
	}
	builder.WriteString(str)
	return
}
