package main

import "strings"

func cleanInput(text string) []string {
	trimmed := strings.TrimSpace(text)
	return strings.Split(trimmed, " ")
}
