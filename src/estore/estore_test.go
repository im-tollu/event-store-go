package estore

import (
	"github.com/google/uuid"
	"log"
	"os"
	"testing"
)

const envVar = "ESTORE_DB_URL"

var repo *Repo

func TestMain(m *testing.M) {
	setupRepo()
	result := m.Run()
	os.Exit(result)
}

func setupRepo() {
	dbUrl := os.Getenv(envVar)
	if dbUrl == "" {
		log.Fatal("Provide database connection string in environment variable " + envVar)
	}
	log.Printf("Running tests with database [%s]", dbUrl)
	repo = NewRepo(dbUrl)
}

func TestRetrieveStreamNotFound(t *testing.T) {
	streamKey := "non-existing-stream"
	_, retrieveErr := repo.RetrieveStream(streamKey)
	notFound, ok := retrieveErr.(StreamNotFoundErr)
	if !ok {
		t.Fatalf("expected StreamNotFoundErr but got %v", retrieveErr)
	}
	if notFound.Key != streamKey {
		t.Fatalf("expected key to be [%s] but got [%s]", streamKey, notFound.Key)
	}

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
	_, createExistingErr := repo.CreateStream(def)
	alreadyExistsErr, ok := createExistingErr.(StreamAlreadyExistsErr)
	if !ok {
		t.Fatalf("expected StreamAlreadyExistsError, but got %v", createExistingErr)
	}
	if alreadyExistsErr.Key != def.Key {
		t.Errorf("expected Key to be [%s], but got [%s]", def.Key, alreadyExistsErr.Key)
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
