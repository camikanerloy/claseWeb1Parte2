package handlers

import (
	"net/http"
	"strconv"

	"github.com/camikanerloy/claseWeb1Parte2/internal/domain"
	"github.com/camikanerloy/claseWeb1Parte2/internal/product"
	"github.com/camikanerloy/claseWeb1Parte2/pkg/response"
	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

func GetPong(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}

func GetProducts(ctx *gin.Context) {
	//response
	ctx.JSON(http.StatusOK, response.Ok("Ok", product.Products))
}

func GetProductById(ctx *gin.Context) {
	//request
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Err(err))
		return
	}
	//process
	prod, err := product.GetByID(id)
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
	productsQuery := product.GetProductsQuery(priceQuery)

	// response
	ctx.JSON(http.StatusOK, response.Ok("Ok", productsQuery))
}

// Ejercitacion 2

func CreateProduct(ctx *gin.Context) {
	//request
	var request domain.Request
	if err := ctx.ShouldBind(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Err(err))
	}

	validator := validator.New()
	if err := validator.Struct(&request); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.Err(err))
		return
	}

	//process
	prod, err := product.Create(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Err(err))
		return
	}

	//response
	ctx.JSON(http.StatusCreated, response.Ok("suceed to create website", prod))
}
