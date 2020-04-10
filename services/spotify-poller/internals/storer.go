package internals

import (
	models "github.com/pocockn/mono-repo/pkg/models/api/recs"
)

// Storer represents the database interactions.
type Storer interface {
	Create(rec *models.Rec) error
	FetchAll() (recs models.Recs, err error)
}
