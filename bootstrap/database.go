package bootstrap

import (
	"errors"
	"fmt"
	"gohub/pkg/logger"
	"time"

	"gohub/pkg/config"
	"gohub/pkg/database"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDB() {
	var dbConfig gorm.Dialector
	switch config.Get("database.connection") {
	case "mysql":
		dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&multiStatements=true&loc=Local",
			config.Get("database.mysql.username"),
			config.Get("database.mysql.password"),
			config.Get("database.mysql.host"),
			config.Get("database.mysql.port"),
			config.Get("database.mysql.database"),
			config.Get("database.mysql.charset"),
		)
		dbConfig = mysql.New(mysql.Config{
			DSN: dsn,
		})
	//case "sqlite":
	//dsn := config.Get("database.sqlite.database")
	//dbConfig = sqlite.Open(dsn)
	default:
		panic(errors.New("database connection not supported"))
	}

	database.Connect(dbConfig, logger.NewGormLogger())

	database.SQLDB.SetMaxOpenConns(config.GetInt("database.mysql.max_open_connections"))
	database.SQLDB.SetMaxIdleConns(config.GetInt("database.mysql.max_idle_connections"))
	database.SQLDB.SetConnMaxIdleTime(time.Duration(config.GetInt("database.mysql.max_life_seconds")) * time.Second)
}
