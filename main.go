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

		context, err := prompter(contextPrompt(text))
		fmt.Println(context)
		text, _ = modifyPrompt(text, context)
		response, err := prompter(text)
		if err != nil {
			fmt.Println("There was an error with the response")
			return
		}
		parseAndRunResponse(response)
		fmt.Println("Commands Executed Successfully")
	}
}
