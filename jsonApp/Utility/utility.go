package utility

import (
	"fmt"
	"io"
	"jsonApp/model"
	"os"
	"time"
)

func ReadCompaniesFromFile(filePath string) (data []byte, error error) {
	// Open the JSON file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	// Read the contents of the file
	fileContents, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	return fileContents, err
}

func ParseTimestamps(createdAt, updatedAt string) (time.Time, time.Time, *model.MyError) {

	parseTime := func(timeStr, label string) (time.Time, error) {
		// Try parsing with RFC3339
		t, err := time.Parse(time.RFC3339, timeStr)
		if err == nil {
			return t, nil
		}

		// If RFC3339 parsing fails, try the custom format
		t, err = time.Parse("2006-01-02 15:04:05 -0700 UTC", timeStr)
		if err != nil {
			return time.Time{}, fmt.Errorf("failed to parse time for %s: %v", label, err)
		}
		return t, nil
	}

	createdAtParsed, err := parseTime(createdAt, "Created At")
	if err != nil {
		str := "failed to parse time for CreatedAt" + createdAt
		err := &model.MyError{Message: str}
		return time.Time{}, time.Time{}, err
	}

	updatedAtParsed, err := parseTime(updatedAt, "Updated At")
	if err != nil {
		str := "failed to parse time for UpdatedAt" + updatedAt
		err := &model.MyError{Message: str}
		return time.Time{}, time.Time{}, err
	}

	return createdAtParsed, updatedAtParsed, nil
}

// // ValidateData checks if the data is valid for a given field and regex.
// func ValidateData(field model.Field) *model.MyError {
// 	if strings.TrimSpace(field.Value) == "" {
// 		str := field.Name + " is empty."
// 		err := &model.MyError{Message: str}
// 		return err
// 	}

// 	if field.Regex != nil && !field.Regex.MatchString(field.Value) {
// 		str := field.Name + " is not a valid field."
// 		err := &model.MyError{Message: str}
// 		return err
// 	}

// 	return nil
// }

// // ValidateEntity validates fields for different entities (e.g., Company, Admin, HR, Employee)
// func ValidateEntity(fields []model.Field) *model.MyError {
// 	for _, field := range fields {
// 		if err := ValidateData(field); err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }
