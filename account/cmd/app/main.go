package main

import (
    "fmt"
    "github.com/Teexijis/PetAdopt/account/config"
    "github.com/Teexijis/PetAdopt/account/internal/bootstrap"
)

func main() {
    cfg, err := config.LoadConfig("config.yaml")
    if err != nil {
        panic(fmt.Sprintf("ошибка парсинга конфига: %v", err))
    }

    accountStorage := bootstrap.InitPGStorage(cfg)
    defer accountStorage.Close()
}