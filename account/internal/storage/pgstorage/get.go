package pgstorage

import (
    "context"
    "github.com/Teexijis/PetAdopt/account/internal/models"
    "github.com/Masterminds/squirrel"
    "github.com/pkg/errors"
)


func (storage *PGstorage) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
    query := storage.getUserByEmailQuery(email)
    queryText, args, err := query.ToSql()
    if err != nil {
        return nil, errors.Wrap(err, "ошибка формирования запроса")
    }

    var user models.User
    err = storage.db.QueryRow(ctx, queryText, args...).Scan(&user.ID, &user.FullName, &user.Email, &user.Phone,
		&user.PasswordHash, &user.Role)
    if err != nil {
        return nil, errors.Wrap(err, "ошибка получения пользователя по email")
    }

    return &user, nil
}

func (storage *PGstorage) getUserByEmailQuery(email string) squirrel.Sqlizer {
    q := squirrel.Select(IDColumnName, FullNameColumnName, EmailColumnName, PhoneColumnName, PasswordHashColumnName,
		RoleColumnName).From(tableName).Where(squirrel.Eq{EmailColumnName: email}).PlaceholderFormat(squirrel.Dollar)
	return q
}


func (storage *PGstorage) GetUserByID(ctx context.Context, id int64) (*models.User, error) {
    query := storage.getUserByIDQuery(id)
    queryText, args, err := query.ToSql()
    if err != nil {
        return nil, errors.Wrap(err, "ошибка формирования запроса")
    }

    var user models.User
    err = storage.db.QueryRow(ctx, queryText, args...).Scan(&user.ID, &user.FullName, &user.Email, &user.Phone, 
		&user.PasswordHash, &user.Role)
    if err != nil {
        return nil, errors.Wrap(err, "ошибка получения пользователя по ID")
    }

    return &user, nil
}


func (storage *PGstorage) getUserByIDQuery(id int64) squirrel.Sqlizer {
    q := squirrel.Select(IDColumnName, FullNameColumnName, EmailColumnName, PhoneColumnName, PasswordHashColumnName,
		RoleColumnName).From(tableName).Where(squirrel.Eq{IDColumnName: id}).PlaceholderFormat(squirrel.Dollar)
    return q
}


func (storage *PGstorage) GetAdminIDs(ctx context.Context) ([]int64, error) {
    query := storage.getAdminIDsQuery()
    queryText, args, err := query.ToSql()
    if err != nil {
        return nil, errors.Wrap(err, "ошибка формирования запроса")
    }

    rows, err := storage.db.Query(ctx, queryText, args...)
    if err != nil {
        return nil, errors.Wrap(err, "ошибка получения администраторов")
    }
	defer rows.Close()
    var adminIDs []int64
    for rows.Next() {
        var id int64
        if err := rows.Scan(&id); err != nil {
            return nil, errors.Wrap(err, "ошибка чтения ID администратора")
        } 
        adminIDs = append(adminIDs, id)
    }

    return adminIDs, nil
}


func (storage *PGstorage) getAdminIDsQuery() squirrel.Sqlizer {
    q := squirrel.Select(IDColumnName).From(tableName).Where(squirrel.Eq{RoleColumnName: "admin"}).PlaceholderFormat(squirrel.Dollar)
	
	return q
}
