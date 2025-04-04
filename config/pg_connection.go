package config

import (
	"context"
	"fmt"
	"time"

	"github.com/adwinugroho/wedding-management-system/internals/logger"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type EnvPostgreSQLConfig struct {
	PostgreSQLHost     string `mapstructure:"db_host"`
	PostgreSQLPort     string `mapstructure:"db_port"`
	PostgreSQLUser     string `mapstructure:"db_user"`
	PostgreSQLPassword string `mapstructure:"db_password"`
	PostgreSQLDBName   string `mapstructure:"db_name"`
}

var (
	PostgreSQLConfig EnvPostgreSQLConfig
)

type (
	PostgresDB struct {
		DB *pgxpool.Pool
	}
)

func InitConnectDB(ctx context.Context, dbHost, dbUser, dbPass, dbName string, dbPort int32) (*PostgresDB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", dbHost, dbUser, dbPass, dbName, dbPort)
	logger.LogWithFields(logrus.Fields{
		"info": "Connecting to database",
		"dsn":  dsn,
	}, "info connecting to database")

	poolConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		logger.LogFatal("Failed to parse config:" + err.Error())
		// return nil, err
	}

	poolConfig.MaxConns = 100
	poolConfig.MaxConnIdleTime = time.Minute * 1
	poolConfig.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeDescribeExec
	poolConfig.ConnConfig.RuntimeParams = map[string]string{}
	poolConfig.ConnConfig.RuntimeParams["application_name"] = "wedding_service"

	dbPool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		logger.LogFatal("Failed to create pool:" + err.Error())
		// return nil, err
	}

	err = dbPool.Ping(ctx)
	if err != nil {
		logger.LogFatal("Failed to ping:" + err.Error())
		// return nil, err
	}

	return &PostgresDB{DB: dbPool}, nil
}
