package estore

import (
	"os"
	"testing"
)

var repo *Repo

func TestMain(m *testing.M) {
	setupRepo()
	result := m.Run()
	os.Exit(result)
}

func setupRepo() {
	dbUrl := os.Getenv(DbUrlEnvVar)
	if dbUrl == "" {
		panic(DbUrlEnvVar + " env var must be set!")
	}
	repo = NewRepo(dbUrl)
}
