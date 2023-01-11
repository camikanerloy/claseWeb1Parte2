package interfaces

import "github.com/camikanerloy/claseWeb1Parte2/internal/domain"

type Service interface {
	GetProducts() (prod []domain.Product, err error)
	GetByID(id int) (product domain.Product, err error)
	GetProductsQuery(price float64) (productsQuery []domain.Product, err error)
	//ExistCodeValue(code string) bool
	ValidateExpiration(expiration string) (err error)
	Create(resq domain.Request) (prod domain.Product, err error)
	Update(id int, p domain.Product) (domain.Product, error)
}
