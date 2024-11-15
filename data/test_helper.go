package data

import (
	"context"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"
	"gitlab.ctyuncdn.cn/ias/ias-core/biz/streaming"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/enttest"
	"testing"
)

func NewTestDB(t *testing.T) *ent.Database {
	// 创建内存数据库
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	// 创建表结构
	if err := client.Schema.Create(context.Background(), schema.WithForeignKeys(false)); err != nil {
		t.Fatal(err)
	}
	return ent.NewDatabase(client)
}

func NewTestData(t *testing.T) (*Data, func()) {
	db := NewTestDB(t)
	rdb := NewTestRedis(t)
	cleanup := func() {
		require.NoError(t, db.Close())
		require.NoError(t, rdb.Close())
	}
	return &Data{db: db, rdb: rdb}, cleanup
}

func NewTestRedis(t *testing.T) *redis.Client {
	s := miniredis.RunT(t)
	rdb := redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})
	return rdb
}

func NewTestStreamingProtocol() streaming.Protocol {
	sp, _ := streaming.NewProtocol("rtsp://0.0.0.0/test", streaming.ProtocolTypeRtsp)
	return sp
}
