package service

import (
	"errors"
	"reading-list/models"
	"time"

	"github.com/google/uuid"
)

func CreateUser(input models.CreateUserInput) models.User {
	user := models.User{FirstName: input.FirstName, LastName: input.LastName, Email: input.Email}
	models.DB.Create(&user)
	return user
}

func FindUser(id uuid.UUID) (models.User, error) {

	var user models.User

	err := models.DB.Where("id = ?", id).First(&user).Error

	if err == nil {
		models.DB.Preload("ReadingLists").Preload("ReadingLists.ReadingListItems").Find(&user)
	}

	if user.IsDeleted {
		err = errors.New("User is deleted")
	}

	return user, err
}

func UpdateUser(user models.User, input models.UpdateUserInput) (models.User, error) {

	models.DB.Model(user).Updates(input).Update("UpdatedAt", time.Now())

	return FindUser(user.ID)
}

func SoftDeleteUser(user models.User) error {

	if user.IsDeleted {
		return errors.New("User already deleted")
	}

	r := models.DB.Model(user).Update("IsDeleted", true)
	return r.Error
}

func RestoreUser(user models.User) (models.User, error) {

	var err error

	if user.IsDeleted {
		err = models.DB.Model(user).Update("IsDeleted", false).Error
	}

	return user, err

}

func HardDeleteUser(user models.User) error {
	r := models.DB.Delete(user)
	return r.Error
}
