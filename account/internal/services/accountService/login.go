package accountService

import (
	"context"
	"errors"
	"strings"
	"golang.org/x/crypto/bcrypt"
)

type LoginResult struct {
	UserID int64
	Role string
}

func (s *AccountService) Login(ctx context.Context, email, password string) (*LoginResult, error) {
	email = strings.ToLower(strings.TrimSpace(email))

	user, err := s.userStorage.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("неверный email или пароль")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, errors.New("неверный email или пароль")
	}

	return &LoginResult{UserID: user.ID, Role: user.Role,}, nil
}
