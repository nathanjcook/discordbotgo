package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

// Struct for commands
type Command struct {
	Command string `json:"command"`
	Info    string `json:"info"`
	Usage   string `json:"usage"`
}

// Function with type gin context
func GetHelp(c *gin.Context) {
	helpPath := "../../microservice_templates/Go/controllers/help.json"

	help, err := getHelp(helpPath)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	// Return as Raw JSON
	c.JSON(200, json.RawMessage(help))
}

// Function to read JSON file and return as a struct
func readJSONFile(filePath string) ([]Command, error) {
	var commands []Command

	// Open the JSON file
	jsonFile, err := os.Open(filePath)
	if err != nil {
		return commands, fmt.Errorf("failed to open file: %w", err)
	}
	defer jsonFile.Close()

	// Read the file into slice
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return commands, fmt.Errorf("failed to read file: %w", err)
	}

	// Unmarshal the JSON data
	if err := json.Unmarshal(byteValue, &commands); err != nil {
		return commands, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return commands, nil
}

// Return data as a JSON string
func getHelp(filePath string) (string, error) {
	// Read the JSON file
	commands, _ := readJSONFile(filePath)

	// Marshal the data with indents
	jsonData, _ := json.MarshalIndent(commands, "", "")

	return string(jsonData), nil
}
