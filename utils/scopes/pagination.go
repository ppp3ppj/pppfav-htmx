package scopes

import "fmt"

func Paginate(page, pageSize int) (string, []any) {
    fmt.Printf("%d and %d\n", page, pageSize)
    if page <= 0 {
        page = 1
    }

    switch {
    case pageSize > 100:
        pageSize = 100
    case pageSize <= 0:
        pageSize = 10
    }

    offset := (page - 1) * pageSize
    query := "LIMIT $1 OFFSET $2"
    args := []any{pageSize, offset}

    return query, args
}
