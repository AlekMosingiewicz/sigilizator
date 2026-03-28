package main

import (
	"fmt"
	"bufio"
	"strings"
	"unicode"
	"os"
)


func sigilize(phrase *string) string {
	var builder strings.Builder
	seen := make(map[rune]bool)
	for _, char := range *phrase {
		lowercase := unicode.ToLower(char)
		if char == ' ' {
			continue
		}
		if seen[lowercase] {
			continue
		}
		seen[lowercase] = true
		builder.WriteRune(lowercase)
	}
	return builder.String()
}

func main() {
	fmt.Println("Enter phrase to be sigilized: ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	sigilized := sigilize(&text)
	fmt.Println("Sigilized string: " + sigilized)

}