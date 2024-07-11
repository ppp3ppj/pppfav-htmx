package dashboard

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/ppp3ppj/pppfav-htmx/pkg/models"
	"github.com/ppp3ppj/pppfav-htmx/template"
	"github.com/ppp3ppj/pppfav-htmx/utils/validate/person_validate"
	views_alert "github.com/ppp3ppj/pppfav-htmx/views/components/alert"
	views_person_card "github.com/ppp3ppj/pppfav-htmx/views/components/personCard"
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

func (fe *DashboardFrontend) IndexPersons(c echo.Context) error {
    mockPerson := mockData()
    personDisplay := views_person_card.PersonCards(mockPerson)
    return template.AssertRender(c, http.StatusOK, personDisplay)
}


func (fe *DashboardFrontend) PersonsImageMasonry(c echo.Context) error {
    opts := views_variables.DashboardOpts{
        Nav: nav(0),
    }
    mockPerson := mockData()
    personVm := views_dashboards_persons.PersonsImageVm{
        Opts: opts,
        Persons: mockPerson,
    }
    personsImageMasonryDashboard := views_dashboards_persons.PersonsImageViewContent(personVm)

    return template.AssertRender(c, http.StatusOK, personsImageMasonryDashboard)
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
            ImageURL:    "/uploads/6431467b554c450fbe29a9b9e30ce1b4.jpg",
            Description: "Alice is a software engineer with a passion for open-source projects.",
        },
        {
            ID:          "24234",
            Name:        "Bob Smith",
            Age:         35,
            BirthDate:   time.Date(1988, time.February, 20, 0, 0, 0, 0, time.UTC),
            ImageURL:    "/uploads/6431467b554c450fbe29a9b9e30ce1b4.jpg",
            Description: "Bob is a graphic designer who loves to work on creative projects.",
        },
        {
            ID:          "ouoeu",
            Name:        "Carol Davis",
            Age:         42,
            BirthDate:   time.Date(1981, time.July, 15, 0, 0, 0, 0, time.UTC),
            ImageURL:    "/uploads/6431467b554c450fbe29a9b9e30ce1b4.jpg",
            Description: "Carol is a project manager with over 20 years of experience in the industry.",
        },
        {
            ID:          "oueu",
            Name:        "David Brown",
            Age:         30,
            BirthDate:   time.Date(1993, time.December, 5, 0, 0, 0, 0, time.UTC),
            ImageURL:    "/uploads/6431467b554c450fbe29a9b9e30ce1b4.jpg",
            Description: "David is a data scientist who specializes in machine learning and AI.",
        },
        {
            ID:          "oeuou",
            Name:        "Eve Wilson",
            Age:         25,
            BirthDate:   time.Date(1999, time.May, 30, 0, 0, 0, 0, time.UTC),
            ImageURL:    "/uploads/6431467b554c450fbe29a9b9e30ce1b4.jpg",
            Description: "Eve is a marketing specialist with a talent for digital campaigns.",
        },
        {
            ID:          "oeuou",
            Name:        "Eve Wilson",
            Age:         25,
            BirthDate:   time.Date(1999, time.May, 30, 0, 0, 0, 0, time.UTC),
            ImageURL:    "/uploads/6431467b554c450fbe29a9b9e30ce1b4.jpg",
            Description: "Eve is a marketing specialist with a talent for digital campaigns.",
        },
        {
            ID:          "oeuou",
            Name:        "Eve Wilson",
            Age:         25,
            BirthDate:   time.Date(1999, time.May, 30, 0, 0, 0, 0, time.UTC),
            ImageURL:    "/uploads/6431467b554c450fbe29a9b9e30ce1b4.jpg",
            Description: "Eve is a marketing specialist with a talent for digital campaigns.",
        },
        {
            ID:          "oeuou",
            Name:        "Eve Wilson",
            Age:         25,
            BirthDate:   time.Date(1999, time.May, 30, 0, 0, 0, 0, time.UTC),
            ImageURL:    "/uploads/6431467b554c450fbe29a9b9e30ce1b4.jpg",
            Description: "Eve is a marketing specialist with a talent for digital campaigns.",
        },
        {
            ID:          "oeuou",
            Name:        "Eve Wilson",
            Age:         25,
            BirthDate:   time.Date(1999, time.May, 30, 0, 0, 0, 0, time.UTC),
            ImageURL:    "/uploads/6431467b554c450fbe29a9b9e30ce1b4.jpg",
            Description: "Eve is a marketing specialist with a talent for digital campaigns.",
        },
        {
            ID:          "oeuou",
            Name:        "Eve Wilson",
            Age:         25,
            BirthDate:   time.Date(1999, time.May, 30, 0, 0, 0, 0, time.UTC),
            ImageURL:    "/uploads/6431467b554c450fbe29a9b9e30ce1b4.jpg",
            Description: "Eve is a marketing specialist with a talent for digital campaigns.",
        },
        {
            ID:          "oeuou",
            Name:        "Eve Wilson",
            Age:         25,
            BirthDate:   time.Date(1999, time.May, 30, 0, 0, 0, 0, time.UTC),
            ImageURL:    "/uploads/6431467b554c450fbe29a9b9e30ce1b4.jpg",
            Description: "Eve is a marketing specialist with a talent for digital campaigns.",
        },
        {
            ID:          "oeuou",
            Name:        "Eve Wilson",
            Age:         25,
            BirthDate:   time.Date(1999, time.May, 30, 0, 0, 0, 0, time.UTC),
            ImageURL:    "/uploads/6431467b554c450fbe29a9b9e30ce1b4.jpg",
            Description: "Eve is a marketing specialist with a talent for digital campaigns.",
        },
        {
            ID:          "oeuou",
            Name:        "Eve Wilson",
            Age:         25,
            BirthDate:   time.Date(1999, time.May, 30, 0, 0, 0, 0, time.UTC),
            ImageURL:    "/uploads/6431467b554c450fbe29a9b9e30ce1b4.jpg",
            Description: "Eve is a marketing specialist with a talent for digital campaigns.",
        },
        {
            ID:          "oeuou",
            Name:        "Eve Wilson",
            Age:         25,
            BirthDate:   time.Date(1999, time.May, 30, 0, 0, 0, 0, time.UTC),
            ImageURL:    "/uploads/6431467b554c450fbe29a9b9e30ce1b4.jpg",
            Description: "Eve is a marketing specialist with a talent for digital campaigns.",
        },
        {
            ID:          "oeuou",
            Name:        "Eve Wilson",
            Age:         25,
            BirthDate:   time.Date(1999, time.May, 30, 0, 0, 0, 0, time.UTC),
            ImageURL:    "/uploads/6431467b554c450fbe29a9b9e30ce1b4.jpg",
            Description: "Eve is a marketing specialist with a talent for digital campaigns.",
        },
    }
}

func (fe *DashboardFrontend) PersonsPush(c echo.Context) error {

    var req PersonCreateRequest

    nameReq := c.FormValue("name")
    nameOk, nameErr := person_validate.ValidateName(nameReq)
    if nameOk {
        nameExists, err := fe.PersonRepo.CheckNameExists(nameReq)
        fmt.Println("Data is ")
        fmt.Printf("%t", nameExists)
        _ = err
        if nameExists {
            fmt.Println("Name Exists")
        } else {
            fmt.Println("OK")
        }
    } else {
            fmt.Println("Error not ok")
            fail := views_alert.AlertFailure("Failed to save person:", nameErr)
            return template.AssertRender(c, http.StatusOK, fail)
    }

    ageReq := c.FormValue("age")

    ageOk, ageErr := person_validate.ValidateAge(ageReq)
    if len(ageErr) != 0 && !ageOk {
        fail := views_alert.AlertFailure("Failed to save person:", ageErr)
        return template.AssertRender(c, http.StatusOK, fail)
    }

    ageNunber, err := strconv.ParseUint(ageReq, 10, 32); if err != nil {
        req.Age = 0
    }

    descriptionReq := c.FormValue("description")
    dateString := c.FormValue("birthDate")
    birthDate, err := time.Parse("2006-01-02", dateString)
    if err != nil {
        fmt.Println("Error parsing date: ", err)
        fail := views_alert.AlertFailure("Failed to save person:", "Error parsing date")
        return template.AssertRender(c, http.StatusOK, fail)
    }
    req.BirthDate = birthDate
    req.Name = nameReq
    req.Age = int(ageNunber)
    req.Description = descriptionReq

    println("req is ")
    println(req.Name)
    println(req.Age)
    println(req.BirthDate.String())
    println(req.Description)

    opts := views_variables.DashboardOpts{
        Nav: nav(0),
    }

    vm := views_dashboards_persons_new.NewPersonVM{
        Opts: opts,
    }

    imageURL := ""
    file, err := c.FormFile("imageData")
    if err == nil {
        fmt.Println("Not Error")
        fmt.Println(file.Size)
        newImageURL, err := fe.ImageService.UploadImage(file, "uploads")
        if err != nil {
            imageURL = ""
        }
        imageURL = fe.BaseURL + newImageURL
    }


    personsNew := views_dashboards_persons_new.New(vm)
    personData := models.Person{
        Name: req.Name,
        Age: req.Age,
        BirthDate: req.BirthDate,
        ImageURL: imageURL,
        Description: req.Description,
    };
    fe.PersonRepo.Insert(c, &personData)
    fmt.Println("test user")
    fmt.Println(personData)
    if len(personData.ID) > 0 {
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
    } else {
        return template.AssertRender(c, http.StatusOK, personsNew)
    }
}
