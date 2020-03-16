package estore

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Repo struct {
	db *sql.DB
}

type StreamDef struct {
	Key string
}

type Stream struct {
	Key     string
	Version uint
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

func (r *Repo) CreateStream(def StreamDef) (Stream, error) {
	stream := Stream{Key: def.Key}
	tx, txErr := r.db.Begin()
	if txErr != nil {
		return Stream{}, fmt.Errorf("cannot start DB transaction: %w", txErr)
	}
	row := tx.QueryRow("insert into STREAMS (KEY, VERSION) values ($1, 1) returning VERSION", def.Key)
	insertErr := row.Scan(&stream.Version)
	if insertErr != nil {
		return stream, fmt.Errorf("cannot insert stream %v into DB: %w", def, insertErr)
	}
	commitErr := tx.Commit()
	if commitErr != nil {
		return stream, fmt.Errorf("cannot commit transaction after inserting stream %v into DB: %w", def, commitErr)
	}
	return stream, nil
}

func (r *Repo) RetrieveStream(key string) (Stream, error) {
	stream := Stream{}
	row := r.db.QueryRow("select s.key, s.version from STREAMS s where KEY = $1", key)
	selectErr := row.Scan(&stream.Key, &stream.Version)
	if selectErr != nil {
		return stream, fmt.Errorf("cannot retrieve stream [%v] from DB: %w", key, selectErr)
	}
	return stream, nil
}
