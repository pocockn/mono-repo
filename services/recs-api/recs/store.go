package recs

import (
	"github.com/pocockn/mono-repo/pkg/models/api/recs"
)

// Store represents the database interactions.
type Store interface {
	Fetch(id uint) (rec recs.Rec, err error)
	FetchAll() (recs recs.Recs, err error)
	Update(rec *recs.Rec) error
}
