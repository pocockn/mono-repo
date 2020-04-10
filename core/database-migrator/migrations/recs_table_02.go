package migrations

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"gopkg.in/gormigrate.v1"
)

// RecsTableMigration holds the migration to create a rec table in the DB.
func RecsTableMigration() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: uuid.Must(uuid.NewV4(), nil).String(),
		Migrate: func(tx *gorm.DB) error {
			type Rec struct {
				gorm.Model
				Rating    int64
				Review    string
				SpotifyID string
				Title     string
			}
			return tx.AutoMigrate(&Rec{}).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.DropTable("recs").Error
		},
	}
}
