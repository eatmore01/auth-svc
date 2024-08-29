package app

import (
	ra "auth/service/internal/app/rest"
	"auth/service/internal/config"
	"auth/service/internal/rest/auth"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

const (
	GRPC     = "grpc"
	REST     = "rest"
	TOGETHER = "together"
)

type App struct {
	Logger  *slog.Logger
	AppCfg  *config.AppConfig
	AuthApi *auth.AuthApi
}

func NewApp(logger *slog.Logger, cfg *config.AppConfig, AuthApi *auth.AuthApi) *App {
	return &App{
		Logger:  logger,
		AppCfg:  cfg,
		AuthApi: AuthApi,
	}
}

func (a *App) Run(flag string) {
	if flag == REST {
		restApp := ra.NewRestApp(a.Logger, a.AppCfg, a.AuthApi)
		restApp.Run()

		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

		<-quit

		a.Logger.Info("server start shutdown...")
		restApp.GraceFullShutDown()
		a.Logger.Info("server shutdown success")
	}
}
