package pgstorage

import (
    "context"
    "fmt"
    "github.com/jackc/pgx/v5/pgxpool"
    "github.com/pkg/errors"
)

type PGstorage struct {
    db *pgxpool.Pool
}

const (
    tableName = "users"
    IDColumnName           = "id"
    FullNameColumnName     = "full_name"
    EmailColumnName        = "email"
    PhoneColumnName        = "phone"
    PasswordHashColumnName = "password_hash"
    RoleColumnName         = "role"
)

func NewPGStorage(connString string) (*PGstorage, error) {
    config, err := pgxpool.ParseConfig(connString)
    if err != nil {
        return nil, errors.Wrap(err, "ошибка парсинга конфига")
    }
    db, err := pgxpool.NewWithConfig(context.Background(), config)
    if err != nil {
        return nil, errors.Wrap(err, "ошибка создания пула соединений")
    }
    storage := &PGstorage{
        db: db,
    }

    if err := storage.initTables(); err != nil {
        db.Close()
        return nil, err
    }
    return storage, nil
}

func (s *PGstorage) initTables() error {
    sql := fmt.Sprintf(`
        CREATE TABLE IF NOT EXISTS %s (
            %s BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
            %s VARCHAR(150) NOT NULL,
            %s VARCHAR(255) UNIQUE NOT NULL,
            %s VARCHAR(20),
            %s TEXT NOT NULL,
            %s VARCHAR(20) NOT NULL DEFAULT 'user',
            CONSTRAINT users_role_check
                CHECK (%s IN ('user', 'admin'))
        )`, tableName, IDColumnName, FullNameColumnName, EmailColumnName, PhoneColumnName, PasswordHashColumnName, 
		RoleColumnName,RoleColumnName)
    _, err := s.db.Exec(context.Background(), sql)
    if err != nil {
        return errors.Wrap(err, "ошибка инициализации таблицы users")
    }
    return nil
}

func (s *PGstorage) Close() {
    s.db.Close()
}
