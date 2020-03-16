package estore

import (
	"flag"
	"github.com/google/uuid"
	"log"
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
	var dbUrl string
	flag.StringVar(&dbUrl, "db", "", "URL-encoded database connection string (required)")
	flag.Parse()
	if dbUrl == "" {
		log.Fatal("Parameter 'db' is required!")
	}
	log.Printf("Running tests with database [%s]", dbUrl)
	repo = NewRepo(dbUrl)
}

func TestCreateStream(t *testing.T) {
	streamKey := uuid.New().String()
	def := StreamDef{
		Key: streamKey,
	}
	newStream, createErr := repo.CreateStream(def)
	if createErr != nil {
		t.Fatalf("cannot create stream because of error: %v", createErr)
	}
	if newStream.Key != def.Key {
		t.Errorf("expected that created stream has Key [%s], but got [%s]", def.Key, newStream.Key)
	}
	if newStream.Version != 1 {
		t.Errorf("expected that created stream has version [1], but got [%d]", newStream.Version)
	}
	stream, retrieveErr := repo.RetrieveStream(streamKey)
	if retrieveErr != nil {
		t.Fatalf("cannot retrieve stream because of error: %v", retrieveErr)
	}
	if stream.Key != def.Key {
		t.Errorf("expected that retrieved stream has Key [%s], but got [%s]", def.Key, stream.Key)
	}
	if stream.Version != 1 {
		t.Errorf("expected that retrieved stream has version [1], but got [%d]", stream.Version)
	}
}
