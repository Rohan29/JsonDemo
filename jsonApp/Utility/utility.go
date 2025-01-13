package utility

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"time"

	"github.com/go-playground/validator"
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

func ParseTimestamps(createdAt, updatedAt string) (time.Time, time.Time, error) {

	parseTime := func(timeStr, label string) (time.Time, error) {
		// Try parsing with RFC3339
		t, err := time.Parse(time.RFC3339, timeStr)
		if err == nil {
			return t, nil
		}

		// If RFC3339 parsing fails, try the custom format
		t, err = time.Parse("2006-01-02 15:04:05 -0700 UTC", timeStr)
		if err != nil {
			return time.Time{}, err
		}
		return t, nil
	}

	createdAtParsed, err := parseTime(createdAt, "Created At")
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	updatedAtParsed, err := parseTime(updatedAt, "Updated At")
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	return createdAtParsed, updatedAtParsed, nil
}

func regexValidator(fl validator.FieldLevel) bool {
	tag := fl.Param()
	match, _ := regexp.MatchString(tag, fl.Field().String())
	return match
}

// ValidateStruct validates any struct based on its tags
func ValidateStruct(s interface{}) error {
	validate := validator.New()
	validate.RegisterValidation("regex", regexValidator)

	if err := validate.Struct(s); err != nil {
		errorMessages := ""
		for _, err := range err.(validator.ValidationErrors) {
			errorMessages += fmt.Sprintf("Field '%s' failed validation with tag '%s'. ", err.StructField(), err.Tag())
		}
		return err
	}
	return nil
}
