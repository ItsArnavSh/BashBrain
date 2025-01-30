package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	// cmdStruct := exec.Command("ls")
	// out, err := cmdStruct.Output()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(out))
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
		fmt.Print(response)
		commands := strings.Split(response, "~~")
		for _, command := range commands {
			cmdStruct := exec.Command("bash", "-c", command)
			out, err := cmdStruct.Output()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(string(out))
		}
		fmt.Print(response)
	}
}
