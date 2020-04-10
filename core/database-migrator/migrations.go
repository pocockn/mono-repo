package main

import (
	"github.com/pocockn/mono-repo/core/database-migrator/migrations"
	"gopkg.in/gormigrate.v1"
)

// GenerateMigrations generates migrations for our db.
func GenerateMigrations() []*gormigrate.Migration {
	var gormMigrations []*gormigrate.Migration

	allMigrations := getMigrations()

	for _, migration := range allMigrations {
		gormMigrations = append(gormMigrations, migration())
	}

	return gormMigrations
}

func getMigrations() []func() *gormigrate.Migration {
	return []func() *gormigrate.Migration{
		migrations.ShoutTableMigration,
		migrations.RecsTableMigration,
	}
}
