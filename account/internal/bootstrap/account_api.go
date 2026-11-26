package bootstrap

import (
    "github.com/Teexijis/PetAdopt/account/config"
    "github.com/Teexijis/PetAdopt/account/internal/api/account_service_api"
    "github.com/Teexijis/PetAdopt/account/internal/services/accountService"
)

func InitAccountServiceAPI(accountService *accountService.AccountService, cfg *config.Config) *account_service_api.AccountServiceAPI {
    return account_service_api.NewAccountServiceAPI(accountService, []byte(cfg.JWT.Secret), cfg.JWT.TTL)
}