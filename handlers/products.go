package handlers

import (
	"net/http"
	"strconv"

	"github.com/camikanerloy/claseWeb1Parte2/pkg/response"
	"github.com/camikanerloy/claseWeb1Parte2/services"
	"github.com/gin-gonic/gin"
)

func GetPong(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}

func GetProducts(ctx *gin.Context) {
	//response
	ctx.JSON(http.StatusOK, response.Ok("Ok", services.Products))
}

func GetProductById(ctx *gin.Context) {
	//request
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Err(err))
		return
	}
	//process
	prod, err := services.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Err(err))
		return
	}

	//response
	ctx.JSON(http.StatusOK, response.Ok("Ok", prod))
}

func GetProductQuery(ctx *gin.Context) {
	//request
	priceQuery, err := strconv.ParseFloat(ctx.Query("priceGt"), 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Err(err))
		return
	}
	//process
	productsQuery := services.GetProductsQuery(priceQuery)

	// response
	ctx.JSON(http.StatusOK, response.Ok("Ok", productsQuery))
}
