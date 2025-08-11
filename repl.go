package main

import (
	"fmt"
	"bufio"
	"os"
	//"io"
	//"log"
	"strings"
	//"net/http"
	//"encoding/json"
)

type config struct {
	Next string
	Previous string
}

type cliCommand struct {
	name string
	description string
	callback func(*config) error
}

func startRepl(){
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	config := config{}
	//bufio.Scanner
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleaned := cleanInput(input)
		if len(cleaned) == 0{
			continue
		}
		firstWord := cleaned[0]
		command, exists := commands[firstWord]
		if exists {
			err := command.callback(&config)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Commands does not exist")
			continue
		}
		//fmt.Printf("Your command was: %s\n", firstWord)
	}
}

func cleanInput(text string) []string{
	ret := []string{}
	tmp := strings.ToLower(text) //lower and trim then split
	tmp = strings.TrimSpace(tmp)
	ret = strings.Fields(tmp)

	return ret
}
