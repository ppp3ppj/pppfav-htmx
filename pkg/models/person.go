package models

import (
	"time"
	"github.com/labstack/echo/v4"
)

type Person struct {
    ID          string    `json:"id" db:"id"`                   // UUID stored as string
    Name        string    `json:"name" db:"name"`
    Age         int       `json:"age" db:"age"`
    BirthDate   time.Time `json:"birth_date" db:"birth_date"`   // Uses time.Time to represent date
    ImageURL    string    `json:"image_url" db:"image_url"`     // URL to the image of the person
    Description string    `json:"description" db:"description"` // Description or biography of the person
}

type PersonRepository interface {
    Insert(ctx echo.Context, value *Person) error
    Count() int
}
