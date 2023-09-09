package console

import (
	"errors"
	"os"
	"os/signal"

	"github.com/fajarachmadyusup13/product-management/internal/config"
	"github.com/fajarachmadyusup13/product-management/internal/db"
	"github.com/fajarachmadyusup13/product-management/internal/delivery/httpsvc"
	"github.com/fajarachmadyusup13/product-management/internal/repository"
	"github.com/fajarachmadyusup13/product-management/internal/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "server",
	Short: "run server",
	Long:  "This subcommand start the server",
	Run:   run,
}

func init() {
	RootCmd.AddCommand(runCmd)
}

func run(cmd *cobra.Command, args []string) {
	db.InitializePostgreSQLConn()

	productRepo := repository.NewProductRepository(db.PostgreSQL)
	// productRepo := repository.NewProductRepository()
	productUsecase := usecase.NewProductUsecase(productRepo)

	httpServer := httpsvc.NewHTTPService()
	httpServer.SetProductUsecase(productUsecase)

	sigCh := make(chan os.Signal, 1)
	errCh := make(chan error, 1)
	signal.Notify(sigCh, os.Interrupt)
	go func() {
		<-sigCh
		errCh <- errors.New("received an interrupt")
		db.StopTickerCh <- true
	}()

	go func() {
		// Start HTTP server
		e := echo.New()
		e.Pre(middleware.AddTrailingSlash())
		e.Use(middleware.Recover())
		if config.Env() == config.EnvDevelopment {
			e.Use(middleware.Logger())
		}

		httpServer.InitRoutes(e)
		errCh <- e.Start(":" + config.HTTPPort())
	}()

	log.Error(<-errCh)
}
