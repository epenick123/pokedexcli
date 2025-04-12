package main

import (
	"fmt"
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
	fmt.Printf("Hello, World!")
}
