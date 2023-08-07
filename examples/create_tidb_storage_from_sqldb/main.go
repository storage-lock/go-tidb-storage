package main

import (
	"context"
	"database/sql"
	"fmt"
	storage "github.com/storage-lock/go-storage"
	tidb_storage "github.com/storage-lock/go-tidb-storage"
)

func main() {

	// 假设已经在其它地方初始化数据库连接得到了一个*sql.DB
	testDsn := "root:@tcp(127.0.0.1:4000)/storage_lock_test"
	db, err := sql.Open("mysql", testDsn)
	if err != nil {
		panic(err)
	}

	// 则可以从这个*sql.DB中创建一个Tidb Storage
	connectionManager := storage.NewFixedSqlDBConnectionManager(db)
	options := tidb_storage.NewTidbStorageOptions().SetConnectionManager(connectionManager)
	storage, err := tidb_storage.NewTidbStorage(context.Background(), options)
	if err != nil {
		panic(err)
	}
	fmt.Println(storage.GetName())

}
