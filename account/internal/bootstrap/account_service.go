package bootstrap

import (
    "github.com/Teexijis/PetAdopt/account/config"
    "github.com/Teexijis/PetAdopt/account/internal/services/accountService"
    "github.com/Teexijis/PetAdopt/account/internal/storage/pgstorage"
)

func InitAccountService(storage *pgstorage.PGstorage, cfg *config.Config) *accountService.AccountService {
    return accountService.NewAccountService(storage,
        cfg.AccountServiceSettings.MinNameLen,
        cfg.AccountServiceSettings.MaxNameLen,
        cfg.AccountServiceSettings.MaxEmailLen,
        cfg.AccountServiceSettings.MinPasswordLen,
        cfg.AccountServiceSettings.MaxPasswordLen)
}