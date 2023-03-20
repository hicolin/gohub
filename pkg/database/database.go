package database

import (
	"database/sql"
	"errors"
	"fmt"
	"gohub/pkg/config"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var SQLDB *sql.DB

func Connect(dbConfig gorm.Dialector, _logger logger.Interface) {
	var err error
	DB, err = gorm.Open(dbConfig, &gorm.Config{
		Logger: _logger,
	})
	if err != nil {
		fmt.Println(err.Error())
	}

	// 获取底层的 sqlDB
	SQLDB, err = DB.DB()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func CurrentDatabase() (dbname string) {
	dbname = DB.Migrator().CurrentDatabase()
	return
}

func DeleteAllTables() error {
	var err error
	switch config.Get("database.connection") {
	case "mysql":
		err = deleteMySQLTables()
	case "sqlite":
		err = deleteAllSqliteTables()
		if err != nil {
			return err
		}
	default:
		panic(errors.New("database connection not supported"))
	}
	return err
}

func deleteAllSqliteTables() error {
	tables := []string{}

	err := DB.Select(&tables, "select name FROM sqlite_master WHERE type = 'table'").Error
	if err != nil {
		return err
	}

	for _, table := range tables {
		err = DB.Migrator().DropTable(table)
		if err != nil {
			return err
		}
	}

	return nil
}

func deleteMySQLTables() error {
	dbname := CurrentDatabase()
	var tables []string

	err := DB.Table("information_schema.tables").
		Where("table_schema = ?", dbname).
		Pluck("table_name", &tables).
		Error
	if err != nil {
		return err
	}

	DB.Exec("SET foreign_key_checks = 0;")

	for _, table := range tables {
		err = DB.Migrator().DropTable(table)
		if err != nil {
			return err
		}
	}

	DB.Exec("SET foreign_key_checks = 1;")
	return nil
}
