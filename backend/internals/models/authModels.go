package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/ra00d/event_management/internals/utils"
)

type LoginModel struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	RememberMe bool   `json:"remember_me"`
}

func (loginModel *LoginModel) Validate() error {
	return validation.ValidateStruct(
		loginModel,
		validation.Field(loginModel.Email, validation.Required.Error("email is required"),
			is.Email),
		validation.Field(
			loginModel.Password,
			validation.Required.Error("password is required"),
			validation.Length(6, 16),
		),
	)
}

type SignUpModel struct {
	Name                 string `json:"name"`
	Username             string `json:"username"`
	Phone                string `json:"phone"`
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

func (signUpModel *SignUpModel) Validate() error {
	// fmt.Println(signUpModel)
	return validation.ValidateStruct(
		signUpModel,
		validation.Field(&signUpModel.Email, validation.Required.Error("email is required"),
			is.Email),
		validation.Field(
			&signUpModel.Password,
			validation.Required.Error("password is required"),
			validation.Length(6, 16),
		),
		validation.Field(
			&signUpModel.PasswordConfirmation,
			validation.Required.Error("password confirmation is required"),
			validation.Length(6, 16),
			validation.By(utils.Confirmed(signUpModel.Password)),
			// validation.
		),
	)
}
