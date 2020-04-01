package internals

import (
	"github.com/pocockn/recs-api/models"
)

// Storer represents the database interactions.
type Storer interface {
	Create(rec *models.Rec) error
	FetchAll() (recs models.Recs, err error)
}
