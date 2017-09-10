//go:generate goagen bootstrap -d github.com/vpino/merakiVE/CEHDUN/core/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/merakiVE/CEHDUN/core/app"
)

func main() {
	// Create service
	service := goa.New("CEHDUN")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "syndesi" controller
	c := NewSyndesiController(service)
	app.MountSyndesiController(service, c)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}
