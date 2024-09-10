package main

// import (
// 	"assessment/internal/app/inquirer_service/controller/health"
// 	"assessment/internal/app/inquirer_service/route"
// 	"context"
// 	"fmt"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"go.uber.org/fx"
// )

// func init() {

// }

// func StartServer(lifecycle fx.Lifecycle, r *gin.Engine) {
// 	server := &http.Server{
// 		Addr:    fmt.Sprintf(":%d", "8080"),
// 		Handler: r,
// 	}

// 	lifecycle.Append(fx.Hook{
// 		OnStart: func(context.Context) error {
// 			go func() {
// 				if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
// 					panic(err)
// 				}
// 			}()
// 			return nil
// 		},
// 		OnStop: func(ctx context.Context) error {
// 			fmt.Println("Shutting down server...")
// 			return server.Shutdown(ctx)
// 		},
// 	})
// }

// func main() {

// 	app := fx.New(

// 		fx.Invoke(

// 			// route.SetupUserRoute,
// 			// route.SetupAuthRoute,
// 			StartServer,
// 		),
// 	)

// 	app.Run()
// }
