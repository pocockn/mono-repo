package api

import (
	"github.com/jinzhu/gorm"
)

type (
	// Player holds data about a player on each team.
	Player struct {
		gorm.Model
		Name   string
		TeamID uint
	}

	// Players holds a list of player structs.
	Players []Player
)
