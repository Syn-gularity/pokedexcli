package main

import (
	"fmt"
	"os"
	"github.com/Syn-gularity/pokedexcli/jsonPokemonAPI"
)

func commandExit(config *config) error{
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
	//return fmt.Errorf("something happened during exit")
}

func commandHelp(config *config) error{
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, cmd := range getCommands(){
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap(config *config) error{
	url := "https://pokeapi.co/api/v2/location-area/"
	if config.Next != ""{
		url = config.Next
	}

	mapData, err := jsonPokemonAPI.LocationAreaList(url)
	if err != nil{
		return err
	}

	for i:=0;i<20;i++{
		fmt.Println(mapData.Results[i].Name)
	}
	//fmt.Println(config1.Next)
	config.Next = mapData.Next
	config.Previous = mapData.Previous
	return nil
}

func commandMapb(config *config) error{
	url := "https://pokeapi.co/api/v2/location-area/"
	if config.Previous != ""{
		url = config.Previous
	}

	mapData, err := jsonPokemonAPI.LocationAreaList(url)
	if err != nil{
		return err
	}

	for i:=0;i<20;i++{
		fmt.Println(mapData.Results[i].Name)
	}
	//fmt.Println(config1.Next)
	config.Next = mapData.Next
	config.Previous = mapData.Previous
	return nil
}

func getCommands() map[string]cliCommand{
	commands := map[string]cliCommand{
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help": {
        name:        "help",
        description: "Displays a help message",
        callback:    commandHelp,
    	},
		"map": {
        name:        "map",
        description: "Displays 20 locations",
        callback:    commandMap,
    	},
		"mapb": {
        name:        "mapb",
        description: "Displays the previous 20 locations",
        callback:    commandMapb,
    	},
	}
	return commands
}