package accountService

import "gotest.tools/v3/assert"

func (s *AccountServiceSuite) TestGetAdminIDsSuccess() {
    expected := []int64{1, 3, 5}
    s.userStorage.EXPECT().GetAdminIDs(s.ctx).Return(expected, nil)
    ids, err := s.accountService.GetAdminIDs(s.ctx)

    assert.NilError(s.T(), err)
    assert.DeepEqual(s.T(), ids, expected)
}

func (s *AccountServiceSuite) TestGetAdminIDsEmpty() {
    s.userStorage.EXPECT().GetAdminIDs(s.ctx).Return([]int64{}, nil)
    ids, err := s.accountService.GetAdminIDs(s.ctx)

    assert.NilError(s.T(), err)
    assert.Equal(s.T(), len(ids), 0)
}