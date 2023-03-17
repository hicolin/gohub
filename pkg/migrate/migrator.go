package migrate

import (
	"gohub/pkg/database"
	"gorm.io/gorm"
)

type Migrator struct {
	Folder   string
	DB       *gorm.DB
	Migrator gorm.Migrator
}

func (migrator *Migrator) createMigrationsTable() {
	migration := Migration{}

	if !migrator.Migrator.HasTable(&migration) {
		migrator.Migrator.CreateTable(&migration)
	}
}

type Migration struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement;"`
	Migration string `gorm:"type:varchar(255);not null;unique;"`
	Batch     int
}

func NewMigrator() *Migrator {
	migrator := &Migrator{
		Folder:   "database/migration/",
		DB:       database.DB,
		Migrator: database.DB.Migrator(),
	}
	migrator.createMigrationsTable()
	return migrator
}
