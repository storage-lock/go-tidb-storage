package main

import (
	"context"
	"fmt"

	tidb_storage "github.com/storage-lock/go-tidb-storage"
)

func main() {

	// 使用一个DSN形式的数据库连接字符串创建ConnectionManager
	testDsn := "root:@tcp(127.0.0.1:4000)/storage_lock_test"
	connectionManager := tidb_storage.NewTidbConnectionManagerFromDsn(testDsn)

	// 然后从这个ConnectionManager创建Tidb Storage
	options := tidb_storage.NewTidbStorageOptions().SetConnectionManager(connectionManager)
	storage, err := tidb_storage.NewTidbStorage(context.Background(), options)
	if err != nil {
		panic(err)
	}
	fmt.Println(storage.GetName())

}
