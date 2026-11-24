package accountService

import (
	"context"
	"errors"
	"testing"
	"github.com/Teexijis/PetAdopt/account/internal/services/accountService/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gotest.tools/v3/assert"
	"strings"
	"github.com/Teexijis/PetAdopt/account/internal/models"
	"golang.org/x/crypto/bcrypt"
)


type AccountServiceSuite struct {
	suite.Suite
	ctx            context.Context
	userStorage    *mocks.UserStorage
	accountService *AccountService
}


func (s *AccountServiceSuite) SetupTest() {
	s.userStorage = mocks.NewUserStorage(s.T())
	s.ctx = context.Background()
	s.accountService = NewAccountService(s.userStorage, 5, 150, 254, 8, 72)
}

func (s *AccountServiceSuite) TestRegisterSuccess() {
    s.userStorage.EXPECT().CreateUser(s.ctx, mock.Anything).Return(int64(1), nil)
    userID, err := s.accountService.Register(s.ctx, "Катя Сидорова", "sidorova@mail.ru", "852963741", nil)
    assert.NilError(s.T(), err)
    assert.Equal(s.T(), userID, int64(1))
}

func (s *AccountServiceSuite) TestRegisterStorageError() {
	wantErr := errors.New("error")

	s.userStorage.EXPECT().CreateUser(s.ctx, mock.Anything).Return(int64(0), wantErr)

	userID, err := s.accountService.Register(s.ctx, "Катя Сидорова", "sidorova@mail.ru", "852963741", nil)
	assert.ErrorIs(s.T(), err, wantErr)
	assert.Equal(s.T(), userID, int64(0))
}

func (s *AccountServiceSuite) TestRegisterEmptyNameError() {
	userID, err := s.accountService.Register(s.ctx, "", "sidorova@mail.ru", "852963741", nil)
	assert.Check(s.T(), err != nil)
	assert.Equal(s.T(), userID, int64(0))
}

func (s *AccountServiceSuite) TestRegisterInvalidEmailError() {
	userID, err := s.accountService.Register(s.ctx, "Катя Сидорова", "email-email", "852963741", nil)
	assert.Check(s.T(), err != nil)
	assert.Equal(s.T(), userID, int64(0))
}


func (s *AccountServiceSuite) TestRegisterShortNameError() {
    userID, err := s.accountService.Register(s.ctx, "Катя", "sidorova@mail.ru", "852963741", nil)
    assert.Check(s.T(), err != nil)
    assert.Equal(s.T(), userID, int64(0))
}

func (s *AccountServiceSuite) TestRegisterLongNameError() {
    fullName := strings.Repeat("А", 151)
    userID, err := s.accountService.Register(s.ctx, fullName, "sidorova@mail.ru", "852963741", nil)
    assert.Check(s.T(), err != nil)
    assert.Equal(s.T(), userID, int64(0))
}

func (s *AccountServiceSuite) TestRegisterShortPasswordError() {
    userID, err := s.accountService.Register(s.ctx, "Катя Сидорова", "sidorova@mail.ru", "1234567", nil)
    assert.Check(s.T(), err != nil)
    assert.Equal(s.T(), userID, int64(0))
}

func (s *AccountServiceSuite) TestRegisterLongPasswordError() {
    password := strings.Repeat("и", 37) // тут граничный тест 37 символов 74 байта 
    userID, err := s.accountService.Register(s.ctx, "Катя Сидорова", "sidorova@mail.ru", password, nil)
    assert.Check(s.T(), err != nil)
    assert.Equal(s.T(), userID, int64(0))
}


func (s *AccountServiceSuite) TestRegisterInvalidPhoneError() {
    phone := "7989qwe123"
    userID, err := s.accountService.Register(s.ctx, "Катя Сидорова", "sidorova@mail.ru", "852963741", &phone)

    assert.Check(s.T(), err != nil)
    assert.Equal(s.T(), userID, int64(0))
}


func (s *AccountServiceSuite) TestRegisterLongPhoneError() {
    phone := strings.Repeat("9", 16)
    userID, err := s.accountService.Register(s.ctx, "Катя Сидорова", "sidorova@mail.ru", "852963741", &phone)
    assert.Check(s.T(), err != nil)
    assert.Equal(s.T(), userID, int64(0))
}


func (s *AccountServiceSuite) TestRegisterPasswordHash() {
    password := "852963741"
    checkUser := func(ctx context.Context, user *models.User) {
        err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
        assert.NilError(s.T(), err)
    }

    s.userStorage.EXPECT().CreateUser(s.ctx, mock.Anything).Run(checkUser).Return(int64(1), nil)
    userID, err := s.accountService.Register(s.ctx, "Катя Сидорова", "sidorova@mail.ru", password, nil)

    assert.NilError(s.T(), err)
    assert.Equal(s.T(), userID, int64(1))
}


func TestAccountServiceSuite(t *testing.T) {
	suite.Run(t, new(AccountServiceSuite))
}
