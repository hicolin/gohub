package migrate

import (
	"database/sql"
	"gorm.io/gorm"
)

type migrationFunc func(migrator gorm.Migrator, db sql.DB)

type MigrationFile struct {
	Up       migrationFunc
	Down     migrationFunc
	FileName string
}

var migrationFiles []MigrationFile

func Add(name string, up migrationFunc, down migrationFunc) {
	migrationFiles = append(migrationFiles, MigrationFile{
		FileName: name,
		Up:       up,
		Down:     down,
	})
}

func getMigrationFile(name string) MigrationFile {
	for _, mfile := range migrationFiles {
		if name == mfile.FileName {
			return mfile
		}
	}
	return MigrationFile{}
}

func (mfile MigrationFile) isNotMigrated(migrations []Migration) bool {
	for _, migration := range migrations {
		if migration.Migration == mfile.FileName {
			return false
		}
	}
	return true
}
