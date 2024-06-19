# pppfav-htmx
Go + HTMX(templ)
## How to run

    air -c .air.dev.toml


## How to migration

    go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

    make migrate-up

## ISSUE

    migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" force 1

    and then make migrate-{up - down}
