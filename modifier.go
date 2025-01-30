package main

import "strings"

func modifyPrompt(prompt string) (string, error) {
	var modif strings.Builder

	modif.WriteString(prompt)
	modif.WriteString("Just give a series of commands for a ubuntu machine to do this. No other text please. No markdown or indicators, just the plain commands. If there are multiple commands, give a '~~' seperated list of commands please")

	return modif.String(), nil
}
