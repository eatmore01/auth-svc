package main

import (
	"auth/service/internal/app"
	"auth/service/internal/config"
	"auth/service/internal/rest/auth"
	as "auth/service/internal/services/auth"
	"auth/service/internal/storage/user"
	"auth/service/pkg/client/postgresql"
	"auth/service/pkg/lg"
	"context"
	"fmt"

	"github.com/spf13/viper"
)

const (
	GRPC     = "grpc"
	REST     = "rest"
	TOGETHER = "together"
)

func init() {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

func main() {
	cfg := config.MustLoad()
	expTime := viper.GetString("EXPIRED_TIME")
	appSecret := viper.GetString("APP_SECRET")
	serverPort := viper.GetString("SERVER_PORT")

	logger := lg.SetupLogger(cfg.Env)

	sc := postgresql.StorageConfig{
		PgUser: cfg.DB.User,
		PgDB:   cfg.DB.Name,
		PgPass: cfg.DB.Password,
		PgPort: cfg.DB.Port,
	}

	ac := config.AppConfig{
		ExpiriesTime: expTime,
		AppSecret:    appSecret,
		SrvPort:      serverPort,
	}

	PostgreSqlClient, err := postgresql.NewClient(context.TODO(), sc)
	if err != nil {
		panic(fmt.Sprintf("error with connect to db: %s", err))
	}

	userRepo := user.NewUserRepo(PostgreSqlClient)
	userService := as.NewAuthService(logger, *userRepo)
	authAPi := auth.NewAuthApi(&ac, userService)

	app := app.NewApp(logger, &ac, authAPi)
	app.Run(REST)
}
