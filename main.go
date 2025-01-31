package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("How can I assist you?")
	for {
		text, _ := reader.ReadString('\n')
		text, _ = modifyPrompt(text)
		response, err := prompter(text)
		if err != nil {
			fmt.Println("There was an error with the response")
			return
		}
		parseAndRunResponse(response)
	}
}
