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

// CheckNameExists checks if a name exists in the Person table
func (r *repository) CheckNameExists(name string) (bool, error) {
    fmt.Printf("%s%s", "Name is db, ", name)
    var exists bool
    query := `SELECT EXISTS(SELECT 1 FROM "Person" WHERE "name"=$1)`
    err := r.db.Get(&exists, query, name)
    if err != nil {
        return false, fmt.Errorf("error checking name existence: %v", err)
    }
    return exists, nil
}


func NewPersonRepository(db *sqlx.DB) models.PersonRepository {
    return &repository{
        db: db,
    }
}
