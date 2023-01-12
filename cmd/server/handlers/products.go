package handlers

import (
	"errors"
	"net/http"
	"os"
	"strconv"

	"github.com/camikanerloy/claseWeb1Parte2/internal/domain"
	"github.com/camikanerloy/claseWeb1Parte2/internal/interfaces"
	"github.com/camikanerloy/claseWeb1Parte2/pkg/response"
	"github.com/camikanerloy/claseWeb1Parte2/pkg/web"
	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	// Interfaz?
	ProductService interfaces.Service
}

func NewProductHandler(service interfaces.Service) *ProductHandler {
	return &ProductHandler{
		ProductService: service,
	}
}

func (ph ProductHandler) GetPong() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		web.Success(ctx, http.StatusOK, "pong")
	}
}

func (ph ProductHandler) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token == "" {
			web.Failure(ctx, 401, errors.New("token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(ctx, http.StatusUnauthorized, errors.New("error: invalid token"))
			return
		}

		//response
		prod, err := ph.ProductService.GetProducts()

		if err != nil {
			web.Failure(ctx, http.StatusInternalServerError, err)
		}

		web.Success(ctx, http.StatusOK, prod)
		//ctx.JSON(http.StatusOK, response.Ok("Ok", prod))
	}
}

func (ph ProductHandler) GetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token == "" {
			web.Failure(ctx, 401, errors.New("token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(ctx, http.StatusUnauthorized, errors.New("error: invalid token"))
			return
		}

		//request
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, err)
			return
		}
		//process
		prod, err := ph.ProductService.GetByID(id)
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, err)
			return
		}

		//response
		ctx.JSON(http.StatusOK, response.Ok("Ok", prod))
	}
}

func (ph ProductHandler) Search() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token == "" {
			web.Failure(ctx, 401, errors.New("token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(ctx, http.StatusUnauthorized, errors.New("error: invalid token"))
			return
		}

		//request
		priceQuery, err := strconv.ParseFloat(ctx.Query("priceGt"), 64)
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, err)
			return
		}
		//process
		productsQuery, err := ph.ProductService.GetProductsQuery(priceQuery)

		if err != nil {
			web.Failure(ctx, http.StatusInternalServerError, err)
		}

		// response
		web.Success(ctx, http.StatusOK, productsQuery)
		//ctx.JSON(http.StatusOK, response.Ok("Ok", productsQuery))
	}
}

// Ejercitacion 2

func (ph ProductHandler) Post() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token == "" {
			web.Failure(ctx, 401, errors.New("token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(ctx, http.StatusUnauthorized, errors.New("error: invalid token"))
			return
		}

		//request
		var request domain.Request
		if err := ctx.ShouldBind(&request); err != nil {
			web.Failure(ctx, http.StatusBadRequest, err)
		}

		validator := validator.New()
		if err := validator.Struct(&request); err != nil {
			web.Failure(ctx, http.StatusUnprocessableEntity, err)
			return
		}

		//process
		prod, err := ph.ProductService.Create(request)
		if err != nil {
			web.Failure(ctx, http.StatusInternalServerError, err)
			return
		}

		//response
		web.Success(ctx, http.StatusCreated, prod)
		//ctx.JSON(http.StatusCreated, response.Ok("suceed to create website", prod))
	}
}

// Put actualiza un producto
func (h *ProductHandler) Put() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token == "" {
			web.Failure(ctx, 401, errors.New("token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(ctx, http.StatusUnauthorized, errors.New("error: invalid token"))
			return
		}

		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(ctx, 400, errors.New("error: invalid id"))
			return
		}
		var product domain.Product
		err = ctx.ShouldBindJSON(&product)
		if err != nil {
			web.Failure(ctx, 400, errors.New("error: invalid product"))
			return
		}
		p, err := h.ProductService.Update(id, product)
		if err != nil {
			web.Failure(ctx, 409, err)
			return
		}
		ctx.JSON(200, p)
	}
}

// Patch update selected fields of a product WIP
func (h *ProductHandler) Patch() gin.HandlerFunc {
	type Request struct {
		Name        string  `json:"name,omitempty"`
		Quantity    int     `json:"quantity,omitempty"`
		CodeValue   string  `json:"code_value,omitempty"`
		IsPublished bool    `json:"is_published,omitempty"`
		Expiration  string  `json:"expiration,omitempty"`
		Price       float64 `json:"price,omitempty"`
	}
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token == "" {
			web.Failure(ctx, 401, errors.New("token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(ctx, http.StatusUnauthorized, errors.New("error: invalid token"))
			return
		}

		var r Request
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(ctx, 400, errors.New("error: invalid id"))
			return
		}
		if err := ctx.ShouldBindJSON(&r); err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New("error: invalid request"))
			//			ctx.JSON(400, gin.H{"error": "invalid request"})
			return
		}
		update := domain.Product{
			Name:        r.Name,
			Quantity:    r.Quantity,
			CodeValue:   r.CodeValue,
			IsPublished: r.IsPublished,
			Expiration:  r.Expiration,
			Price:       r.Price,
		}

		p, err := h.ProductService.Update(id, update)
		if err != nil {
			web.Failure(ctx, 409, err)
			return
		}
		ctx.JSON(200, p)
	}
}

func (ph *ProductHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//Sacar el token del context
		token := ctx.GetHeader("TOKEN")
		if token == "" {
			web.Failure(ctx, 401, errors.New("token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(ctx, http.StatusUnauthorized, errors.New("error: invalid token"))
			return
		}

		//recuperar el id
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, err)
			return
		}

		err = ph.ProductService.Delete(id)
		if err != nil {
			web.Failure(ctx, 404, err)
			return
		}

		web.Success(ctx, 204, nil)
	}
}
