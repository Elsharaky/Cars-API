package app

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Elsharaky/Cars-API.git/config"
	"github.com/Elsharaky/Cars-API.git/routes"
	"github.com/gorilla/mux"
)

func SetupAndRunAPI() error {
	if err := config.LoadENV(); err != nil {
		return err
	}

	if err := config.SetupDB(); err != nil {
		return err
	}

	if err := config.RunDBMigrations(); err != nil {
		return err
	}

	router := mux.NewRouter()
	routes.SetupTypeRoutes(router)
	routes.SetupCarsRoutes(router)

	fmt.Printf("Server running on http://%s:%s\n", os.Getenv("API_HOST"), os.Getenv("API_PORT"))

	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", os.Getenv("API_HOST"), os.Getenv("API_PORT")), router); err != nil {
		return err
	}

	return nil
}
