package interfaces

import (
	"github.com/camikanerloy/claseWeb1Parte2/internal/domain"
)

type Repository interface {
	GetAll() ([]domain.Product, error)
	GetProdByID(int) (domain.Product, error)
	GetProductsQuery(float64) ([]domain.Product, error)
	CreateProduct(newProducto domain.Product) (domain.Product, error)
	ExistCodeValue(code string) bool
	Update(id int, p domain.Product) (domain.Product, error)
}
