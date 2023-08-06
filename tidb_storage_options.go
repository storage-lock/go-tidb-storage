package tidb_storage

import mysql_storage "github.com/storage-lock/go-mysql-storage"

type TidbStorageOptions struct {
	*mysql_storage.MySQLStorageOptions
}

func NewTidbStorageOptions() *TidbStorageOptions {
	return &TidbStorageOptions{mysql_storage.NewMySQLStorageOptions()}
}
