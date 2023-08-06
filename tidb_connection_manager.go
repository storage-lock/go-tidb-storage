package tidb_storage

import (
	"database/sql"
	mysql_storage "github.com/storage-lock/go-mysql-storage"
	"github.com/storage-lock/go-storage"
)

// TidbConnectionManager 创建一个TIDB的连接
type TidbConnectionManager struct {
	// tidb底层实际上都是跟mysql通用的
	storage.ConnectionManager[*sql.DB]
}

var _ storage.ConnectionManager[*sql.DB] = &TidbConnectionManager{}

// NewTidbConnectionProviderFromDSN 从DSN创建tidb连接
func NewTidbConnectionProviderFromDSN(dsn string) *TidbConnectionManager {
	return &TidbConnectionManager{
		ConnectionManager: mysql_storage.NewMySQLConnectionManagerFromDSN(dsn),
	}
}

// NewTidbStorageConnectionProvider 从服务器属性创建数据库连接
func NewTidbStorageConnectionProvider(host string, port uint, user, passwd, databaseName string) *TidbConnectionManager {
	return &TidbConnectionManager{
		ConnectionManager: mysql_storage.NewMySQLConnectionProvider(host, port, user, passwd, databaseName),
	}
}
