package seeds

import (
	"github.com/jinzhu/gorm"
)

type (
	Rec struct {
		gorm.Model
		Rating    int64
		Review    string
		SpotifyID string
		Title     string
	}
)

// SeedRecsTable seeds the recs table with test data.
func SeedRecsTable(db *gorm.DB) error {
	recSeeds := seedRecs()

	for _, rec := range recSeeds {
		err := db.Create(&rec).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func seedRecs() []Rec {
	var recsList []Rec

	recsList = append(
		recsList,
		Rec{
			Rating:    4,
			Review:    "Banging track",
			SpotifyID: "12345",
			Title:     "Tune",
		},
		Rec{
			Rating:    10,
			Review:    "Naughty track",
			SpotifyID: "123454",
			Title:     "Water",
		},
	)

	return recsList
}
