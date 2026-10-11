package tidb_storage

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/storage-lock/go-storage"
	"github.com/stretchr/testify/assert"
)

type takeErrorManager struct{}

func (takeErrorManager) Name() string                          { return "take-error" }
func (takeErrorManager) Take(context.Context) (*sql.DB, error) { return nil, errors.New("take failed") }
func (takeErrorManager) Return(context.Context, *sql.DB) error { return nil }
func (takeErrorManager) Shutdown(context.Context) error        { return nil }

func TestTidbConnectionManager(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	mock.ExpectClose()

	fromDB := NewTidbConnectionManagerFromSqlDb(db)
	got, err := fromDB.Take(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, db, got)
	assert.NoError(t, fromDB.Return(context.Background(), got))
	assert.NoError(t, fromDB.Shutdown(context.Background()))

	fromConfig := NewTidbConnectionManager("localhost", 4000, "user", "pass", "db")
	assert.Equal(t, "db.local", fromConfig.SetHost("db.local").Host)
	assert.Equal(t, TidbConnectionManagerName, fromConfig.Name())
	assert.Contains(t, fromConfig.GetDSN(), "db.local:4000")
	opened, err := fromConfig.Take(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, opened)
	again, err := fromConfig.Take(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, opened, again)
	assert.NoError(t, fromConfig.Shutdown(context.Background()))

	failedOpen := NewTidbConnectionManagerFromDsn("bad")
	failedOpen.open = func(string, string) (*sql.DB, error) { return nil, errors.New("open failed") }
	_, err = failedOpen.Take(context.Background())
	assert.EqualError(t, err, "open failed")
	_, err = failedOpen.Take(context.Background())
	assert.EqualError(t, err, "open failed")

	preset := &TidbConnectionManager{err: errors.New("open failed")}
	_, err = preset.Take(context.Background())
	assert.EqualError(t, err, "open failed")

	fromDSN := NewTidbConnectionManagerFromDsn("user:pass@tcp(localhost:4000)/db")
	assert.Equal(t, "user:pass@tcp(localhost:4000)/db", fromDSN.GetDSN())
	assert.NoError(t, (&TidbConnectionManager{}).Shutdown(context.Background()))
}

func TestTidbStorageOptions_EmptyTableName(t *testing.T) {
	opts := NewTidbStorageOptions()
	opts.TableName = ""
	opts.SetConnectionManager(takeErrorManager{})
	assert.NoError(t, opts.Check())
	assert.Equal(t, storage.DefaultStorageTableName, opts.TableName)
}

func TestNewTidbStorage_CheckError(t *testing.T) {
	_, err := NewTidbStorage(context.Background(), NewTidbStorageOptions())
	assert.Error(t, err)
}

func TestNewTidbStorage_MysqlError(t *testing.T) {
	opts := NewTidbStorageOptions().SetConnectionManager(takeErrorManager{})
	_, err := NewTidbStorage(context.Background(), opts)
	assert.EqualError(t, err, "take failed")
}

func TestNewTidbStorage_SuccessAndWatch(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	mock.ExpectExec("CREATE TABLE").WillReturnResult(sqlmock.NewResult(0, 0))
	s, err := NewTidbStorage(context.Background(), NewTidbStorageOptions().
		SetConnectionManager(storage.NewFixedSqlDBConnectionManager(db)))
	assert.NoError(t, err)
	assert.NotEmpty(t, s.Capabilities())
	_, err = s.Watch(context.Background(), "")
	assert.Error(t, err)
}
