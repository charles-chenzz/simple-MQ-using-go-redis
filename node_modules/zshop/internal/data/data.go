package data

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"zshop/internal/conf"
	"zshop/internal/types"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewMysql, NewRedis, NewGreeterRepo, NewOrderRepo)

// Data .
type Data struct {
	db *sqlx.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, db *sqlx.DB) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{db: db}, cleanup, nil
}

func NewMysql(c *conf.Data) (*sqlx.DB, error) {
	// if you don't create volume here, when you running on k8s or docker it will panic

	db, err := sqlx.Connect("mysql", types.DSN)
	if err != nil {
		fmt.Printf("error:%v", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		fmt.Printf("error:%v", err)
	}

	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(50)

	return db, nil
}

func NewRedis() (*redis.Client, error) {
	// it should be read from remote if you intend to run in cloud and fill with the option that you need
	rds := redis.NewClient(
		types.RedisOptions(),
	)

	log.Debugf("redis connect:%v", rds)

	return rds, nil
}
