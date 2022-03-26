package service

import (
	"errors"
	"reading-list/models"
	"time"

	"github.com/google/uuid"
)

func FindReadingList(id uuid.UUID) (models.ReadingList, error) {

	var readingList models.ReadingList

	err := models.DB.Where("id = ?", id).First(&readingList).Error

	if err == nil {
		models.DB.Preload("ReadingListItems").Find(&readingList)
	}

	if readingList.IsDeleted {
		err = errors.New("ReadingList is deleted")
	}
	return readingList, err
}

func CreateReadingList(input models.CreateReadingListInput) models.ReadingList {
	readingList := models.ReadingList{Title: input.Title, UserID: input.UserID}
	models.DB.Create(&readingList)
	return readingList
}

func UpdateReadingList(readingList models.ReadingList, input models.UpdateReadingListInput) (models.ReadingList, error) {

	models.DB.Model(readingList).Updates(input).Update("UpdatedAt", time.Now())

	return FindReadingList(readingList.ID)
}
