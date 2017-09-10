package main

import (
	"github.com/goadesign/goa"
	"github.com/merakiVE/CEHDUN/core/app"
	"github.com/merakiVE/CEHDUN/core/types"
    "github.com/merakiVE/CEHDUN/common"
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
	
    db := types.DataBase{
    	Host: ctx.Payload.Host,
    	User: ctx.Payload.User, 
		Password: ctx.Payload.Password,
		Type: ctx.Payload.Type,
		Name: ctx.Payload.Name,
    }
	
	var message string

	con := common.Connect(db)

    if con == nil {
        message = "Error mi pana :)"
    }

	message = "Success :)"
	
	res := &app.SyndesiMedia{Message: message}
	return ctx.OK(res)

}
