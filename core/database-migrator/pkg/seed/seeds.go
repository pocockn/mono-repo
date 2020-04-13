package seed

import (
	"github.com/jinzhu/gorm"
	"github.com/pocockn/mono-repo/core/database-migrator/seeds"
)

// Process processes the seeds for the local DB.
func Process(db *gorm.DB) error {
	seeds := get()
	for _, s := range seeds {
		err := s(db)
		if err != nil {
			return err
		}
	}

	return nil
}

func get() []func(db *gorm.DB) error {
	return []func(db *gorm.DB) error{
		seeds.SeedShoutTable,
		seeds.SeedRecsTable,
	}
}
