package main

import "strings"

func cleanInput(text string) []string {
	lowerCase := strings.ToLower(text)
	fields := strings.Fields(lowerCase)
	return fields 
}

