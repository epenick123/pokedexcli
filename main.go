package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

// Global commands map
var commands map[string]cliCommand

func cleanInput(text string) []string {
	cleanedText := []string{}
	processedText := strings.TrimSpace(text)
	processedText = strings.ToLower(processedText)
	cleanedText = strings.Fields(processedText)
	return cleanedText
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil // This line will never actually execute because os.Exit terminates the program
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	// Iterate over the commands map
	for _, cmd := range commands {
		// Print each command's name and description
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message with available commands",
			callback:    commandHelp,
		},
	}
	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleanedInput := cleanInput(input)

		// Check if input is empty
		if len(cleanedInput) == 0 {
			fmt.Println("No command entered")
			continue
		}

		firstWord := cleanedInput[0]
		// Use the whole phrase as the key to look up commands
		command, exists := commands[firstWord]
		if exists {
			err := command.callback()
			if err != nil {
				fmt.Println("Error:", err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}
