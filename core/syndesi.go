package main

import (
	"github.com/goadesign/goa"
	"github.com/merakiVE/CEHDUN/core/app"
	"github.com/merakiVE/CEHDUN/core/types"
    "github.com/merakiVE/CEHDUN/common"
    "fmt"
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
	
	result := common.GetData(db)

	return ctx.OK([]byte(fmt.Sprintf("%v", result)))

}
