package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"

	"github.com/merakiVE/CEHDUN/common"
	"github.com/merakiVE/CEHDUN/core/types"

	"fmt"
)

func ConnectDB(ctx context.Context) {

	var dataDB types.DataBase

	err := ctx.ReadJSON(&dataDB)

	fmt.Println(dataDB)

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

	result, errr := response.(types.BaseData)

	fmt.Println(response)

	if !errr {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(types.ResponseAPI{
			Message: "Invalid data",
			Data:    nil,
			Errors:  nil,
		})
		return
	} 

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(types.ResponseAPI{
		Message: "Connection success",
		Data:    result,
		Errors:  nil,
	})
}
