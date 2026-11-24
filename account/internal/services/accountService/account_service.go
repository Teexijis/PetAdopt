package accountService

import (
    "context"
    "github.com/Teexijis/PetAdopt/account/internal/models"
)


type UserStorage interface {
    CreateUser(ctx context.Context, user *models.User) (int64, error)
    GetUserByEmail(ctx context.Context, email string) (*models.User, error)
    GetUserByID(ctx context.Context, id int64) (*models.User, error)
    GetAdminIDs(ctx context.Context) ([]int64, error)
}



type AccountService struct {
    userStorage UserStorage
	minNameLen int
    maxNameLen int
    minPasswordLen int
	maxPasswordLen int 
	maxEmailLen int 
}


func NewAccountService(userStorage UserStorage, minNameLen, maxNameLen, maxEmailLen, minPasswordLen, maxPasswordLen int) *AccountService {
    return &AccountService{
        userStorage: userStorage,
        minNameLen: minNameLen,
        maxNameLen: maxNameLen,
        minPasswordLen: minPasswordLen,
        maxPasswordLen: maxPasswordLen,
		maxEmailLen: maxEmailLen}
}