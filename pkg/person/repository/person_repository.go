package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/ppp3ppj/pppfav-htmx/pkg/models"
)

type repository struct {
    db *sqlx.DB
}

func (r *repository) Insert(ctx echo.Context, value *models.Person) error {
    fmt.Println("Hii Insert This...")
    return nil
}

func NewPersonRepository(db *sqlx.DB) models.PersonRepository {
    return &repository{
        db: db,
    }
}
