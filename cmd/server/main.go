package main

import (
	"log"

	"github.com/camikanerloy/claseWeb1Parte2/cmd/server/handlers"
	"github.com/camikanerloy/claseWeb1Parte2/internal/domain"
	"github.com/camikanerloy/claseWeb1Parte2/internal/product"

	"github.com/gin-gonic/gin"
)

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

type Response struct {
	Message string
	Data    []domain.Product
}

type ResponseByID struct {
	Message string
	Data    domain.Product
}

func main() {
	repo := product.NewProductRepository()
	productService := product.NewProductService(repo)
	productHandler := handlers.NewProductHandler(*productService)
	//server
	sv := gin.Default()

	//router
	sv.GET("/ping", productHandler.GetPong())
	svProducts := sv.Group("/products")
	{
		svProducts.GET("/", productHandler.GetProducts())

		svProducts.GET("/:id", productHandler.GetProductById())

		svProducts.GET("/search", productHandler.GetProductQuery())

		//Ejercitacion 2
		svProducts.POST("/", productHandler.CreateProduct())
	}
	//start

	if err := sv.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
