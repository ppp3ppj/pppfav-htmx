package dashboard

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/ppp3ppj/pppfav-htmx/pkg/models"
	"github.com/ppp3ppj/pppfav-htmx/template"
	views_dashboards_persons "github.com/ppp3ppj/pppfav-htmx/views/pages/dashboards/persons"
	views_dashboards_persons_new "github.com/ppp3ppj/pppfav-htmx/views/pages/dashboards/persons/create"
	views_variables "github.com/ppp3ppj/pppfav-htmx/views/variables"
)

func (fe *DashboardFrontend) PersonsNew(c echo.Context) error {
    opts := views_variables.DashboardOpts{
        Nav: nav(0),
    }

    vm := views_dashboards_persons_new.NewPersonVM{
        Opts: opts,
    }

    personsNew := views_dashboards_persons_new.New(vm)

    return template.AssertRender(c, http.StatusOK, personsNew)
}


func (fe *DashboardFrontend) Articles(c echo.Context) error {

    opts := views_variables.DashboardOpts{
        Nav: nav(0),
    }
    mockPerson := mockData()
    personVm := views_dashboards_persons.PersonsVm{
        Opts: opts,
        Persons: mockPerson,
        CreatePath: "/dashboard/persons/new",

    }
    dashboard := views_dashboards_persons.PersonsContent(personVm)

    return template.AssertRender(c, http.StatusOK, dashboard)
}

func mockData() []models.Person {
    return []models.Person{
        {
            ID:          "3414",
            Name:        "Alice Johnson",
            Age:         28,
            BirthDate:   time.Date(1996, time.April, 12, 0, 0, 0, 0, time.UTC),
            ImageURL:    "https://example.com/images/alice.jpg",
            Description: "Alice is a software engineer with a passion for open-source projects.",
        },
        {
            ID:          "24234",
            Name:        "Bob Smith",
            Age:         35,
            BirthDate:   time.Date(1988, time.February, 20, 0, 0, 0, 0, time.UTC),
            ImageURL:    "https://example.com/images/bob.jpg",
            Description: "Bob is a graphic designer who loves to work on creative projects.",
        },
        {
            ID:          "ouoeu",
            Name:        "Carol Davis",
            Age:         42,
            BirthDate:   time.Date(1981, time.July, 15, 0, 0, 0, 0, time.UTC),
            ImageURL:    "https://example.com/images/carol.jpg",
            Description: "Carol is a project manager with over 20 years of experience in the industry.",
        },
        {
            ID:          "oueu",
            Name:        "David Brown",
            Age:         30,
            BirthDate:   time.Date(1993, time.December, 5, 0, 0, 0, 0, time.UTC),
            ImageURL:    "https://example.com/images/david.jpg",
            Description: "David is a data scientist who specializes in machine learning and AI.",
        },
        {
            ID:          "oeuou",
            Name:        "Eve Wilson",
            Age:         25,
            BirthDate:   time.Date(1999, time.May, 30, 0, 0, 0, 0, time.UTC),
            ImageURL:    "https://example.com/images/eve.jpg",
            Description: "Eve is a marketing specialist with a talent for digital campaigns.",
        },
    }
}

func (fe *DashboardFrontend) PersonsPush(c echo.Context) error {
    var req PersonCreateRequest
    if err := c.Bind(&req); err != nil {
    }

    opts := views_variables.DashboardOpts{
        Nav: nav(0),
    }

    vm := views_dashboards_persons_new.NewPersonVM{
        Opts: opts,
    }

    personsNew := views_dashboards_persons_new.New(vm)

    return template.AssertRender(c, http.StatusOK, personsNew)
}
