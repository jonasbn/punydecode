package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"golang.org/x/net/idna"
)

func main() {
	argsWithoutProg := os.Args[1:]

	var punycodeString string

	if len(argsWithoutProg) <= 0 {
		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			punycodeString = scanner.Text()
		}

		if err := scanner.Err(); err != nil {
			log.Println(err)
			os.Exit(1)
		}

	} else {
		punycodeString = os.Args[1] // we only take a single parameter, the string to decode
	}

	decodedString, err := idna.ToUnicode(punycodeString)

	if err == nil {
		fmt.Printf("%s\n", decodedString)
	}
}
