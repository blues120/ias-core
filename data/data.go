package data

import (
	"context"
	"fmt"
	"os"
	path "path/filepath"
	"strings"
	"time"

	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"

	"ariga.io/atlas/sql/migrate"
	atlas "ariga.io/atlas/sql/schema"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/lithammer/shortuuid"
	_ "github.com/xiaoqidun/entps"
	"gitlab.ctyuncdn.cn/ias/ias-core/biz"
	"gitlab.ctyuncdn.cn/ias/ias-core/conf"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/iam"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/scheduler"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewCameraRepo,
	NewTaskRepo,
	NewAlgoRepo,
	NewEntDatabase,
	NewRedisClient,
	NewMqttClient,
	NewMqttRepo,
	NewRedSync,
	NewData,
	NewOssRepo,
	NewTransaction,
	NewUserRepo,
	NewCaptchaRepo,
	NewJwtRepo,
	NewWarnPushRepo,
	NewInformRepo,
	NewDeviceRepo,
	NewYtxClient,
	NewSmsNotifyRepo,
	NewWarningAlertRepo,
	NewSystemRepo,
	NewEventSubsRepo,
	NewSignatureRepo,
	NewUpPlatformRepo,
	NewWarnTypeRepo,
	NewFileUploadRepo,
	iam.NewClient,
	NewDeviceAlgoRepo,
	NewOrganizationRepo,
	NewAreaRepo,
	NewTaskLimitsRepo,
	// scheduler
	scheduler.NewSchedulerRepo,
	scheduler.NewSchedulerVSSRepo,
	scheduler.NewSchedulerRepoSelector,
	NewImageCacheRepo,
)

type Data struct {
	db   *ent.Database
	rdb  *redis.Client
	rs   *redsync.Redsync
	mqtt mqtt.Client

	logger *log.Helper
}

// NewTransaction .
func NewTransaction(data *Data) biz.Transaction {
	return data.db
}

func NewData(db *ent.Database, rdb *redis.Client, rs *redsync.Redsync, mqtt mqtt.Client, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
		_ = db.Close()
		if rdb != nil {
			_ = rdb.Close()
		}
		if mqtt != nil {
			mqtt.Disconnect(uint(3 * time.Second.Milliseconds()))
		}
	}
	return &Data{db: db, rdb: rdb, rs: rs, mqtt: mqtt}, cleanup, nil
}

func NewEntDatabase(logConf *conf.Log, dataConf *conf.Data, logger log.Logger) (*ent.Database, error) {
	var (
		drv *sql.Driver
		err error
	)
	switch dataConf.Database.Driver {
	case dialect.MySQL:
		drv, err = sql.Open(dialect.MySQL, dataConf.Database.Source)
	case dialect.SQLite:
		drv, err = sql.Open(dialect.SQLite, dataConf.Database.Source)
		changeToWAL(drv)
	default:
		err = fmt.Errorf("unsupported database drive: %s", dataConf.Database.Driver)
	}
	if err != nil {
		return nil, err
	}

	// 连接池配置
	db := drv.DB()
	db.SetMaxIdleConns(int(dataConf.Database.MaxIdleConns))
	db.SetMaxOpenConns(int(dataConf.Database.MaxOpenConns))
	db.SetConnMaxLifetime(dataConf.Database.ConnMaxLifetime.AsDuration())

	helper := log.NewHelper(logger)

	observableDrv := NewObservableDriver(drv, func(ctx context.Context, i ...interface{}) {
		helper.WithContext(ctx).Debug(i...)
	}, strings.ToLower(logConf.Level) == "debug", true)

	// 初始化 client
	client := ent.NewClient(
		ent.Driver(observableDrv),
	)

	// 表结构迁移
	switch {
	case dataConf.Database.AutoMigration:
		if err := client.Schema.Create(context.Background(), schema.WithForeignKeys(false)); err != nil {
			log.Fatalf("auto migration err: %v", err)
		}
	default:
		if err := versionedMigration(client); err != nil {
			log.Fatalf("versioned migration err: %v", err)
		}
	}
	return ent.NewDatabase(client), nil
}

func changeToWAL(driver *sql.Driver) {
	// 执行 PRAGMA 查询以获取 WAL 模式状态
	var journalMode string
	err := driver.DB().QueryRow("PRAGMA journal_mode;").Scan(&journalMode)
	if err != nil {
		log.Fatal(err)
	}
	log.Debugf("执行 PRAGMA 查询以获取 WAL 模式状态Journal Mode:%s", journalMode)

	if journalMode != "wal" {
		_, err = driver.DB().Exec("PRAGMA journal_mode = WAL;")
		if err != nil {
			log.Fatal(err)
		}
		log.Debug("执行了WAL模式的修改")
		err = driver.DB().QueryRow("PRAGMA journal_mode;").Scan(&journalMode)
		if err != nil {
			log.Fatal(err)
		}
		log.Debugf("执行 PRAGMA 查询以获取 WAL 模式状态Journal Mode:%s", journalMode)
	}

}

func versionedMigration(client *ent.Client) error {
	root, _ := os.Getwd()
	dir, err := migrate.NewLocalDir(path.Join(root, "../../migrations"))
	if err != nil {
		return err
	}

	var diff int
	opts := []schema.MigrateOption{
		schema.WithDir(dir),           // provide migration directory
		schema.WithForeignKeys(false), // disable foreign keys.
		schema.WithDropColumn(true),
		schema.WithDropIndex(true),
		schema.WithDiffHook(NewEntSchemaDiffHook(&diff)),
	}

	if err := client.Schema.Diff(context.Background(), opts...); err != nil {
		return err
	}
	if diff > 0 {
		log.Fatalf("当前schema与连接的数据库共有%v处差异，已将变更sql生成到migrations文件夹下，请执行后重试", diff)
	}
	return nil
}

func NewEntSchemaDiffHook(diff *int) schema.DiffHook {
	return func(next schema.Differ) schema.Differ {
		return schema.DiffFunc(func(current, desired *atlas.Schema) ([]atlas.Change, error) {
			changes, err := next.Diff(current, desired)
			if err != nil {
				return nil, err
			}
			for _, ch := range changes {
				switch t := ch.(type) {
				case *atlas.ModifyTable:
					*diff += len(t.Changes) // fix disable fk
				default:
					*diff += 1
				}
			}
			return changes, nil
		})
	}
}

func NewRedisClient(c *conf.Data) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		Password:     c.Redis.Password,
		DB:           int(c.Redis.Db),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
	})

	// Enable tracing instrumentation.
	if err := redisotel.InstrumentTracing(rdb); err != nil {
		panic(err)
	}

	// Enable metrics instrumentation.
	if err := redisotel.InstrumentMetrics(rdb); err != nil {
		panic(err)
	}

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}

	return rdb
}

func NewRedSync(conf *conf.Data, rdb *redis.Client) *redsync.Redsync {
	// Create a pool with go-redis which is the pool redisync will
	// use while communicating with Redis.
	pool := goredis.NewPool(rdb)
	// Create an instance of redisync to be used to obtain a mutual exclusion
	// lock.
	return redsync.New(pool)
}

func NewMqttClient(conf *conf.Data) mqtt.Client {
	// ias-lite 不通过配置文件连接 mqtt
	if conf.Mqtt == nil {
		return nil
	}

	opts := mqtt.NewClientOptions()
	opts.AddBroker(conf.Mqtt.Addr)
	opts.ClientID = shortuuid.New()
	opts.SetUsername(conf.Mqtt.Username)
	opts.SetPassword(conf.Mqtt.Password)
	opts.SetAutoReconnect(true)
	opts.SetCleanSession(false)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	mqttClient := mqtt.NewClient(opts)

	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		log.Error("mqtt无法连接，请确认网络连接")
	}
	return mqttClient
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	log.Info("mqtt connect")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	log.Infof("mqtt connected lost: %v", err)
}
