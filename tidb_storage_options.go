package tidb_storage

import (
	"database/sql"

	"github.com/storage-lock/go-storage"
)

// TidbStorageOptions 基于MySQL为存储引擎时的选项
type TidbStorageOptions struct {

	// 存放锁的表的名字，如果未指定的话则使用默认的表
	TableName string

	// 用于获取数据库连接
	ConnectionManager storage.ConnectionManager[*sql.DB]
}

func NewTidbStorageOptions() *TidbStorageOptions {
	return &TidbStorageOptions{
		TableName: storage.DefaultStorageTableName,
	}
}

func (x *TidbStorageOptions) SetConnectionManager(connManager storage.ConnectionManager[*sql.DB]) *TidbStorageOptions {
	x.ConnectionManager = connManager
	return x
}

func (x *TidbStorageOptions) SetTableName(tableName string) *TidbStorageOptions {
	x.TableName = tableName
	return x
}
