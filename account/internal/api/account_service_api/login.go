package account_service_api

import (
    "context"
    "github.com/Teexijis/PetAdopt/account/internal/pb/account_api"
    "github.com/Teexijis/PetAdopt/account/internal/services/accountService"
)

func (a *AccountServiceAPI) Login(ctx context.Context, req *account_api.LoginRequest) (*account_api.LoginResponse, error) {
    result, err := a.accountService.Login(ctx, req.GetEmail(), req.GetPassword())
    if err != nil {
        return nil, err
    }
    token, err := accountService.GenerateJWT(result.UserID, result.Role, a.jwtSecret, a.jwtTTL)
    if err != nil {
        return nil, err
    }

    return &account_api.LoginResponse{AccessToken: token}, nil
}