package database

import (
	"fmt"
	"loly_api/config"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func NewDatabase(conf *config.Config) *sqlx.DB {

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		conf.Db.User, conf.Db.Password, conf.Db.Host, conf.Db.Port, conf.Db.Name,
	)
	fmt.Printf("Connecting to: %s:[hidden]@tcp(%s:%s)/%s\n", conf.Db.User, conf.Db.Host, conf.Db.Port, conf.Db.Name)
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		panic(fmt.Errorf("failed to connect to database: %w", err))
	}

	err = db.Ping()
	if err != nil {
		panic(fmt.Errorf("failed to ping database: %w", err))
	}

	db.SetMaxOpenConns(conf.Db.MaxOpenConns)
	db.SetMaxIdleConns(conf.Db.MaxIdleConns)
	db.SetConnMaxLifetime(time.Duration(conf.Db.ConnMaxLifetime) * time.Second)

	return db
}
