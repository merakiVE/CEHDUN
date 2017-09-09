package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/merakiVE/CVDI/src/models"
	"github.com/merakiVE/CVDI/core/db"
	"github.com/merakiVE/CVDI/core/types"

	arangoDB "github.com/hostelix/aranGO"
)

type ProcedureController struct {
	mvc.Controller
	DB *arangoDB.Database
}

func (c *ProcedureController) Get() {
	result := make([]models.ProcedureModel, 0)
	var err error

	q := arangoDB.NewQuery(`FOR proc in procedures RETURN proc`)

	cur, err := c.DB.Execute(q)

	if err != nil {

		c.Ctx.StatusCode(iris.StatusInternalServerError)
		c.Ctx.JSON(types.ResponseAPI{
			Message: "Fail",
			Data:    nil,
			Errors:  nil,
		})
		return
	}

	err = cur.FetchBatch(&result)

	if err != nil {

		c.Ctx.StatusCode(iris.StatusInternalServerError)
		c.Ctx.JSON(types.ResponseAPI{
			Message: "Fail",
			Data:    nil,
			Errors:  nil,
		})
		return
	}

	c.Ctx.StatusCode(iris.StatusOK)
	c.Ctx.JSON(types.ResponseAPI{
		Message: "Success",
		Data:    result,
		Errors:  nil,
	})

}

func (c *ProcedureController) GetBy(key string) {

	var procedure models.ProcedureModel

	procedure.SetKey(key)

	success := db.GetModel(c.DB, &procedure)

	if !success {
		c.Ctx.StatusCode(iris.StatusNotFound)
		c.Ctx.JSON(types.ResponseAPI{
			Message: "Error get procedure, key not found",
			Data:    nil,
			Errors:  nil,
		})
		return
	}

	c.Ctx.StatusCode(iris.StatusOK)
	c.Ctx.JSON(types.ResponseAPI{
		Message: "Procedure " + key,
		Data:    procedure,
		Errors:  nil,
	})
}

func (c *ProcedureController) Post() {
	//diagram_xml, x, _:= c.Ctx.FormFile("diagram_xml")
}
