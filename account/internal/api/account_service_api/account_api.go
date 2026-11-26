package account_service_api

import (
    "context"
    "time"
    "github.com/Teexijis/PetAdopt/account/internal/services/accountService"
    "github.com/Teexijis/PetAdopt/account/internal/pb/account_api"
)

type accountServiceInterface interface {
    Register(ctx context.Context, fullName, email, password string, phone *string) (int64, error)
    Login(ctx context.Context, email, password string) (*accountService.LoginResult, error)
    GetProfile(ctx context.Context, userID int64) (*accountService.UserProfile, error)
	GetAdminIDs(ctx context.Context) ([]int64, error)
}


type AccountServiceAPI struct {
    account_api.UnimplementedAccountServiceServer
    accountService accountServiceInterface
    jwtSecret []byte
    jwtTTL time.Duration
}

func NewAccountServiceAPI(service accountServiceInterface, jwtSecret []byte, jwtTTL time.Duration) *AccountServiceAPI {
    return &AccountServiceAPI{accountService: service, jwtSecret: jwtSecret, jwtTTL: jwtTTL}
}