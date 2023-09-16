package tidb_storage

import (
	"context"
	"os"
	"testing"

	storage_test_helper "github.com/storage-lock/go-storage-test-helper"
	"github.com/stretchr/testify/assert"
)

func TestNewTidbStorage(t *testing.T) {
	envName := "STORAGE_LOCK_TIDB_DSN"
	dsn := os.Getenv(envName)
	assert.NotEmpty(t, dsn)
	connectionGetter := NewTidbConnectionManagerFromDsn(dsn)
	storage, err := NewTidbStorage(context.Background(), &TidbStorageOptions{
		ConnectionManager: connectionGetter,
		TableName:         storage_test_helper.TestTableName,
	})
	assert.Nil(t, err)
	storage_test_helper.TestStorage(t, storage)
}
