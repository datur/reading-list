package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReadingList struct {
	ID               uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Title            string    `json:"title"`
	UserID           uuid.UUID `json:"user" gorm:"type:uuid;foreign_key"`
	ReadingListItems []ReadingListItem
	IsDeleted        bool `json:"isDeleted"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type CreateReadingListInput struct {
	Title  string    `json:"title"`
	UserID uuid.UUID `json:"userID"`
}

type UpdateReadingListInput struct {
	Title string `json:"title"`
}

func (rL *ReadingList) BeforeCreate(tx *gorm.DB) (err error) {
	rL.ID = uuid.New()
	return
}
