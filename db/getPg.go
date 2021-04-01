package db

import (
	"time"

	configurations "movieserver/config"

	"github.com/go-pg/pg"
	"go.uber.org/zap"
)

func getDbSession() (*pg.DB, error) {
	cfg := configurations.InitConfigs()

	pgIns, err := cfg.GetPostgresIni("pg")
	if err != nil {
		zap.L().Info("getting the postgres config failed in postgres ", zap.Error(err))
		return nil, err
	}

	options := &pg.Options{
		Addr:                  pgIns.PgConfig.Host,
		User:                  pgIns.PgConfig.Username,
		Password:              pgIns.PgConfig.Password,
		Database:              pgIns.PgConfig.Database,
		PoolSize:              200,
		IdleTimeout:           time.Minute * 2,
		DialTimeout:           time.Second * 2,
		PoolTimeout:           time.Second * 2,
		RetryStatementTimeout: false,
		MaxConnAge:            time.Minute * 2,
	}
	return pg.Connect(options), nil
}
