package tidb_storage

import (
	"context"
	mysql_storage "github.com/storage-lock/go-mysql-storage"
	storage_test_helper "github.com/storage-lock/go-storage-test-helper"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNewTidbStorage(t *testing.T) {
	envName := "STORAGE_LOCK_TIDB_DSN"
	dsn := os.Getenv(envName)
	assert.NotEmpty(t, dsn)
	connectionGetter := NewTidbConnectionManagerFromDSN(dsn)
	storage, err := NewTidbStorage(context.Background(), &TidbStorageOptions{
		MySQLStorageOptions: &mysql_storage.MySQLStorageOptions{
			ConnectionManager: connectionGetter,
			TableName:         storage_test_helper.TestTableName,
		},
	})
	assert.Nil(t, err)
	storage_test_helper.TestStorage(t, storage)
}
