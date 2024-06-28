package person

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/ppp3ppj/pppfav-htmx/pkg/models"
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
    return nil
}

