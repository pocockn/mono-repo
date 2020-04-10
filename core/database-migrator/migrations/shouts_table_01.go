package migrations

import (
	"github.com/jinzhu/gorm"
	"github.com/pocockn/models/api/shouts"
	"github.com/satori/go.uuid"
	"gopkg.in/gormigrate.v1"
)

// ShoutTableMigration holds the migration to create a shout table in the DB.
func ShoutTableMigration() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: uuid.Must(uuid.NewV4(),nil).String(),
		Migrate: func(tx *gorm.DB) error {
			type Shout struct {
				gorm.Model
				image string
			}
			return tx.AutoMigrate(&shouts.Shout{}).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.DropTable("shouts").Error
		},
	}
}
