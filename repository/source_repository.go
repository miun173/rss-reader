package repository

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/miun173/rss-reader/model"
)

// SourceRepository :nodoc:
type SourceRepository struct {
	db *gorm.DB
}

// NewSourceRepository :nodoc:
func NewSourceRepository(db *gorm.DB) *SourceRepository {
	return &SourceRepository{db: db}
}

// FindAll :nodoc:
func (s *SourceRepository) FindAll(limit, offset int64) (sources []model.Source, err error) {
	err = s.db.
		Limit(limit).
		Offset(offset).
		Find(&sources).
		Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		log.Println("error: ", err)
		return sources, err
	}

	return sources, nil
}