package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Komsorn-sookdang/go-template/internal/app/inquirer_service/router"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func init() {

}

func StartServer(lifecycle fx.Lifecycle, r *gin.Engine) {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", 8080),
		Handler: r,
	}

	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					panic(err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("Shutting down server...")
			return server.Shutdown(ctx)
		},
	})
}

func main() {

	app := fx.New(
		/// repository.Module,
		/// service.Module,
		router.Module,
		fx.Provide(router.SetupBaseGinRouter),
		fx.Invoke(
			// route.SetupAuthRoute,
			StartServer,
		),
	)

	app.Run()
}
