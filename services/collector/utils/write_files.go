package utils

import (
	"os"
)

// WriteResponseToFile takes a response.Body and a file name
// Writes a json file to the root directory with the name specified
func WriteJSONToFile(respBody []byte, name string) error {
	fileName := "./" + name + ".json"
	// Write response body to file
	err := os.WriteFile(fileName, respBody, 0644)
	if err != nil {
		return err
	}
	return nil
}
