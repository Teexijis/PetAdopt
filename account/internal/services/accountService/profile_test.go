package accountService

import (
	"github.com/Teexijis/PetAdopt/account/internal/models"
	"gotest.tools/v3/assert"
)

func (s *AccountServiceSuite) TestGetProfileSuccess() {
	phone := "+79966418175"
	user := &models.User{ID: 17, FullName: "Катя Сидорова", Email: "sidorova@mail.com",
		Phone: &phone, PasswordHash: "486465465845646546", Role: "user"}
	s.userStorage.EXPECT().GetUserByID(s.ctx, int64(17)).Return(user, nil)
	profile, err := s.accountService.GetProfile(s.ctx, 17)

	assert.NilError(s.T(), err)
	assert.Assert(s.T(), profile != nil)
	assert.Equal(s.T(), profile.ID, int64(17))
	assert.Equal(s.T(), profile.FullName, "Катя Сидорова")
	assert.Equal(s.T(), profile.Email, "sidorova@mail.com")
	assert.Equal(s.T(), *profile.Phone, phone)
	assert.Equal(s.T(), profile.Role, "user")
}

func (s *AccountServiceSuite) TestGetProfileNotFound() {
	s.userStorage.EXPECT().GetUserByID(s.ctx, int64(999)).Return(nil, nil)
	profile, err := s.accountService.GetProfile(s.ctx, 999)
	assert.Error(s.T(), err, "пользователь не найден")
	assert.Assert(s.T(), profile == nil)
}
