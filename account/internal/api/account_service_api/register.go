package account_service_api

import (
    "context"
    "github.com/Teexijis/PetAdopt/account/internal/pb/account_api"
)

func (a *AccountServiceAPI) Register(ctx context.Context, req *account_api.RegisterRequest) (*account_api.RegisterResponse, error) {
    userID, err := a.accountService.Register(ctx, req.GetFullName(), req.GetEmail(), req.GetPassword(), req.Phone)
    if err != nil {
        return nil, err
    }
    return &account_api.RegisterResponse{UserId: userID}, nil
}