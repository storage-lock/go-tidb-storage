package tidb_storage

import (
	"context"
	"testing"

	"github.com/storage-lock/go-storage"
	storage_lock "github.com/storage-lock/go-storage-lock"
	"github.com/stretchr/testify/assert"
)

func TestTidbStorageMysqlCompatSmoke(t *testing.T) {
	dsn := "root:rootpass@tcp(127.0.0.1:13307)/storage_lock_test"
	ctx := context.Background()
	cm := NewTidbConnectionManagerFromDsn(dsn)
	s, err := NewTidbStorage(ctx, NewTidbStorageOptions().SetConnectionManager(cm).SetTableName("tidb_smoke_test"))
	assert.Nil(t, err)
	caps := s.Capabilities()
	assert.Contains(t, caps, storage.CapabilityCAS)
	info := &storage.LockInformation{LockId: "smoke-lock", OwnerId: "owner-1", Version: 1}
	_ = s.DeleteWithVersion(ctx, "smoke-lock", 999, info)
	assert.Nil(t, s.CreateWithVersion(ctx, "smoke-lock", 1, info))
	lock, err := storage_lock.NewStorageLock(s, "smoke-lock-2")
	assert.Nil(t, err)
	assert.Nil(t, lock.Lock(ctx, "smoke-owner"))
	assert.Nil(t, lock.UnLock(ctx, "smoke-owner"))
	_ = s.DeleteWithVersion(ctx, "smoke-lock", 1, info)
}
