package postgresql

import (
	"site/internal/config"
)

func CreateNewUser() (uint, error) {
	user := config.User{
		Login:    "boer",
		Email:    "example@example.com",
		Password: "Qwerty1234",
		Role:     3,
		Name:     "Bohdan Yerokhin",
	}

	result := DB.Omit("ID", "UpdatedAt", "DeletedAt").Create(&user)

	return user.ID, result.Error
}
