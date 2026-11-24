package accountService

import (
	"context"
	"strings"
	"github.com/Teexijis/PetAdopt/account/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func (s *AccountService) Register(ctx context.Context, fullName, email, password string, phone *string) (int64, error) {
	fullName = strings.TrimSpace(fullName)
	email = strings.ToLower(strings.TrimSpace(email))
	err := s.validateInfo(fullName, email, password, phone)
	if err != nil {
    	return 0, err
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	if phone != nil && *phone == "" { // телефон не обязательный 
		phone = nil
	}

	user := &models.User{
		FullName: fullName,
		Email: email,
		Phone: phone,
		PasswordHash: string(passwordHash),
		Role: "user"}

	userID, err := s.userStorage.CreateUser(ctx, user)
	if err != nil {
		return 0, err
	}

	return userID, nil
}