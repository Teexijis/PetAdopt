package pgstorage

import (
    "context"
    "github.com/Teexijis/PetAdopt/account/internal/models"
    "github.com/Masterminds/squirrel"
    "github.com/pkg/errors"
)


func (storage *PGstorage) CreateUser(ctx context.Context, user *models.User) (int64, error) {
    query := storage.createUserQuery(user)
    queryText, args, err := query.ToSql()
    if err != nil {
        return 0, errors.Wrap(err, "ошибка формирования запроса")
    }

    var userID int64
    err = storage.db.QueryRow(ctx, queryText, args...).Scan(&userID)
    if err != nil {
        return 0, errors.Wrap(err, "ошибка создания пользователя")
    }
    return userID, nil
}


func (storage *PGstorage) createUserQuery(user *models.User) squirrel.Sqlizer {
    return squirrel.Insert(tableName).Columns(FullNameColumnName, EmailColumnName, PhoneColumnName, 
		PasswordHashColumnName, RoleColumnName).Values(user.FullName, user.Email, user.Phone, 
			user.PasswordHash, user.Role).Suffix("RETURNING id").PlaceholderFormat(squirrel.Dollar)
}