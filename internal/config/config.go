package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	v *viper.Viper
}

type Server struct {
	Port     string
	GrpcPort string
}

func NewConfig(v *viper.Viper) *Config {
	return &Config{
		v: v,
	}
}

func (c *Config) NewServer() *Server {
	return &Server{
		Port:     c.v.GetString("service.port"),
		GrpcPort: c.v.GetString("service.grpc-port"),
	}
}

func (c *Config) NewDatabaseConfig(configName string) *DatabaseConfiguration {
	return &DatabaseConfiguration{
		host:     c.v.GetString(fmt.Sprintf("database.%s.host", configName)),
		port:     c.v.GetInt(fmt.Sprintf("database.%s.port", configName)),
		user:     c.v.GetString(fmt.Sprintf("database.%s.user", configName)),
		password: c.v.GetString(fmt.Sprintf("database.%s.password", configName)),
		database: c.v.GetString(fmt.Sprintf("database.%s.name", configName)),
		DatabasePoolConf: DatabasePoolConf{
			MaxConn:           c.v.GetInt(fmt.Sprintf("database.%s.max-conn", configName)),
			MinConn:           c.v.GetInt(fmt.Sprintf("database.%s.min-conn", configName)),
			KeepAliveInterval: c.v.GetDuration(fmt.Sprintf("database.%s.keep-alive-interval", configName)),
			MaxConnLifetime:   c.v.GetDuration(fmt.Sprintf("database.%s.max-life-time", configName)),
		},
	}
}
