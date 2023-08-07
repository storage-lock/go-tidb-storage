package tidb_storage

import (
	"database/sql"
	mysql_storage "github.com/storage-lock/go-mysql-storage"
	"github.com/storage-lock/go-storage"
)

type TidbStorageOptions struct {
	*mysql_storage.MySQLStorageOptions
}

func NewTidbStorageOptions() *TidbStorageOptions {
	return &TidbStorageOptions{mysql_storage.NewMySQLStorageOptions()}
}

func (x *TidbStorageOptions) SetConnectionManager(connManager storage.ConnectionManager[*sql.DB]) *TidbStorageOptions {
	x.ConnectionManager = connManager
	return x
}

func (x *TidbStorageOptions) SetTableName(tableName string) *TidbStorageOptions {
	x.TableName = tableName
	return x
}
