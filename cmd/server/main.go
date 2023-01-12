package main

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/camikanerloy/claseWeb1Parte2/cmd/server/routes"
	"github.com/camikanerloy/claseWeb1Parte2/internal/domain"
	"github.com/joho/godotenv"

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
	if err := godotenv.Load("/Users/CKANER/bootcamp/claseWeb1Parte2/cmd/server/.env"); err != nil {
		panic("Error loading .env file: " + err.Error())
	}

	//server
	sv := gin.Default()
	r := routes.NewRoute(sv)
	r.SetRoutes()
	//start

	if err := sv.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func GetProductsStruct() (data []domain.Product, err error) {

	jsonFile, err := os.Open("/Users/CKANER/bootcamp/claseWeb1Parte2/products.json")

	if err != nil {
		return
	}

	byteValue, err := io.ReadAll(jsonFile)

	if err != nil {
		return
	}

	if err = json.Unmarshal(byteValue, &data); err != nil {
		return
	}

	defer jsonFile.Close()
	return
}
