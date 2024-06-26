package scopes

func Paginate(page, pageSize int) (string, []interface{}) {
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
    args := []interface{}{pageSize, offset}

    return query, args
}
