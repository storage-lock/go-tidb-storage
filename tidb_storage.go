package tidb_storage

import (
	"context"
	"github.com/golang-infrastructure/go-iterator"
	mysql_storage "github.com/storage-lock/go-mysql-storage"
	"github.com/storage-lock/go-storage"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// TidbStorage 把锁存储在Tidb数据库中
type TidbStorage struct {

	// 其实底层实现跟MySQL是一样一样的，这里就直接复用mysql的storage的逻辑了
	*mysql_storage.MySQLStorage

	// 创建storage的选项
	options *TidbStorageOptions
}

var _ storage.Storage = &TidbStorage{}

// NewTidbStorage 创建一个基于tidb的storage
func NewTidbStorage(ctx context.Context, options *TidbStorageOptions) (*TidbStorage, error) {

	mysqlStorage, err := mysql_storage.NewMySQLStorage(ctx, options.MySQLStorageOptions)
	if err != nil {
		return nil, err
	}

	s := &TidbStorage{
		options:      options,
		MySQLStorage: mysqlStorage,
	}

	err = s.Init(ctx)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (x *TidbStorage) Init(ctx context.Context) error {
	return x.MySQLStorage.Init(ctx)
}

func (x *TidbStorage) UpdateWithVersion(ctx context.Context, lockId string, exceptedVersion, newVersion storage.Version, lockInformation *storage.LockInformation) error {
	return x.MySQLStorage.UpdateWithVersion(ctx, lockId, exceptedVersion, newVersion, lockInformation)
}

func (x *TidbStorage) InsertWithVersion(ctx context.Context, lockId string, version storage.Version, lockInformation *storage.LockInformation) error {
	return x.MySQLStorage.CreateWithVersion(ctx, lockId, version, lockInformation)
}

func (x *TidbStorage) DeleteWithVersion(ctx context.Context, lockId string, exceptedVersion storage.Version, lockInformation *storage.LockInformation) error {
	return x.MySQLStorage.DeleteWithVersion(ctx, lockId, exceptedVersion, lockInformation)
}

func (x *TidbStorage) Get(ctx context.Context, lockId string) (string, error) {
	return x.MySQLStorage.Get(ctx, lockId)
}

func (x *TidbStorage) GetTime(ctx context.Context) (time.Time, error) {
	return x.MySQLStorage.GetTime(ctx)
}

func (x *TidbStorage) Close(ctx context.Context) error {
	return x.MySQLStorage.Close(ctx)
}

func (x *TidbStorage) List(ctx context.Context) (iterator.Iterator[*storage.LockInformation], error) {
	return x.MySQLStorage.List(ctx)
}
