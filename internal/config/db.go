package config

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"net/url"
	"time"
)

type DatabasePoolConf struct {
	MaxConn           int
	MinConn           int
	MaxConnLifetime   time.Duration
	KeepAliveInterval time.Duration
}

type DatabaseConfiguration struct {
	host     string
	port     int
	user     string
	password string
	database string
	DatabasePoolConf
}

func (dc *DatabaseConfiguration) dbConnectionUrl() string {
	u := url.URL{
		Scheme: "postgres",
		Host:   fmt.Sprintf("%s:%d", dc.host, dc.port),
		User:   url.UserPassword(dc.user, dc.password),
		Path:   dc.database,
	}

	values := url.Values{}
	values.Add("sslmode", "disable")
	u.RawQuery = values.Encode()

	return u.String()
}

func (dc *DatabaseConfiguration) GetDbConnection() *pgxpool.Pool {
	connectionUrl := dc.dbConnectionUrl()
	log.Printf("Connecting to database at %s", connectionUrl)
	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunc()

	dbPoolConfig, err := pgxpool.ParseConfig(connectionUrl)
	if err != nil {
		log.Fatalf("fail to connect to database: %v", err)
	}

	dbPoolConfig.MaxConns = int32(dc.MaxConn)
	dbPoolConfig.MinConns = int32(dc.MinConn)
	dbPoolConfig.MaxConnLifetime = dc.MaxConnLifetime
	dbPoolConfig.HealthCheckPeriod = dc.KeepAliveInterval

	poolDb, err := pgxpool.NewWithConfig(ctx, dbPoolConfig)
	if err != nil {
		log.Fatalf("fail to connect to database: %v", err)
	}

	return poolDb
}