#!/usr/bin/env bash

docker rm -f db-storage-tidb-server

# 启动TiDB实例，默认的用户名为root，密码为空，监听在4000端口
docker run --name db-storage-tidb-server -d -p 4000:4000 -p 10080:10080 pingcap/tidb:v6.2.0

export STORAGE_LOCK_TIDB_DSN="root:@tcp(127.0.0.1:4000)/storage_lock_test"
