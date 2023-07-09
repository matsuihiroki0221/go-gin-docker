package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/matsuihiroki0221/gin-docker/db"
	"github.com/matsuihiroki0221/gin-docker/model"
)

type UserRepository struct{}

// User is alias of model.User struct
type User model.User

// GetAll is get all User
func (r UserRepository) GetAll() ([]User, error) {
	db := db.GetDB()

	var u []User
	if err := db.Find(&u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

func (s UserRepository) CreateModel(c *gin.Context) (User, error) {
	db := db.GetDB()
	var u User

	if err := c.BindJSON(&u); err != nil {
		return u, err
	}

	if err := db.Create(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}

func (r UserRepository) GetByID(id int) (User, error) {
	db := db.GetDB()
	var u User

	if err := db.First(&u, id).Error; err != nil {
		return u, err
	}

	return u, nil
}

func (r UserRepository) UpdateByID(id int, c *gin.Context) (User, error) {
	db := db.GetDB()
	var u User

	if err := db.Where("id = ?", id).First(&u).Error; err != nil {
		return u, err
	}

	if err := c.BindJSON(&u); err != nil {
		return u, err
	}

	db.Save(&u)

	return u, nil
}

func (r UserRepository) DeleteByID(id string) error {
	db := db.GetDB()
	var u User

	if err := db.Where("id = ?", id).Delete(&u).Error; err != nil {
		return err
	}
	return nil
}
