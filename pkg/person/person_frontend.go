package person

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/ppp3ppj/pppfav-htmx/pkg/models"
	"github.com/ppp3ppj/pppfav-htmx/template"
	views_dashboards_persons_new_components "github.com/ppp3ppj/pppfav-htmx/views/pages/dashboards/persons/create/components"
)

type PersonFrontend struct {
    repo models.PersonRepository
}

func NewPersonFrontend(
    g *echo.Group,
    repo models.PersonRepository,
) {
    fe := &PersonFrontend{
        repo: repo,
    }
    _ = fe

    g.POST("/persons/validate/name", fe.ValidateName)
    g.POST("/persons/validate/age", fe.ValidateAge)
    g.POST("/persons/validate/birthdate", fe.ValidateBirthDate)
    g.POST("/persons/validate/description", fe.ValidateDescription)
}
func (fe *PersonFrontend) ValidateBirthDate(c echo.Context) error {
    birthDate := c.FormValue("birthDate")
    ok, errorMessage := validateBirthDate(birthDate)
    birthDateComponent := views_dashboards_persons_new_components.NewPersonVadidateVM{
        LabelId: "lblFieldBirthDate",
        InputId: "birthDate",
        InputType: "date",
        TitleName: "What is your birthdate",
        BasePath: "/persons/validate/birthdate",
        ContentName: birthDate,
    }

    if ok {
        birthDateComponent.StatusType = "Success"
        validateViewOK := views_dashboards_persons_new_components.NameFieldLabelValidation(birthDateComponent)
        return template.AssertRender(c, http.StatusOK, validateViewOK)
    } else {
        birthDateComponent.StatusType = "Error"
        birthDateComponent.ErrorMessage = errorMessage
        validateViewErr := views_dashboards_persons_new_components.NameFieldLabelValidation(birthDateComponent)
        return template.AssertRender(c, http.StatusOK, validateViewErr)
    }
}

func (fe *PersonFrontend) ValidateAge(c echo.Context) error {
    age := c.FormValue("age")
    ok, errorMessage := validateAge(age)
    ageComponent := views_dashboards_persons_new_components.NewPersonVadidateVM{
        LabelId: "lblFieldAge",
        InputId: "age",
        InputType: "text",
        TitleName: "What is your age",
        Placeholder: "Enter age",
        BasePath: "/persons/validate/age",
        ContentName: age,
    }

    if ok {
        ageComponent.StatusType = "Success"
        validateViewOK := views_dashboards_persons_new_components.NameFieldLabelValidation(ageComponent)
        return template.AssertRender(c, http.StatusOK, validateViewOK)
    } else {
        ageComponent.StatusType = "Error"
        ageComponent.ErrorMessage = errorMessage
        validateViewErr := views_dashboards_persons_new_components.NameFieldLabelValidation(ageComponent)
        return template.AssertRender(c, http.StatusOK, validateViewErr)
    }
}

func (fe *PersonFrontend) ValidateDescription(c echo.Context) error {
    description := c.FormValue("description")
    fmt.Println(description)
    isValid, message := validateDescription(description)
    descriptionComponent := views_dashboards_persons_new_components.NewPersonVadidateVM{
        LabelId: "lblFieldDescription",
        InputId: "description",
        InputType: "textarea",
        TitleName: "What is your description",
        Placeholder: "Enter description",
        BasePath: "/persons/validate/description",
        ContentName: description,
    }
    if !isValid {
        descriptionComponent.StatusType = "Error"
        descriptionComponent.ErrorMessage = message
        validateViewErr := views_dashboards_persons_new_components.TextAreaValidation(descriptionComponent)
        return template.AssertRender(c, http.StatusOK, validateViewErr)
    }

    validateViewOk := views_dashboards_persons_new_components.TextAreaValidation(descriptionComponent)
    descriptionComponent.StatusType = "Success"
    return template.AssertRender(c, http.StatusOK, validateViewOk)
}

func (fe *PersonFrontend) ValidateName(c echo.Context) error {
    name := c.FormValue("name")
    ok, strErr := validateName(name)
    newVM := views_dashboards_persons_new_components.NewPersonVadidateVM{
        LabelId: "lblFieldName",
        InputId: "name",
        InputType: "text",
        TitleName: "What is your name",
        Placeholder: "Enter name",
        BasePath: "/persons/validate/name",
        ContentName: name,
    }
    if ok {
        nameExists, err := fe.repo.CheckNameExists(newVM.ContentName)
        fmt.Println("Data is ")
        fmt.Printf("%t", nameExists)
        _ = err
        if nameExists {
            newVM.StatusType = "Error"
            newVM.ErrorMessage = "Name is exists"
            validateViewOK := views_dashboards_persons_new_components.NameFieldLabelValidation(newVM)
            return template.AssertRender(c, http.StatusOK, validateViewOK)
        } else {
            newVM.StatusType = "Success"
            validateViewErr := views_dashboards_persons_new_components.NameFieldLabelValidation(newVM)
            return template.AssertRender(c, http.StatusOK, validateViewErr)
        }
    } else {
        newVM.StatusType = "Error"
        newVM.ErrorMessage = strErr
        validateViewErr := views_dashboards_persons_new_components.NameFieldLabelValidation(newVM)
        return template.AssertRender(c, http.StatusOK, validateViewErr)
    }
}

// validateDescription validates a description string and returns a boolean and a message.
func validateDescription(description string) (bool, string) {
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
func validateName(name string) (bool, string) {
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
func validateAge(age string) (bool, string) {
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
func validateBirthDate(birthDate string) (bool, string) {
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
