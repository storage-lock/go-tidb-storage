package tidb_storage

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/storage-lock/go-storage"
	"github.com/stretchr/testify/assert"
)

func TestNewTidbStorageOptions(t *testing.T) {
	opts := NewTidbStorageOptions()
	assert.NotNil(t, opts)
	assert.Equal(t, storage.DefaultStorageTableName, opts.TableName)
	assert.Nil(t, opts.ConnectionManager)
}

func TestTidbStorageOptions_SetConnectionManager(t *testing.T) {
	opts := NewTidbStorageOptions()
	assert.NotNil(t, opts)
	db := &sql.DB{}
	connManager := storage.NewFixedSqlDBConnectionManager(db)
	opts.SetConnectionManager(connManager)
	assert.NotNil(t, opts.ConnectionManager)
}

func TestTidbStorageOptions_SetTableName(t *testing.T) {
	opts := NewTidbStorageOptions()
	assert.NotNil(t, opts)
	opts.SetTableName("custom_table")
	assert.Equal(t, "custom_table", opts.TableName)
}

func TestTidbStorageOptions_Check_NilConnectionManager(t *testing.T) {
	opts := NewTidbStorageOptions()
	assert.NotNil(t, opts)
	err := opts.Check()
	assert.Error(t, err)
	assert.Equal(t, fmt.Errorf("ConnectionManager can not nil"), err)
}

func TestTidbStorageOptions_Check_ValidOptions(t *testing.T) {
	opts := NewTidbStorageOptions()
	assert.NotNil(t, opts)
	db := &sql.DB{}
	connManager := storage.NewFixedSqlDBConnectionManager(db)
	opts.SetConnectionManager(connManager)
	err := opts.Check()
	assert.NoError(t, err)
}

func TestTidbStorageOptions_ToMysqlStorageOptions(t *testing.T) {
	opts := NewTidbStorageOptions()
	assert.NotNil(t, opts)
	db := &sql.DB{}
	connManager := storage.NewFixedSqlDBConnectionManager(db)
	opts.SetConnectionManager(connManager)
	opts.SetTableName("custom_table")
	mysqlOpts := opts.toMysqlStorageOptions()
	assert.NotNil(t, mysqlOpts)
	assert.Equal(t, "custom_table", mysqlOpts.TableName)
	assert.Equal(t, connManager, mysqlOpts.ConnectionManager)
}

func TestNewTidbStorage_NilOptions(t *testing.T) {
	assert.Panics(t, func() {
		NewTidbStorage(context.Background(), nil)
	})
}

func TestTidbStorage_GetName(t *testing.T) {
	s := &TidbStorage{}
	name := s.GetName()
	assert.Equal(t, StorageName, name)
}

func TestTidbStorage_Capabilities(t *testing.T) {
	s := &TidbStorage{}
	caps := s.Capabilities()
	assert.NotNil(t, caps)
	assert.Len(t, caps, 3)
	assert.Equal(t, storage.CapabilityCAS, caps[0])
	assert.Equal(t, storage.CapabilityReliableTime, caps[1])
	assert.Equal(t, storage.CapabilityWatch, caps[2])
}
