package person

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ppp3ppj/pppfav-htmx/pkg/models"
	"github.com/ppp3ppj/pppfav-htmx/template"
	"github.com/ppp3ppj/pppfav-htmx/utils/validate/person_validate"
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
    ok, errorMessage := person_validate.ValidateBirthDate(birthDate)
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
    ok, errorMessage := person_validate.ValidateAge(age)
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
    isValid, message := person_validate.ValidateDescription(description)
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
    ok, strErr := person_validate.ValidateName(name)
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
