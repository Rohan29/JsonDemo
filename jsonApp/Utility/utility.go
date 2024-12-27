package utility

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func ParseTime(timeStr string) (time.Time, error) {
	// Try parsing with RFC3339 first
	t, err := time.Parse(time.RFC3339, timeStr)
	if err == nil {
		return t, nil
	}

	// If RFC3339 parsing fails, try with custom format "2006-01-02 15:04:05 -0700 UTC"
	t, err = time.Parse("2006-01-02 15:04:05 -0700 UTC", timeStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse time: %v", err)
	}
	return t, nil
}

func ReadCompaniesFromFile(filePath string) (data []byte, error error) {
	// Open the JSON file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	// Read the contents of the file
	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	return fileContents, err
}
