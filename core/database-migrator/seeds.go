package main

import (
	"github.com/jinzhu/gorm"
	"github.com/pocockn/database-migrator/seeds"
)

func processSeeds(db *gorm.DB) error {
	seeds := getSeeds()

	for _, seed := range seeds {
		err := seed(db)
		if err != nil {
			return err
		}
	}

	return nil
}

func getSeeds() []func(db *gorm.DB) error {
	return []func(db *gorm.DB) error{
		seeds.SeedShoutTable,
		seeds.SeedRecsTable,
	}
}
