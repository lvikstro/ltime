//go:generate goagen bootstrap -d github.com/lvikstro/ltime/backend/design

package main

import (
	"os"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/lvikstro/ltime/backend/app"
	"github.com/lvikstro/ltime/backend/utils/database"
)

func main() {
	// Create service
	service := goa.New("LTime API")

	db, err := database.Connect(os.Getenv("PGUSER"), os.Getenv("PGPASS"), os.Getenv("PGDB"), os.Getenv("PGHOST"), os.Getenv("PGPORT"))
	if err != nil {
		service.LogError("startup", "err", err)
	}

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "authentication" controller
	c := NewAuthenticationController(service, db)
	app.MountAuthenticationController(service, c)
	// Mount "swagger" controller
	c2 := NewSwaggerController(service)
	app.MountSwaggerController(service, c2)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}
