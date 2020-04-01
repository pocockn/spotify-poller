package store

import (
	"github.com/jinzhu/gorm"
	"github.com/pocockn/recs-api/models"
	"github.com/pocockn/spotify-poller/internals"
)

type store struct {
	Conn *gorm.DB
}

// NewStore creates a new store struct for interacting with the Gorm connection to the DB.
func NewStore(conn *gorm.DB) internals.Storer {
	return &store{conn}
}

// Create adds a new record to the DB.
func (s *store) Create(rec *models.Rec) error {
	return s.Conn.Create(rec).Error
}

// FetchAll fetches all the recs from the DB.
func (s *store) FetchAll() (models.Recs, error) {
	var recs models.Recs
	err := s.Conn.Find(&recs).Error

	return recs, err
}
