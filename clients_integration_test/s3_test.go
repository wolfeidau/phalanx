//go:build integration

package clients_integration_test

import (
	"path/filepath"
	"testing"

	"github.com/joho/godotenv"
	"github.com/mosuka/phalanx/clients"
)

func TestNewS3ClientWithUri(t *testing.T) {
	err := godotenv.Load(filepath.FromSlash("../.env"))
	if err != nil {
		t.Errorf("Failed to load .env file")
	}

	uri := "s3://phalanx-locks-test/indexname/shardname"

	if _, err := clients.NewS3ClientWithUri(uri); err != nil {
		t.Fatalf("error %v\n", err)
	}
}
