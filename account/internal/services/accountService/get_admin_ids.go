package accountService

import "context"

func (s *AccountService) GetAdminIDs(ctx context.Context) ([]int64, error) {
    return s.userStorage.GetAdminIDs(ctx)
}