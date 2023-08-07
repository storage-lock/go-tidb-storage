package tidb_storage

import (
	"context"
	"database/sql"
	mysql_storage "github.com/storage-lock/go-mysql-storage"
	"github.com/storage-lock/go-storage"
	"sync"
)

// TidbConnectionManager 创建一个TIDB的连接
type TidbConnectionManager struct {
	// tidb底层实际上都是跟mysql通用的
	once *sync.Once
	*mysql_storage.MySQLConnectionManager
}

var _ storage.ConnectionManager[*sql.DB] = &TidbConnectionManager{}

// NewTidbConnectionManagerFromDSN 从DSN创建TiDB连接管理器
func NewTidbConnectionManagerFromDSN(dsn string) *TidbConnectionManager {
	return &TidbConnectionManager{
		once:                   &sync.Once{},
		MySQLConnectionManager: mysql_storage.NewMySQLConnectionManagerFromDSN(dsn),
	}
}

// NewTidbConnectionManager 从连接属性创建数据库连接
func NewTidbConnectionManager(host string, port uint, user, passwd, database string) *TidbConnectionManager {
	return &TidbConnectionManager{
		once:                   &sync.Once{},
		MySQLConnectionManager: mysql_storage.NewMySQLConnectionManager(host, port, user, passwd, database),
	}
}

func (x *TidbConnectionManager) SetHost(host string) *TidbConnectionManager {
	x.Host = host
	return x
}

func (x *TidbConnectionManager) SetPort(port uint) *TidbConnectionManager {
	x.Port = port
	return x
}

func (x *TidbConnectionManager) SetUser(user string) *TidbConnectionManager {
	x.User = user
	return x
}

func (x *TidbConnectionManager) SetPasswd(passwd string) *TidbConnectionManager {
	x.Passwd = passwd
	return x
}

func (x *TidbConnectionManager) SetDatabaseName(databaseName string) *TidbConnectionManager {
	x.DatabaseName = databaseName
	return x
}

const TiDBConnectionManagerName = "tidb-connection-manager"

func (x *TidbConnectionManager) Name() string {
	return TiDBConnectionManagerName
}

// Take 获取到数据库的连接
func (x *TidbConnectionManager) Take(ctx context.Context) (*sql.DB, error) {
	return x.MySQLConnectionManager.Take(ctx)
}

func (x *TidbConnectionManager) GetDSN() string {
	return x.MySQLConnectionManager.GetDSN()
}

func (x *TidbConnectionManager) Return(ctx context.Context, db *sql.DB) error {
	return x.MySQLConnectionManager.Return(ctx, db)
}

func (x *TidbConnectionManager) Shutdown(ctx context.Context) error {
	return x.MySQLConnectionManager.Shutdown(ctx)
}
