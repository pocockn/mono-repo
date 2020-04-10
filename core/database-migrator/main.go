package main

import (
	"github.com/pocockn/mono-repo/core/database-migrator/seeds"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pocockn/mono-repo/core/database-migrator/config"
	"github.com/pocockn/mono-repo/pkg/models/api/shouts"
	"github.com/sirupsen/logrus"
	"gopkg.in/gormigrate.v1"
)

func main() {
	logrus.Print("Starting migrations...\n")

	appConfig := config.NewConfig()

	var DB *gorm.DB
	var err error

	for i := 0; i <= 30; i++ {
		DB, err = gorm.Open("mysql", appConfig.Database.URL)
		if err == nil {
			err := DB.DB().Ping()
			if err == nil {
				DB.LogMode(true)
				break
			}
		}

		if i == 15 {
			log.Fatalf("Unable to connect to %s after 30 seconds", appConfig.Database.URL)
		}

		logrus.Infof("%d attempt at connecting to the DB \n", i)
		time.Sleep(2 * time.Second)
	}

	gormMigrator := gormigrate.New(DB, gormigrate.DefaultOptions, GenerateMigrations())

	gormMigrator.InitSchema(func(tx *gorm.DB) error {
		logrus.Print("Creating initial table schema...")
		err := tx.AutoMigrate(
			&shouts.Shout{},
			&seeds.Rec{},
		)

		if err != nil {
			log.Printf("Err : %+v /n creating initial schema", err)
		}

		return nil
	})

	if err := gormMigrator.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
	logrus.Print("Initial schema migration successfull")

	err = processSeeds(DB)
	if err != nil {
		log.Fatalf("Error will seeding database: %v", err)
	}
	logrus.Print("Database seeding successfull")
}
