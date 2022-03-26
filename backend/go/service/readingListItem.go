package service

import (
	"errors"
	"reading-list/models"
	"time"

	"github.com/google/uuid"
)

func CreateReadingListItem(input models.CreateReadingListItemInput) models.ReadingListItem {

	readingListItem := models.ReadingListItem{Title: input.Title, Author: input.Author, Url: input.Url, Tags: input.Tags, ReadingListID: input.ReadingListID}

	models.DB.Create(&readingListItem)

	return readingListItem
}

func FindReadingListItem(uuid uuid.UUID) (models.ReadingListItem, error) {
	var readingListItem models.ReadingListItem

	err := models.DB.Where("id = ?", uuid).First(&readingListItem).Error

	if err == nil {
		models.DB.Find(&readingListItem)
	}

	if readingListItem.IsDeleted {
		err = errors.New("User is deleted")
	}

	return readingListItem, err
}

func UpdateReadingListItem(readingListItem models.ReadingListItem, input models.UpdateReadingListItemInput) (models.ReadingListItem, error) {

	models.DB.Model(readingListItem).Updates(input).Update("UpdatedAt", time.Now())

	return FindReadingListItem(readingListItem.ID)
}

func SoftDeleteReadingListItem(readingListItem models.ReadingListItem) error {

	if readingListItem.IsDeleted {
		return errors.New("User already deleted")
	}

	r := models.DB.Model(readingListItem).Update("IsDeleted", true)
	return r.Error

}
