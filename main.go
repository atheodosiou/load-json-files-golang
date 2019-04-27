package main

import (
	"encoding/json"
	"fmt"
	"os"
)

//Struct that maches config.json file
type Config struct {
	Application string `json:"application"`
	Database    struct {
		Host string `json:"host"`
		Port string `json:"port"`
	} `json:"database"`
}

//Reads the json file
func LoadConfiguration(filename string) (Config, error) {
	var config Config
	configFile, err := os.Open(filename)
	//Closes the file
	defer configFile.Close()

	if err != nil {
		return config, err
	}

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	return config, err
}

func main() {
	fmt.Println("Starting the application...")
	config, _ := LoadConfiguration("config.json")

	fmt.Println("Application name: " + config.Application + "\nDatabase: \nHost: " + config.Database.Host + "\nPort: " + config.Database.Port)
}
