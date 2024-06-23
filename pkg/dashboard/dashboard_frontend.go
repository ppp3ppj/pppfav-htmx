package dashboard

import (
	"net/http"

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
}

func (fe *DashboardFrontend) Index(c echo.Context) error {
    bodyOpts := views_variables.BodyOpts{
        ExtraHeaders: nil,
        Component: nil,
    }

    index := views_pages.Index(bodyOpts)
    return template.AssertRender(c, http.StatusOK, index)
}
