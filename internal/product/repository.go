package product

import (
	"errors"

	"github.com/camikanerloy/claseWeb1Parte2/internal/domain"
	"github.com/camikanerloy/claseWeb1Parte2/internal/interfaces"
)

type repository struct {
	interfaces.Repository
	products []domain.Product
	cantId   int
}

//var Products []domain.Product

func NewRepository(data []domain.Product, cantId int) interfaces.Repository {
	rta := &repository{
		products: data,
		cantId:   cantId,
	}

	return rta
}

func (pr repository) GetAll() ([]domain.Product, error) {
	return pr.products, nil
}

func (pr repository) GetProdByID(id int) (product domain.Product, err error) {
	for _, prod := range pr.products {
		if prod.Id == id {
			product = prod
			return
		}
	}
	return product, errors.New("no se encontro ningun producto con ese id")
}

func (pr repository) GetProductsQuery(price float64) (productsQuery []domain.Product, err error) {
	for _, prod := range pr.products {
		if prod.Price > price {
			productsQuery = append(productsQuery, prod)
		}
	}
	return
}

func (pr *repository) CreateProduct(newProducto domain.Product) (domain.Product, error) {
	pr.cantId++
	newProducto.Id = pr.cantId
	pr.products = append(pr.products, newProducto)

	return newProducto, nil
}

func (pr repository) ExistCodeValue(code string) bool {
	for _, p := range pr.products {
		if p.CodeValue == code {
			return true
		}
	}
	return false
}
