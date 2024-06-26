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
    query := `
        INSERT INTO "Person" ("name", "age", "birth_date", "image_url", "description")
        VALUES ($1, $2, $3, $4, $5)
        RETURNING "id";
    `
    if err := r.db.QueryRow(
        query,
        value.Name,
        value.Age,
        value.BirthDate,
        value.ImageURL,
        value.Description,
    ).Scan(&value.ID); err != nil {
        return err
    }
    return nil
}

func (r *repository) Count() int {
    var count int
    err := r.db.Get(&count, `SELECT COUNT(*) FROM "Person"`)
    if err != nil {
        return 0
    }

    return count
}

func (r *repository) Get(limitQuery string, argsData []any) ([]models.Person, error) {
    var persons []models.Person
    baseQuery := `SELECT id, name, age, birth_date, image_url, description
        FROM "Person"
        ORDER BY id`

    fullQuery := baseQuery + " " + limitQuery
    fmt.Printf("Executing SQL query: %s with args: %v\n", fullQuery, argsData)
    err := r.db.Select(&persons, fullQuery, argsData...)
    if err != nil {
        return nil, err
    }
    return persons, nil
}

func NewPersonRepository(db *sqlx.DB) models.PersonRepository {
    return &repository{
        db: db,
    }
}
