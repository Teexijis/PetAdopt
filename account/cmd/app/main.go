package main

import (
    "fmt"
	"os"
    "github.com/Teexijis/PetAdopt/account/config"
    "github.com/Teexijis/PetAdopt/account/internal/bootstrap"
)

func main() {
    cfg, err := config.LoadConfig(os.Getenv("configPath"))
    if err != nil {
        panic(fmt.Sprintf("ошибка парсинга конфига: %v", err))
    }

    accountStorage := bootstrap.InitPGStorage(cfg)
    defer accountStorage.Close()
	accountService := bootstrap.InitAccountService(accountStorage, cfg)
    accountAPI := bootstrap.InitAccountServiceAPI(accountService, cfg)
    bootstrap.AppRun(*accountAPI)
}