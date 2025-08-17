package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"good-db/pkg/executor"
	"good-db/pkg/parser"
	"good-db/pkg/storage"
)

func main() {
	fmt.Println("Welcome to Good-DB!")
	fmt.Println("Type 'exit' to quit.")

	storage := storage.NewMemoryStorage()
	defer storage.Close()

	parser := parser.NewParser()
	executor := executor.NewExecutor(storage)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("good-db> ")
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}

		if strings.ToLower(input) == "exit" {
			fmt.Println("Goodbye!")
			break
		}

		query, err := parser.Parse(input)
		if err != nil {
			fmt.Printf("Parse error: %v\n", err)
			continue
		}

		result, err := executor.Execute(query)
		if err != nil {
			fmt.Printf("Execution error: %v\n", err)
			continue
		}

		fmt.Printf("Result: %v\n", result)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading input: %v\n", err)
	}
}