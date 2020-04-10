package seeds

import (
	"github.com/jinzhu/gorm"
	"github.com/pocockn/models/api/shouts"
)

// SeedShoutTable seeds the team table with test data.
func SeedShoutTable(db *gorm.DB) error {
	shoutSeeds := seedShouts()

	for _, shout := range shoutSeeds {
		err := db.Create(&shout).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func seedShouts() []shouts.Shout {
	var shoutsList []shouts.Shout

	shoutsList = append(
		shoutsList,
		shouts.Shout{
			Image: "https://adobe99u.files.wordpress.com/2018/01/latoya-dixon-stock-photography-main.jpg?quality=100&resize=1240,920&strip",
		},
		shouts.Shout{
			Image: "https://adobe99u.files.wordpress.com/2018/01/mark-marziars-stock-photography.jpg?quality=100",
		},
	)

	return shoutsList
}
