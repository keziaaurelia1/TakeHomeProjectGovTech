package config

import (
	"time"

	"github.com/keziaaurelia1/TakeHomeProjectGovTech/internal/common/envreader"
)

type MySQLCfg struct {
	DBName          string        `env:"MYSQL_DBNAME" required:"true" default:"dbname"`
	DBUser          string        `env:"MYSQL_DBUSER" required:"true" default:"dbuser"`
	DBPass          string        `env:"MYSQL_DBPASS" required:"true" default:"dbpass"`
	Host            string        `env:"MYSQL_HOST" required:"true" default:"localhost"`
	Port            string        `env:"MYSQL_PORT" required:"true" default:"3306"`
	MaxOpenConns    int           `env:"MYSQL_MAX_OPEN_CONNS" default:"30" required:"true"`
	MaxIdleConns    int           `env:"MYSQL_MAX_IDLE_CONNS" default:"6" required:"true"`
	ConnMaxLifetime time.Duration `env:"MYSQL_CONN_MAX_LIFETIME" default:"3m" required:"true"`
	MaxIdleTime     time.Duration `env:"MYSQL_MAX_IDLE_TIME" default:"0"`
}

func ProvideMySQLConfig() MySQLCfg {
	cfg := MySQLCfg{}
	envreader.BindEnv(&cfg)
	return cfg
}
