package api

import (
	"github.com/jinzhu/gorm"
)

type (

	// Team holds information about a team.
	Team struct {
		gorm.Model
		Name string
	}

	// Teams holds a list of team structs.
	Teams []Team
)
