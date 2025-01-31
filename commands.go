package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Command struct {
	Command string `json:"command"`
	Safe    bool   `json:"safe"`
}

func parseAndRunResponse(response string) {
	start := strings.Index(response, "[")
	end := strings.LastIndex(response, "]") + 1
	if start == -1 || end == 0 {
		fmt.Println("Invalid JSON format")
		os.Exit(1)
	}
	response = response[start:end]
	var commands []Command
	if err := json.Unmarshal([]byte(response), &commands); err != nil {
		fmt.Println("Error parsing JSON:", err)
		os.Exit(1)
	}
	reader := bufio.NewReader(os.Stdin)
	for _, cmd := range commands {
		if !cmd.Safe {
			fmt.Print("Are you sure you want to run: ", cmd.Command, " Y to Run: ")
			text, _ := reader.ReadString('\n')
			if strings.TrimSpace(text) != "Y" {
				fmt.Println("Aborting Execution")
				os.Exit(0)
			}
		}
		fmt.Println(cmd.Command)
		cmdStruct := exec.Command("bash", "-c", cmd.Command)
		out, err := cmdStruct.Output()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(out))
	}
}
