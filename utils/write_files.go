package utils

import "os"

func WriteResponseToFile(respBody []byte) error {
	// Write response body to file
	err := os.WriteFile("../response.json", respBody, 0644)
	if err != nil {
		return err
	}
	return nil
}
