package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func contextPrompt(prompt string) string {
	var modif strings.Builder
	modif.WriteString("Do not tell what commands to write for this instruction. just respond with a json if you would need additional context like pwd or find to do this. The JSON should be just a list of commands whose outputs you would need to execute this. Like ls pwd all those. No extra text please or explaination please")
	modif.WriteString(prompt)
	return modif.String()
}
func modifyPrompt(prompt string, context string) (string, error) {
	var modif strings.Builder
	start := strings.Index(context, "[")
	end := strings.LastIndex(context, "]") + 1
	if start == -1 || end == 0 {
		fmt.Println("Invalid JSON form Context")
		os.Exit(1)
	}
	context = context[start:end]
	var commands []string
	if err := json.Unmarshal([]byte(context), &commands); err != nil {
		fmt.Println("Error parsing JSON:", err)
		os.Exit(1)
	}
	var extraData strings.Builder
	for _, cmd := range commands {
		extraData.WriteString(cmd)
		cmdStruct := exec.Command("bash", "-c", cmd)
		out, err := cmdStruct.Output()
		if err != nil {
			fmt.Println(err)
		}
		extraData.WriteString((string(out)))
	}
	modif.WriteString(prompt)
	modif.WriteString("Just give a series of commands for a ubuntu machine to do this. No other text please. No markdown or codeblocks plain json so that it can be used directly in code. If there are multiple commands, give the commands in JSON please. In each category, list  the command and whether its safe or not. Example ls touch etc are safe but rm dd are unsafe")
	modif.WriteString(extraData.String())
	return modif.String(), nil
}
