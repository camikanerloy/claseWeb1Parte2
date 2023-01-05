package main

import (
	"github.com/camikanerloy/claseWeb1Parte2/handlers"
	"github.com/camikanerloy/claseWeb1Parte2/services"
	"github.com/camikanerloy/claseWeb1Parte2/services/models"
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
	Data    []models.Product
}

type ResponseByID struct {
	Message string
	Data    models.Product
}

func main() {
	//process
	err := services.GetProductsStruct()

	if err != nil {
		panic(err)
	}
	router := gin.Default()

	router.GET("/ping", handlers.GetPong)

	router.GET("/products", handlers.GetProducts)

	router.GET("products/:id", handlers.GetProductById)

	router.GET("/products/search", handlers.GetProductQuery)

	router.Run()
}
