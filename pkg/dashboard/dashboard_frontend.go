package dashboard

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	public_pages "github.com/ppp3ppj/pppfav-htmx/public/pages/index"
	public_variables "github.com/ppp3ppj/pppfav-htmx/public/variables"
	"github.com/ppp3ppj/pppfav-htmx/template"
)

type DashboardFrontend struct {

}

func NewDashBoardFrontend(g *echo.Group) {
    fe := &DashboardFrontend{ }

    g.GET("", fe.Index)
    g.GET("/", fe.Index)
    g.GET("/dashboard", fe.Index)
}

func (fe *DashboardFrontend) Index(c echo.Context) error {
    bodyOpts := public_variables.BodyOpts{
        ExtraHeaders: nil,
        Component: nil,
    }

    fmt.Println("ppp")

    index := public_pages.Index(bodyOpts)
    fmt.Println("ppp - index")
    return template.AssertRenderLog(c, http.StatusOK, index)
    //return template.RenderEmpty(c)
}
