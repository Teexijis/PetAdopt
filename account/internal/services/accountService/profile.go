package accountService

import (
	"context"
	"errors"
)

type UserProfile struct {
	ID int64
	FullName string
	Email string
	Phone *string
	Role string
}


func (s *AccountService) GetProfile(ctx context.Context, userID int64) (*UserProfile, error) {
	user, err := s.userStorage.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("пользователь не найден")
	}

	return &UserProfile{
		ID: user.ID,
		FullName: user.FullName,
		Email: user.Email,
		Phone: user.Phone,
		Role: user.Role}, nil
}