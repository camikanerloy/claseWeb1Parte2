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

type ProductHandler struct {
	// Interfaz?
	ProductService product.ProductService
}

func NewProductHandler(service product.ProductService) *ProductHandler {
	return &ProductHandler{
		ProductService: service,
	}
}

func (ph ProductHandler) GetPong() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	}
}

func (ph ProductHandler) GetProducts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//response
		prod, err := ph.ProductService.GetProducts()

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
		}

		ctx.JSON(http.StatusOK, response.Ok("Ok", prod))
	}
}

func (ph ProductHandler) GetProductById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//request
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, response.Err(err))
			return
		}
		//process
		prod, err := ph.ProductService.GetByID(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, response.Err(err))
			return
		}

		//response
		ctx.JSON(http.StatusOK, response.Ok("Ok", prod))
	}
}

func (ph ProductHandler) GetProductQuery() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//request
		priceQuery, err := strconv.ParseFloat(ctx.Query("priceGt"), 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, response.Err(err))
			return
		}
		//process
		productsQuery, err := ph.ProductService.GetProductsQuery(priceQuery)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
		}

		// response
		ctx.JSON(http.StatusOK, response.Ok("Ok", productsQuery))
	}
}

// Ejercitacion 2

func (ph ProductHandler) CreateProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
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
		prod, err := ph.ProductService.Create(request)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, response.Err(err))
			return
		}

		//response
		ctx.JSON(http.StatusCreated, response.Ok("suceed to create website", prod))
	}
}
