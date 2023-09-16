package tidb_storage

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/storage-lock/go-storage"
	"sync"
)

// TidbConnectionManager 创建一个MySQL的连接管理器
type TidbConnectionManager struct {

	// 主机的名字
	Host string

	// 主机的端口
	Port uint

	// 用户名
	User string

	// 密码
	Passwd string

	DatabaseName string

	DSN string

	// 初始化好的数据库实例
	db   *sql.DB
	err  error
	once sync.Once
}

var _ storage.ConnectionManager[*sql.DB] = &TidbConnectionManager{}

// NewTidbConnectionManagerFromDsn 从DSN创建MySQL连接管理器
func NewTidbConnectionManagerFromDsn(dsn string) *TidbConnectionManager {
	return &TidbConnectionManager{
		DSN: dsn,
	}
}

func NewTidbConnectionManagerFromSqlDb(db *sql.DB) *TidbConnectionManager {
	return &TidbConnectionManager{
		db: db,
	}
}

// NewTidbConnectionManager 从连接属性创建数据库连接
func NewTidbConnectionManager(host string, port uint, user, passwd, database string) *TidbConnectionManager {
	return &TidbConnectionManager{
		Host:         host,
		Port:         port,
		User:         user,
		Passwd:       passwd,
		DatabaseName: database,
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

const TidbConnectionManagerName = "mysql-connection-manager"

func (x *TidbConnectionManager) Name() string {
	return TidbConnectionManagerName
}

// Take 获取到数据库的连接
func (x *TidbConnectionManager) Take(ctx context.Context) (*sql.DB, error) {
	x.once.Do(func() {
		if x.db != nil {
			return
		}
		db, err := sql.Open("mysql", x.GetDSN())
		if err != nil {
			x.err = err
			return
		}
		x.db = db
	})
	return x.db, x.err
}

func (x *TidbConnectionManager) GetDSN() string {
	if x.DSN != "" {
		return x.DSN
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", x.User, x.Passwd, x.Host, x.Port, x.DatabaseName)
}

func (x *TidbConnectionManager) Return(ctx context.Context, db *sql.DB) error {
	return nil
}

func (x *TidbConnectionManager) Shutdown(ctx context.Context) error {
	if x.db != nil {
		return x.db.Close()
	}
	return nil
}
