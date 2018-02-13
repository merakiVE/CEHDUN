package controllers

import (
	"os"
	"io"

	"github.com/kataras/iris"
	"github.com/merakiVE/CVDI/src/models"
	"github.com/merakiVE/koinos/db"
	"github.com/merakiVE/CVDI/core/types"
	arangoDB "github.com/hostelix/aranGO"
)

type ProcedureController struct {
	Ctx iris.Context
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

func (c *ProcedureController) PostUpload() {

	file, info, err := c.Ctx.FormFile("diagram_xml")

	if err != nil {
		c.Ctx.StatusCode(iris.StatusInternalServerError)
		c.Ctx.JSON(types.ResponseAPI{
			Message: "Error upload file",
			Errors:  err.Error(),
		})
		return
	}

	defer file.Close()

	fname := info.Filename

	out, err := os.OpenFile("/tmp/"+fname, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		c.Ctx.StatusCode(iris.StatusInternalServerError)
		c.Ctx.JSON(types.ResponseAPI{
			Message: "Error upload file",
			Errors:  err.Error(),
		})
		return
	}

	defer out.Close()

	io.Copy(out, file)

	c.Ctx.StatusCode(iris.StatusOK)
	c.Ctx.JSON(types.ResponseAPI{
		Message: "File uploaded success",
		Data:    fname,
	})
}
