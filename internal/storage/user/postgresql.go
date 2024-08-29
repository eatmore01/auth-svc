package user

import (
	"auth/service/internal/domain/model"
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepo struct {
	Client *pgxpool.Pool
}

func NewUserRepo(client *pgxpool.Pool) *UserRepo {
	return &UserRepo{
		Client: client,
	}
}

func (ur *UserRepo) GetUser(ctx context.Context, email string) (model.User, error) {
	var u model.User

	q := `--sql
	SELECT id, user_name, email, password FROM users WHERE email = $1
	`
	if err := ur.Client.QueryRow(ctx, q, email).Scan(&u.Id, &u.UserName, &u.Email, &u.Passhash); err != nil {
		return model.User{}, err
	}

	return u, nil
}

func (ur *UserRepo) CreateUser(ctx context.Context, email, user_name string, hashPass []byte) (string, error) {
	q := `--sql
	INSERT INTO users 
	(email, user_name, password) 
	VALUES
	($1, $2, $3)
	RETURNING id
	`
	var newId uuid.UUID
	if err := ur.Client.QueryRow(ctx, q, email, user_name, hashPass).Scan(&newId); err != nil {
		var pgErr *pgconn.PgError
		if errors.Is(err, &pgconn.PgError{}) {
			pgErr = err.(*pgconn.PgError)
			return "", fmt.Errorf("failed to create user, code: %s, mesage: %s, details: %s", pgErr.Code, pgErr.Message, pgErr.Detail)
		}
		return "", err
	}

	return newId.String(), nil
}
