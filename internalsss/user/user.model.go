package user

import (
	"github.com/sas-scoul/bulk-api/database"
	"github.com/sas-scoul/bulk-api/utils"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string `json:"name"`
	TicketNumber string `json:"ticket_number"`
	Password     string `json:"password"`
}

func (u User) Create(user CreateUserDto) (string, error) {
	pwd, err := utils.HashPassword(user.Password)
	if err != nil {
		return "", err
	}

	user.Password = pwd
	result := database.DB.Create(&user)

	if result.Error != nil {
		return "", result.Error
	}

	return "Created succesfuly", nil
}

func (u User) All() ([]User, error) {
	var users []User
	result := database.DB.Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}
