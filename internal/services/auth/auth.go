package auth

import (
	"auth/service/internal/config"
	"auth/service/internal/domain/model"
	"auth/service/internal/lib"
	"auth/service/internal/storage"
	"auth/service/internal/storage/user"
	"context"
	"errors"
	"log/slog"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	logger   *slog.Logger
	userRepo user.UserRepo
}

var (
	InvalidUserCred = errors.New("Invalid credentials")
)

func NewAuthService(logger *slog.Logger, ur user.UserRepo) *Auth {
	return &Auth{
		logger:   logger,
		userRepo: ur,
	}
}

type Storage interface {
	GetUser(ctx context.Context, email string) (model.User, error)
	CreateUser(ctx context.Context, email, user_name string, hashPass []byte) (string, error)
}

func (a *Auth) Login(ctx context.Context, email, password string, ac *config.AppConfig) (token string, loginErr error) {
	log := a.logger.With("function", "services/auth/Login")

	hashedPass, genErr := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if genErr != nil {
		log.Error(genErr.Error())
		return "", genErr
	}

	u, getErr := a.userRepo.GetUser(ctx, email)
	if getErr != nil {
		if errors.Is(getErr, storage.UserExist) {
			log.Error(storage.UserExist.Error())
			return "", storage.UserExist
		} else if errors.Is(getErr, storage.UserNotFound) {
			log.Error(storage.UserNotFound.Error())
			return "", InvalidUserCred
		} else {
			return "", getErr
		}
	}

	equalErr := bcrypt.CompareHashAndPassword(hashedPass, []byte(password))
	if equalErr != nil {
		log.Error(equalErr.Error())
		return "", InvalidUserCred
	}

	expTime, parseErr := time.ParseDuration(ac.ExpiriesTime)
	if parseErr != nil {
		log.Error(parseErr.Error())
		return "", parseErr
	}

	token, jwtErr := lib.NewJWTToken(u, ac.AppSecret, expTime)
	if jwtErr != nil {
		log.Error(jwtErr.Error())
		return "", jwtErr
	}

	return token, nil
}

func (a *Auth) Register(ctx context.Context, email, password, user_name string) (id string, registerErr error) {
	log := a.logger.With("function", "services/auth/Register")
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Error(err.Error())
		return "", err
	}

	if id, err = a.userRepo.CreateUser(ctx, email, user_name, hashedPass); err != nil {
		if errors.Is(err, storage.UserExist) {
			return "", storage.UserExist
		} else {
			return "", err
		}
	}

	return id, nil
}
