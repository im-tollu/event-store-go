package main

import (
	"flag"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
	"strings"
)

func main() {
	var (
		dbUrl      string
		migrations string
		clean      bool
	)
	flag.StringVar(&dbUrl, "db", "", "URL-encoded database connection string (required)")
	flag.StringVar(&migrations, "migrations", "migrations", "path to migration scripts (required)")
	flag.BoolVar(&clean, "clean", false, "Drop the database and recreate from scratch")
	flag.Parse()
	if dbUrl == "" {
		log.Fatal("Parameter 'db' is required!")
	}
	if migrations == "" {
		log.Fatal("Parameter 'migrations' is required!")
	}
	m := getMigration(dbUrl, migrations)
	defer closeMigrate(m)
	if clean {
		cleanDb(m)
	}
	migrateDb(m)
}

func cleanDb(m *migrate.Migrate) {
	log.Println("Cleaning database...")
	err := m.Drop()
	if err != nil {
		log.Fatalf("Could not clean database: %v", err)
	}
	log.Println("Database clean.")
}

func migrateDb(m *migrate.Migrate) {
	log.Println("Migrating database...")
	err := m.Up()
	if err != nil {
		log.Fatalf("Could not migrate the database up: %v", err)
	}
	log.Println("Database migrated.")
}

type logger struct {
}

func (l logger) Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func (l logger) Verbose() bool {
	return false
}

func getMigration(dbUrl string, migrations string) *migrate.Migrate {
	m, readMigrationErr := migrate.New("file://"+migrations, dbUrl)
	if readMigrationErr != nil {
		log.Fatalf("could not instantiate migration: %v", readMigrationErr)
	}
	m.Log = logger{}
	return m
}

func closeMigrate(m *migrate.Migrate) {
	srcErr, dbErr := m.Close()
	errs := make([]string, 0)
	if srcErr != nil {
		errs = append(errs, srcErr.Error())
	}
	if dbErr != nil {
		errs = append(errs, dbErr.Error())
	}
	if len(errs) > 0 {
		log.Fatalf("Could not finish Migrate: \n%s", strings.Join(errs, "\n"))
	}
}
