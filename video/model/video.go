// gorm model for the video table in the database

package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Video struct {
	Id          string `gorm:"primaryKey"`
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	URL         string `gorm:"not null"`
	ChannelID   string `gorm:"not null"`
	Views       uint64
	Duration    int32
	Thumbnail   string
	PublishedAt time.Time
	Privacy     string
	Category    string
	Language    string
	Tags        []string
	isDeleted   bool
	Slug        string `gorm:"uniqueIndex"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

// Hook before create to generate uuid and slug
func (u *Video) BeforeCreate(tx *gorm.DB) error {
	uuid := uuid.NewV4()
	u.Id = uuid.String()
	u.Slug = uuid.String()
	return nil
}
