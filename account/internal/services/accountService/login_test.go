package accountService

import (
	"errors"
	"github.com/Teexijis/PetAdopt/account/internal/models"
	"golang.org/x/crypto/bcrypt"
	"gotest.tools/v3/assert"
)


func (s *AccountServiceSuite) TestLoginSuccess() {
	hash, err := bcrypt.GenerateFromPassword([]byte("852963741"), bcrypt.DefaultCost)
	assert.NilError(s.T(), err)

	user := &models.User{
		ID: 15,
		Email: "sidorova@mail.ru",
		PasswordHash: string(hash),
		Role: "admin"}

	s.userStorage.EXPECT().GetUserByEmail(s.ctx, "sidorova@mail.ru").Return(user, nil)

	result, err := s.accountService.Login(s.ctx, "sidorova@mail.ru", "852963741")

	assert.NilError(s.T(), err)
	assert.Assert(s.T(), result != nil)
	assert.Equal(s.T(), result.UserID, int64(15))
	assert.Equal(s.T(), result.Role, "admin")
}


func (s *AccountServiceSuite) TestLoginWrongPassword() {
	hash, err := bcrypt.GenerateFromPassword([]byte("852963741"), bcrypt.DefaultCost)
	assert.NilError(s.T(), err)

	user := &models.User{
		ID: 15,
		Email: "sidorova@mail.ru",
		PasswordHash: string(hash),
		Role: "admin"}

	s.userStorage.EXPECT().GetUserByEmail(s.ctx, "sidorova@mail.ru").Return(user, nil)
	result, err := s.accountService.Login(s.ctx, "sidorova@mail.ru", "efesffsfcwda")

	assert.Error(s.T(), err, "неверный email или пароль")
	assert.Assert(s.T(), result == nil)
}

func (s *AccountServiceSuite) TestLoginUserNotFound() {
	s.userStorage.EXPECT().GetUserByEmail(s.ctx, "mail@mail.ru").Return(nil, nil)

	result, err := s.accountService.Login(s.ctx, "mail@mail.ru", "852963741")

	assert.Error(s.T(), err, "неверный email или пароль")
	assert.Assert(s.T(), result == nil)
}


func (s *AccountServiceSuite) TestLoginStorageError() {
	wantErr := errors.New("database error")
	s.userStorage.EXPECT().GetUserByEmail(s.ctx, "sidorova@mail.ru").Return(nil, wantErr)

	result, err := s.accountService.Login(s.ctx, "sidorova@mail.ru", "852963741")

	assert.ErrorIs(s.T(), err, wantErr)
	assert.Assert(s.T(), result == nil)
}

func (s *AccountServiceSuite) TestLoginNormalizeEmail() {
	hash, err := bcrypt.GenerateFromPassword([]byte("852963741"), bcrypt.DefaultCost)
	assert.NilError(s.T(), err)

	user := &models.User{
		ID: 15,
		Email: "sidorova@mail.ru",
		PasswordHash: string(hash),
		Role: "admin"}

	s.userStorage.EXPECT().GetUserByEmail(s.ctx, "sidorova@mail.ru").Return(user, nil)

	result, err := s.accountService.Login(s.ctx,"   SIdOROVA@MAIL.RU  ", "852963741")

	assert.NilError(s.T(), err)
	assert.Assert(s.T(), result != nil)
	assert.Equal(s.T(), result.UserID, int64(15))
}


func (s *AccountServiceSuite) TestLoginDifferentRole() {
	hash, err := bcrypt.GenerateFromPassword([]byte("852963741"), bcrypt.DefaultCost)
	assert.NilError(s.T(), err)

	user := &models.User{
		ID: 20,
		Email: "user@mail.ru",
		PasswordHash: string(hash),
		Role: "user"}

	s.userStorage.EXPECT().GetUserByEmail(s.ctx, "user@mail.ru").Return(user, nil)
	result, err := s.accountService.Login(s.ctx, "user@mail.ru", "852963741")

	assert.NilError(s.T(), err)
	assert.Assert(s.T(), result != nil)
	assert.Equal(s.T(), result.UserID, int64(20))
	assert.Equal(s.T(), result.Role, "user")
}