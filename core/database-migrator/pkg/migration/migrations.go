package migration

import (
	"github.com/pocockn/mono-repo/core/database-migrator/migrations"
	"gopkg.in/gormigrate.v1"
)

// Generate generates migrations for our db.
func Generate() []*gormigrate.Migration {
	var gormMigrations []*gormigrate.Migration

	allMigrations := get()
	for _, m := range allMigrations {
		gormMigrations = append(gormMigrations, m())
	}

	return gormMigrations
}

func get() []func() *gormigrate.Migration {
	return []func() *gormigrate.Migration{
		migrations.ShoutTableMigration,
		migrations.RecsTableMigration,
	}
}
