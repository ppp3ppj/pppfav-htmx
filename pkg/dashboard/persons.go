package dashboard

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ppp3ppj/pppfav-htmx/template"
	views_dashboards_persons "github.com/ppp3ppj/pppfav-htmx/views/pages/dashboards/persons"
	views_variables "github.com/ppp3ppj/pppfav-htmx/views/variables"
)


func (fe *DashboardFrontend) Articles(c echo.Context) error {

    opts := views_variables.DashboardOpts{ }
    personVm := views_dashboards_persons.PersonsVm{
        Opts: opts,
    }
    dashboard := views_dashboards_persons.PersonsContent(personVm)

    return template.AssertRender(c, http.StatusOK, dashboard)
}
