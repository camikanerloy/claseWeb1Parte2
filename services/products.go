package services

import (
	"encoding/json"
	"errors"
	"io"
	"os"

	"github.com/camikanerloy/claseWeb1Parte2/services/models"
)

var Products []models.Product

func OpenJsonFile() (jsonFile *os.File, err error) {
	jsonFile, err = os.Open("products.json")

	if err != nil {
		return
	}
	return
}

func GetProductsStruct() (err error) {

	jsonFile, err := OpenJsonFile()

	if err != nil {
		return
	}

	byteValue, err := io.ReadAll(jsonFile)

	if err != nil {
		return
	}

	if err = json.Unmarshal(byteValue, &Products); err != nil {
		return
	}

	defer jsonFile.Close()
	return
}

func GetByID(id int) (product models.Product, err error) {
	for _, prod := range Products {
		if prod.Id == id {
			product = prod
			return
		}
	}
	return product, errors.New("no se encontro ningun producto con ese id")
}

func GetProductsQuery(price float64) (productsQuery []models.Product) {

	for _, prod := range Products {
		if prod.Price > price {
			productsQuery = append(productsQuery, prod)
		}
	}
	return
}
