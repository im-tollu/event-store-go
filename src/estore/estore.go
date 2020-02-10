package estore

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const DbUrlEnvVar = "ESTORE_DB_URL"

type Repo struct {
	db *sql.DB
}

func NewRepo(url string) *Repo {
	db, connErr := sql.Open("postgres", url)
	if connErr != nil {
		panic(fmt.Errorf("could not connect to db, url[%s]:\n%w", url, connErr))
	}
	pingErr := db.Ping()
	if pingErr != nil {
		panic(fmt.Errorf("db ping failed, url[%s]:\n%w", url, pingErr))
	}
	return &Repo{db: db}
}
