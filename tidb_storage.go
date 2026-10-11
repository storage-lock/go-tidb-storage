package tidb_storage

import (
	"context"
	mysql_storage "github.com/storage-lock/go-mysql-storage"
	"github.com/storage-lock/go-storage"
	"time"

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

func (x *TidbStorage) Capabilities() []storage.StorageCapability {
	return x.MysqlStorage.Capabilities()
}

func (x *TidbStorage) Watch(ctx context.Context, lockId string) (<-chan *storage.LockInformation, error) {
	return storage.NewPollingLockWatcher(x, 500*time.Millisecond).Watch(ctx, lockId)
}
