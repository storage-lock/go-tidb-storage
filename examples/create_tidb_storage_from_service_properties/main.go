package main

import (
	"context"
	"fmt"
	tidb_storage "github.com/storage-lock/go-tidb-storage"
)

func main() {

	// 数据库连接不是DSN的形式，就是一堆零散的属性，则依次设置，可以得到一个连接管理器
	host := "127.0.0.1"
	port := uint(4000)
	username := "root"
	passwd := ""
	database := "storage_lock_test"
	connectionManager := tidb_storage.NewTidbConnectionManager(host, port, username, passwd, database)

	// 然后从这个连接管理器创建Tidb Storage
	options := tidb_storage.NewTidbStorageOptions().SetConnectionManager(connectionManager)
	storage, err := tidb_storage.NewTidbStorage(context.Background(), options)
	if err != nil {
		panic(err)
	}
	fmt.Println(storage.GetName())

}
