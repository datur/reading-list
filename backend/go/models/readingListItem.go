package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReadingListItem struct {
	ID            uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	ReadingListID uuid.UUID `json:"readingListId" gorm:"type:uuid;foreign_key"`
	Title         string    `json:"title"`
	Author        string    `json:"author"`
	Url           string    `json:"url"`
	Tags          string    `json:"tags"`
	IsDeleted     bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type CreateReadingListItemInput struct {
	Title         string    `json:"title"`
	Author        string    `json:"author"`
	Url           string    `json:"url"`
	Tags          string    `json:"tags"`
	ReadingListID uuid.UUID `json:"readingListID"`
}

type UpdateReadingListItemInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Url    string `json:"url"`
	Tags   string `json:"tags"`
}

func (rLI *ReadingListItem) BeforeCreate(tx *gorm.DB) (err error) {
	rLI.ID = uuid.New()
	return
}
