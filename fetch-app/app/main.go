package main

import (
	"context"
	"efishery/api"
	"efishery/app/modules"
	"efishery/config"
	"efishery/util"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"

	authBusiness "efishery/business/auth"
	authRepo "efishery/modules/repository/auth"
)

func main() {
	//load config if available or set to default
	config := config.GetConfig()

	//initialize database connection based on given config
	dbCon := util.NewDatabaseConnection(config)

	//initiate item repository
	controllers := modules.RegisterController(dbCon)

	//create echo http
	e := echo.New()

	// index route
	e.GET("/", func(c echo.Context) error {
		message := `Yang Fana adalah Waktu
Yang fana adalah waktu. Kita abadi:
Memungut detik demi detik,
Merangkainya seperti bunga
Sampai pada suatu hari
Kita lupa untuk apa.
“Tapi, yang fana adalah waktu, bukan?”
Tanyamu.
Kita abadi.
-- Sapardi Djoko Damono`
		return c.String(http.StatusOK, message)
	})

	authPermitRepo := authRepo.RepositoryFactory()
	authService := authBusiness.NewService(authPermitRepo)
	//register API path and handler
	api.RegisterPath(e, controllers, authService)

	// run server
	go func() {
		address := fmt.Sprintf("localhost:%d", config.Port)

		if err := e.Start(address); err != nil {
			log.Info("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	//close db
	defer dbCon.CloseConnection()

	// a timeout of 10 seconds to shutdown the server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
