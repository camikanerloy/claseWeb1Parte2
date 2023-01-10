package product

import (
	"encoding/json"
	"errors"
	"io"
	"os"

	"github.com/camikanerloy/claseWeb1Parte2/internal/domain"
	"github.com/camikanerloy/claseWeb1Parte2/internal/interfaces"
)

type ProductRepository struct {
	interfaces.IProductRepository
	Products []domain.Product
	cantId   int
}

//var Products []domain.Product

func NewProductRepository() *ProductRepository {
	rta := &ProductRepository{}
	rta.GetProductsStruct()
	rta.cantId = 500
	return rta
}

func (pr *ProductRepository) GetProductsStruct() (err error) {

	jsonFile, err := os.Open("/Users/CKANER/bootcamp/claseWeb1Parte2/products.json")

	if err != nil {
		return
	}

	byteValue, err := io.ReadAll(jsonFile)

	if err != nil {
		return
	}

	if err = json.Unmarshal(byteValue, &pr.Products); err != nil {
		return
	}

	defer jsonFile.Close()
	return
}

func (pr ProductRepository) GetAll() ([]domain.Product, error) {
	return pr.Products, nil
}

func (pr ProductRepository) GetProdByID(id int) (product domain.Product, err error) {
	for _, prod := range pr.Products {
		if prod.Id == id {
			product = prod
			return
		}
	}
	return product, errors.New("no se encontro ningun producto con ese id")
}

func (pr ProductRepository) GetProductsQuery(price float64) (productsQuery []domain.Product, err error) {
	for _, prod := range pr.Products {
		if prod.Price > price {
			productsQuery = append(productsQuery, prod)
		}
	}
	return
}

func (pr *ProductRepository) CreateProduct(newProducto domain.Product) (domain.Product, error) {
	pr.cantId++
	newProducto.Id = pr.cantId
	pr.Products = append(pr.Products, newProducto)

	return newProducto, nil
}

func (pr ProductRepository) ExistCodeValue(code string) bool {
	for _, p := range pr.Products {
		if p.CodeValue == code {
			return true
		}
	}
	return false
}
