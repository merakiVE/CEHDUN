package main

import (
	"github.com/goadesign/goa"
	"github.com/vpino/merakiVE/CEHDUN/core/app"
)

// SyndesiController implements the syndesi resource.
type SyndesiController struct {
	*goa.Controller
}

// NewSyndesiController creates a syndesi controller.
func NewSyndesiController(service *goa.Service) *SyndesiController {
	return &SyndesiController{Controller: service.NewController("SyndesiController")}
}

// Connect runs the connect action.
func (c *SyndesiController) Connect(ctx *app.ConnectSyndesiContext) error {
	
	user := ctx.Payload.Name

	var message string

	if user == "Victor" {
		message = "Success :)"
	} else {
		message = "Error mi pana :("
	}
	
	res := &app.SyndesiMedia{Message: message}
	return ctx.OK(res)

}
