package postgresql

import (
	"errors"
	"site/internal/config"
)

func CreateNewUser(user config.User) (uint, error) {

	result := DB.Omit("ID", "UpdatedAt").Create(&user)

	return user.ID, result.Error
}

func UserAuthorization(login, password string) (string, error) {

	user := config.User{}
	err := DB.Find(&user, "login = ?", login)
	if err.Error != nil {
		return "", err.Error
	}
	if login == user.Login && password == user.Password {
		return user.Name, nil
	}

	return "", errors.New("Невірний логін або пароль ")
}
