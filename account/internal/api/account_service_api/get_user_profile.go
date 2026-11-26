package account_service_api

import (
    "context"
    "github.com/Teexijis/PetAdopt/account/internal/pb/account_api"
)

// вызывается adoption сервисом 
func (a *AccountServiceAPI) GetUserProfile(ctx context.Context, req *account_api.GetUserProfileRequest) (*account_api.GetUserProfileResponse, error) {
    profile, err := a.accountService.GetProfile(ctx, req.GetUserId())
    if err != nil {
        return nil, err
    }

    return &account_api.GetUserProfileResponse{
        FullName: profile.FullName,
        Email: profile.Email,
        Phone: profile.Phone}, nil
}