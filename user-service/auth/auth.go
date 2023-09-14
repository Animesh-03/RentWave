package auth

import (
	"errors"
	"fmt"
	"time"
	"user/models"
	"user/repositories"

	"github.com/golang-jwt/jwt/v5"
)

func Register(userRepository *repositories.UserRepository, username string, password string) error {
	_, err := userRepository.GetByUsername(username)
	if err == nil {
		fmt.Println("Username already exists")
		return errors.New("Username already exists")
	}

	err = userRepository.AddEntity(&models.User{Username: username, PasswordHash: password})
	if err != nil {
		fmt.Println("Error registering user: ", err)
		return err
	}

	return nil
}

func Login(userRepository *repositories.UserRepository, username string, password string) (string, error) {
	user, err := userRepository.GetByUsername(username)
	if err != nil {
		fmt.Println("User not found")
		return "", errors.New("username is invalid")
	}
	if user.PasswordHash == password {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"uid":       user.ID,
			"username":  user.Username,
			"timestamp": time.Now().UnixMilli(),
		})
		secret := "secret"
		tokenString, err := token.SignedString([]byte(secret))
		if err != nil {
			fmt.Println("Could not generate JWT: ", err)
		}

		return tokenString, nil
	}
	return "", errors.New("password is invalid")
}
