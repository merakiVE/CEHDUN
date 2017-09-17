package controllers

import (
	"github.com/kataras/iris"
)

func ConnectDB(ctx iris.Context) {

	var dataDB types.DataBase

	err = _context.ReadJSON(&dataDB)

	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(types.ResponseAPI{
			Message: "Invalid data",
			Data:    nil,
			Errors:  nil,
		})
		return
	}
	
	response := common.GetData(dataDB)

	result, _ := response.(types.BaseData)

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(types.ResponseAPI{
		Message: "Connection success",
		Data:    result,
		Errors:  nil,
	})
}
