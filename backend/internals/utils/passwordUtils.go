package utils

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CompareHash(pass string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	if err != nil {
		// fmt.Println(err.Error())
		return false
	} else {
		return true
	}
}

func Confirmed(pass string) validation.RuleFunc {
	return func(value interface{}) error {
		s, _ := value.(string)
		if s != pass {
			return errors.New("passwords does not match")
		}
		return nil
	}
}
