package tidb_storage

import (
	"context"
	mysql_storage "github.com/storage-lock/go-mysql-storage"
	"github.com/storage-lock/go-storage"

	_ "github.com/go-sql-driver/mysql"
)

type TidbStorage struct {
	*mysql_storage.MysqlStorage
}

var _ storage.Storage = &TidbStorage{}

func NewTidbStorage(ctx context.Context, options *TidbStorageOptions) (*TidbStorage, error) {

	if err := options.Check(); err != nil {
		return nil, err
	}

	mysqlStorage, err := mysql_storage.NewMysqlStorage(ctx, options.toMysqlStorageOptions())
	if err != nil {
		return nil, err
	}

	s := &TidbStorage{
		MysqlStorage: mysqlStorage,
	}

	return s, nil
}

const StorageName = "tidb-storage"

func (x *TidbStorage) GetName() string {
	return StorageName
}

//func (x *TidbStorage) Init(ctx context.Context) (returnError error) {
//	db, err := x.options.ConnectionManager.Take(ctx)
//	if err != nil {
//		return err
//	}
//	defer func() {
//		err := x.options.ConnectionManager.Return(ctx, db)
//		if returnError == nil {
//			returnError = err
//		}
//	}()
//
//	// TODO 要不要自动创建数据库呢？这是一个值得讨论的问题。
//	// 用户有可能是想把数据库连接放到当前的数据库下，也可能是想放到别的数据库下
//	// 如果想放到别的数据库下，用户应该为其创建专门的数据库
//	// 如果是复用连接的话，则有可能会有需求是切换数据库，也许这里只应该标记一下，作为能够用之后的优化项
//
//	// 创建存储锁信息需要的表
//	// TODO 这个参数后面涉及到多处拼接sql，可能会有sql注入，是否需要做一些安全措施？
//	tableFullName := x.options.TableName
//	if tableFullName == "" {
//		tableFullName = fmt.Sprintf("`%s`.`%s`", storage.DefaultStorageDatabaseName, storage.DefaultStorageTableName)
//	}
//	createTableSql := `CREATE TABLE IF NOT EXISTS %s (
//    lock_id VARCHAR(255) NOT NULL PRIMARY KEY,
//    owner_id VARCHAR(255) NOT NULL,
//    version BIGINT NOT NULL,
//    lock_information_json_string VARCHAR(255) NOT NULL
//)`
//	_, err = db.Exec(fmt.Sprintf(createTableSql, tableFullName))
//	if err != nil {
//		return err
//	}
//
//	x.tableFullName = tableFullName
//
//	return nil
//}
