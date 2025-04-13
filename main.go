package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	cleanedText := []string{}
	processedText := strings.TrimSpace(text)
	processedText = strings.ToLower(processedText)
	cleanedText = strings.Fields(processedText)
	return cleanedText
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleanedInput := cleanInput(input)
		firstWord := cleanedInput[0]
		fmt.Printf("Your command was: %s\n", firstWord)
	}
}
