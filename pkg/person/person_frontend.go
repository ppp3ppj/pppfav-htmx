package person

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"

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
    //g.POST("/persons/validate/birthdate", )
}

func (fe *PersonFrontend) ValidateAge(c echo.Context) error {
    age := c.FormValue("age")
    ok, errorMessage := validateAge(age)
    ageComponent := views_dashboards_persons_new_components.NewPersonVadidateVM{
        LabelId: "lblFieldAge",
        InputId: "age",
        TitleName: "What is your age",
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

func (fe *PersonFrontend) ValidateName(c echo.Context) error {
    name := c.FormValue("name")
    ok, strErr := validateName(name)
    newVM := views_dashboards_persons_new_components.NewPersonVadidateVM{
        LabelId: "lblFieldName",
        InputId: "name",
        TitleName: "What is your name",
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
