package person_validate

import (
	"regexp"
	"strconv"
	"strings"
	"time"
)

// validateDescription validates a description string and returns a boolean and a message.
func ValidateDescription(description string) (bool, string) {
	trimmedDescription := strings.TrimSpace(description)

	if len(trimmedDescription) == 0 {
		return false, "The description cannot be empty."
	}

	if len(trimmedDescription) < 10 {
		return false, "The description is too short. It must be at least 10 characters long."
	}

	if len(trimmedDescription) > 500 {
		return false, "The description is too long. It should be no more than 500 characters."
	}

	// Optional: Add a regex check for allowed characters.
	// Example: only allow letters, digits, spaces, punctuation.
	allowedCharacters := regexp.MustCompile(`^[a-zA-Z0-9\s.,;!?'"()-]+$`)
	if !allowedCharacters.MatchString(trimmedDescription) {
		return false, "The description contains invalid characters."
	}

	return true, "The description is valid."
}

// Enhanced regex for international names and special characters.
func ValidateName(name string) (bool, string) {
	trimmedName := strings.TrimSpace(name)

	if len(trimmedName) == 0 {
		return false, "The name cannot be empty."
	}

	// Allow alphabetic characters, spaces, hyphens, and apostrophes.
	matched, err := regexp.MatchString(`^[\p{L}\s'-]+$`, trimmedName)
	if err != nil {
		return false, "Error in validating the name."
	}
	if !matched {
		return false, "The name can only contain alphabetic characters, spaces, hyphens, and apostrophes."
	}

	if len(trimmedName) > 50 {
		return false, "The name is too long. It should be no more than 50 characters."
	}

	return true, "The name is valid."
}

// Validate the age.
// if -0 is 0 will not error
func ValidateAge(age string) (bool, string) {
	trimmedAge := strings.TrimSpace(age)

	if len(trimmedAge) == 0 {
		return false, "The age cannot be empty."
	}

	// Check if the age is a valid number.
	if _, err := strconv.Atoi(trimmedAge); err != nil {
		return false, "The age must be a valid number."
	}

	// Convert age to an integer.
	intAge, err := strconv.Atoi(trimmedAge)
	if err != nil {
		return false, "Error in processing the age."
	}

	// Validate the age range.
	if intAge < 0 || intAge > 100 {
		return false, "The age must be between 0 and 100."
	}

	return true, "The age is valid."
}

// Validate the birth date.
func ValidateBirthDate(birthDate string) (bool, string) {
	trimmedDate := strings.TrimSpace(birthDate)

	if len(trimmedDate) == 0 {
		return false, "The birth date cannot be empty."
	}

	// Define the date format. Assuming format is YYYY-MM-DD.
	const layout = "2006-01-02"
	parsedDate, err := time.Parse(layout, trimmedDate)
	if err != nil {
		return false, "The birth date must be in the format YYYY-MM-DD."
	}

	// Check if the date is in the future.
	if parsedDate.After(time.Now()) {
		return false, "The birth date cannot be in the future."
	}

	// Optionally, check if the date is more than 150 years in the past.
	earliestAllowedDate := time.Now().AddDate(-100, 0, 0)
	if parsedDate.Before(earliestAllowedDate) {
		return false, "The birth date cannot be more than 100 years ago."
	}

	return true, "The birth date is valid."
}
