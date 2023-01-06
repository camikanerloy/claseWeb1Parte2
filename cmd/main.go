package main

import (
	"log"

	"github.com/camikanerloy/claseWeb1Parte2/cmd/handlers"
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
	//process
	err := product.GetProductsStruct()

	if err != nil {
		panic(err)
	}
	//server
	sv := gin.Default()

	//router
	svProducts := sv.Group("/products")
	svProducts.GET("/ping", handlers.GetPong)

	svProducts.GET("/", handlers.GetProducts)

	svProducts.GET("/:id", handlers.GetProductById)

	svProducts.GET("/search", handlers.GetProductQuery)

	//Ejercitacion 2
	svProducts.POST("/", handlers.CreateProduct)
	//start
	if err := sv.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
