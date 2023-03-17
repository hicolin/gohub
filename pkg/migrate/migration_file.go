package migrate

import (
	"database/sql"
	"gorm.io/gorm"
)

type migrationFunc func(migrator gorm.Migrator, db sql.DB)

type MigrationFile struct {
	Up       migrationFunc
	down     migrationFunc
	FileName string
}

var migrationFiles []MigrationFile

func Add(name string, up migrationFunc, down migrationFunc) {
	migrationFiles = append(migrationFiles, MigrationFile{
		FileName: name,
		Up:       up,
		down:     down,
	})
}
