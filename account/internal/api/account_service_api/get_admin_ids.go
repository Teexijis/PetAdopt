package account_service_api

import (
    "context"
    "github.com/Teexijis/PetAdopt/account/internal/pb/account_api"
)

// вызывает сервис уведомлений
func (a *AccountServiceAPI) GetAdminIds(ctx context.Context, req *account_api.GetAdminIdsRequest) (*account_api.GetAdminIdsResponse, error) {
    adminIDs, err := a.accountService.GetAdminIDs(ctx)
    if err != nil {
        return nil, err
    }

    return &account_api.GetAdminIdsResponse{
        AdminIds: adminIDs}, nil
}