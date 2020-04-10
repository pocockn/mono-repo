package shouts

import "github.com/jinzhu/gorm"

type (
	// Shout holds all the information about a shout.
	Shout struct {
		gorm.Model
		image string
	}
)
