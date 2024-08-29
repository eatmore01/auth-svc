package rest

import (
	"auth/service/internal/config"
	ar "auth/service/internal/rest/auth"
	"auth/service/pkg/lg"
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type RestApp struct {
	RestSrv *http.Server
	Logger  *slog.Logger
	AppCfg  *config.AppConfig
	AuthApi *ar.AuthApi
}

func NewRestApp(logger *slog.Logger, appCfg *config.AppConfig, AuthApi *ar.AuthApi) *RestApp {
	return &RestApp{
		RestSrv: &http.Server{},
		Logger:  logger,
		AppCfg:  appCfg,
		AuthApi: AuthApi,
	}
}

func (ra *RestApp) Run() {
	r := gin.Default()

	ar.AddAuthHandlers(r, ra.AuthApi)

	port := ":" + ra.AppCfg.SrvPort
	ra.RestSrv.Addr = port
	ra.RestSrv.Handler = r

	go func() {
		err := ra.RestSrv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			ra.Logger.Error("server error: %s", lg.Err(err))
			return

		}
	}()

}

func (ra *RestApp) GraceFullShutDown() {
	shutDownInterval := 5 * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), shutDownInterval)
	defer cancel()

	err := ra.RestSrv.Shutdown(ctx)
	if err != nil {
		ra.Logger.Error("server couldnt shutdown with error: %s", lg.Err(err))
	}

}
