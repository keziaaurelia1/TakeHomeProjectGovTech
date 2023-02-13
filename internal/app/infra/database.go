package infra

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/keziaaurelia1/TakeHomeProjectGovTech/internal/config"
	"github.com/sirupsen/logrus"
)

// ProvideMySQL ...
func ProvideMySQL() *sql.DB {
	cfg := config.ProvideMySQLConfig()
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?tls=false&parseTime=true", cfg.DBUser, cfg.DBPass, cfg.Host, cfg.Port, cfg.DBName))
	if err != nil {
		log.Fatalf("mysql: %s", err.Error())
		return nil
	}

	if db == nil {
		log.Fatalf("Got nil db")
		return nil
	}

	db.SetConnMaxLifetime(cfg.ConnMaxLifetime)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetConnMaxIdleTime(cfg.MaxIdleTime)
	if err = db.Ping(); err != nil {
		logrus.Fatalf("mysql: %s", err.Error())
	}
	return db
}
