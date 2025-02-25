//go:build integration

package lock_integration_test

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/joho/godotenv"
	"github.com/mosuka/phalanx/lock"
	"github.com/mosuka/phalanx/logging"
	"github.com/thanhpk/randstr"
)

func TestEtcdLockManagerWithUri(t *testing.T) {
	err := godotenv.Load(filepath.FromSlash("../.env"))
	if err != nil {
		t.Errorf("Failed to load .env file")
	}

	tmpDir := randstr.String(8)
	uri := fmt.Sprintf("etcd://phalanx-test/locks/%s", tmpDir)
	logger := logging.NewLogger("WARN", "", 500, 3, 30, false)

	etcdLock, err := lock.NewEtcdLockManagerWithUri(uri, logger)
	if err != nil {
		t.Fatalf("error %v\n", err)
	}
	defer etcdLock.Close()
}

func TestEtcdLockManagerLock(t *testing.T) {
	err := godotenv.Load(filepath.FromSlash("../.env"))
	if err != nil {
		t.Errorf("Failed to load .env file")
	}

	tmpDir := randstr.String(8)
	uri := fmt.Sprintf("etcd://phalanx-test/locks/%s", tmpDir)
	logger := logging.NewLogger("WARN", "", 500, 3, 30, false)

	etcdLock, err := lock.NewEtcdLockManagerWithUri(uri, logger)
	if err != nil {
		t.Fatalf("%v\n", err)
	}
	defer etcdLock.Close()

	rev, err := etcdLock.Lock()
	if err != nil {
		t.Fatalf("%v\n", err)
	}
	if rev == 0 {
		t.Fatalf("expecting the revision greater than 0.\n")
	}

	defer etcdLock.Unlock()
}

func TestEtcdLockManagerLockTimeout(t *testing.T) {
	err := godotenv.Load(filepath.FromSlash("../.env"))
	if err != nil {
		t.Errorf("Failed to load .env file")
	}

	tmpDir := randstr.String(8)
	uri := fmt.Sprintf("etcd://phalanx-test/locks/%s", tmpDir)
	logger := logging.NewLogger("WARN", "", 500, 3, 30, false)

	etcdLock1, err := lock.NewEtcdLockManagerWithUri(uri, logger)
	if err != nil {
		t.Fatalf("%v\n", err)
	}
	defer etcdLock1.Close()

	etcdLock2, err := lock.NewEtcdLockManagerWithUri(uri, logger)
	if err != nil {
		t.Fatalf("%v\n", err)
	}
	defer etcdLock2.Close()

	rev, err := etcdLock1.Lock()
	if err != nil {
		t.Fatalf("%v\n", err)
	}
	if rev == 0 {
		t.Fatalf("expecting the revision greater than 0.\n")
	}

	_, err = etcdLock2.Lock()
	if err == nil {
		t.Fatalf("expect error: contextdeadline exceeded\n")
	}

	defer etcdLock1.Unlock()
	defer etcdLock2.Unlock()
}
