package main

import (
	"flag"
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

func desigilize(sigil *string, bufReader *bufio.Reader) {
	sigilPointer := sigil
	for strings.Count(*sigilPointer, "") > 1 {
		fmt.Println("Choose new character to remove: ")
		char, _, _ := bufReader.ReadRune()
		charAsString := string(char)
		newString := desigilizeOne(sigilPointer, &charAsString)
		fmt.Println("Left to sigilize: " + newString)
		sigilPointer = &newString
	}
}

func desigilizeOne(sigil *string, char *string) string {
	return strings.Replace(*sigil, *char, "", 1)
}

func main() {
    // Assist or not?
	assist := flag.Bool("assist", false, "Assist in desigilization")

	flag.Parse()

	if (*assist) {
		fmt.Println("Chose to assist with sigilization.")
	} else {
		fmt.Println("Chose not to assist with sigilization")
	}

	fmt.Println("Enter phrase to be sigilized: ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	sigilized := sigilize(&text)
	fmt.Println("Sigilized string: " + sigilized)

	if (*assist) {
		desigilize(&sigilized, reader)
	}
}