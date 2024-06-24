package dashboard

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/ppp3ppj/pppfav-htmx/template"
	views_pages "github.com/ppp3ppj/pppfav-htmx/views/pages/index"
	views_variables "github.com/ppp3ppj/pppfav-htmx/views/variables"
)

type DashboardFrontend struct {

}

func NewDashBoardFrontend(g *echo.Group) {
    fe := &DashboardFrontend{ }

    g.GET("", fe.Index)

    g.GET("/dashboard", func(c echo.Context) error {
        return c.Redirect(301, "/dashboard/persons")
    })
    g.GET("/dashboard/persons", fe.Articles)
    g.GET("/dashboard/persons/new", fe.PersonsNew)
    g.POST("/dashboard/persons/push", fe.PersonsPush)
}

func (fe *DashboardFrontend) Index(c echo.Context) error {
    bodyOpts := views_variables.BodyOpts{
        ExtraHeaders: nil,
        Component: nil,
    }

    index := views_pages.Index(bodyOpts)
    return template.AssertRender(c, http.StatusOK, index)
}

type PersonCreateRequest struct {
    Name        string    `json:"name" form:"name"`
    Age         int       `json:"age" form:"age"`
    BirthDate   time.Time `json:"birth_date" form:"birth_date"`   // Uses time.Time to represent date
    ImageURL    string    `json:"image_url" form:"image_url"`     // URL to the image of the person
    Description string    `json:"description" form:"description"` // Description or biography of the person
}
