package account_service_api

import (
    "context"
    "github.com/Teexijis/PetAdopt/account/internal/pb/account_api"
)

func (a *AccountServiceAPI) GetProfile(ctx context.Context, req *account_api.GetProfileRequest) (*account_api.GetProfileResponse, error) {
    userID, _, err := a.GetUserFromToken(ctx)
    if err != nil {
        return nil, err
    }

    profile, err := a.accountService.GetProfile(ctx, userID)
    if err != nil {
        return nil, err
    }

    return &account_api.GetProfileResponse{
        Id: profile.ID,
        FullName: profile.FullName,
        Email: profile.Email,
        Phone: profile.Phone,
        Role: profile.Role}, nil
}