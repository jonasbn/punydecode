package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"golang.org/x/net/idna"
)

// Main function, wrapping the realMain function for unit-test capability
func main() {
	os.Exit(realMain())
}

// realMain function, wrapped by the Main function, processes inpput STDIN and CLI and returns an integer indicating success/failure
func realMain() int {
	argsWithoutProg := os.Args[1:]

	var punycodeString string

	if len(argsWithoutProg) <= 0 {
		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			punycodeString = scanner.Text()
		}

		if err := scanner.Err(); err != nil {
			log.Println(err)
			return 2
		}

	} else {
		punycodeString = os.Args[1] // we only take a single parameter, the string to decode
	}

	if punycodeString != "" {

		decodedString, err := idna.ToUnicode(punycodeString)

		if err == nil {
			fmt.Printf("%s\n", decodedString)
			return 0
		}
	}

	return 1
}
