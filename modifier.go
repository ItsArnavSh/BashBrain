package main

import "strings"

func modifyPrompt(prompt string) (string, error) {
	var modif strings.Builder

	modif.WriteString(prompt)
	modif.WriteString("Just give a series of commands for a ubuntu machine to do this. No other text please. No markdown or codeblocks plain json so that it can be used directly in code. If there are multiple commands, give the commands in JSON please. In each category, list  the command and whether its safe or not. Example ls touch etc are safe but rm dd are unsafe")

	return modif.String(), nil
}
