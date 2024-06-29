package person

import (
	"fmt"
	"net/http"
	"regexp"
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
}

func (fe *PersonFrontend) ValidateName(c echo.Context) error {
    fmt.Println("ppppValidate")
    name := c.FormValue("name")
    fmt.Println(name)
    ok, strErr := validateName(name)
    newVM := views_dashboards_persons_new_components.NewPersonVadidateVM{
        ElementId: "lblFieldName",
        TitleName: "What is your name",
        BasePath: "/persons/validate/name",
        ContentName: name,
    }
    if ok {
        newVM.StatusType = "Success"
        validateViewOK := views_dashboards_persons_new_components.NameFieldLabelValidation(newVM)
        return template.AssertRender(c, http.StatusOK, validateViewOK)
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
