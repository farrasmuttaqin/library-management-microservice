package usecase

import (
	"errors"
	"user_service/helpers"
	"user_service/models"
)

func (u *UserUseCase) validateRegisterForm(data models.RegisterForm) error {
	if helpers.IsEmptyStruct(data.Username) {
		return errors.New("username is required")
	}
	if helpers.IsEmptyStruct(data.Email) {
		return errors.New("email is required")
	}
	if helpers.IsEmptyStruct(data.Password) {
		return errors.New("password is required")
	}

	// Check if the author with the same email already exists
	qValidation := map[string]interface{}{
		"username": data.Username,
	}
	existingUserUserName, _ := u.UserRepo.Get(qValidation)
	if !helpers.IsEmptyStruct(existingUserUserName) {
		return errors.New("username already exist")
	}

	// Additional validation logic can be added here
	return nil
}
