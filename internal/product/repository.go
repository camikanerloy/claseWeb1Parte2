package product

import (
	"errors"

	"github.com/camikanerloy/claseWeb1Parte2/internal/domain"
	"github.com/camikanerloy/claseWeb1Parte2/internal/interfaces"
	"github.com/camikanerloy/claseWeb1Parte2/pkg/store"
)

type repository struct {
	interfaces.Repository
	st     store.Store
	cantId int
}

//var Products []domain.Product

func NewRepository(cantId int, st store.Store) interfaces.Repository {
	rta := &repository{
		st:     st,
		cantId: cantId,
	}

	return rta
}

func (pr repository) GetAll() ([]domain.Product, error) {
	products, err := pr.st.GetAll()
	if err != nil {
		return []domain.Product{}, err
	}
	return products, nil
}

func (pr repository) GetProdByID(id int) (product domain.Product, err error) {
	product, err = pr.st.GetOne(id)
	return
}

func (pr repository) GetProductsQuery(price float64) (productsQuery []domain.Product, err error) {
	products, err := pr.st.GetAll()
	if err != nil {
		return nil, err
	}
	for _, prod := range products {
		if prod.Price > price {
			productsQuery = append(productsQuery, prod)
		}
	}
	return
}

func (pr *repository) CreateProduct(newProducto domain.Product) (domain.Product, error) {
	if pr.ExistCodeValue(newProducto.CodeValue) {
		return domain.Product{}, errors.New("code value already exists")
	}
	pr.cantId++
	newProducto.Id = pr.cantId
	err := pr.st.AddOne(newProducto)
	if err != nil {
		return domain.Product{}, errors.New("error creating product")
	}
	return newProducto, nil
}

func (pr repository) ExistCodeValue(code string) bool {
	products, err := pr.st.GetAll()
	if err != nil {
		return false
	}
	for _, p := range products {
		if p.CodeValue == code {
			return true
		}
	}
	return false
}

func (r *repository) Update(id int, p domain.Product) (domain.Product, error) {
	prod, err := r.st.GetOne(id)

	if err != nil {
		return domain.Product{}, errors.New("product not found")
	}

	if r.ExistCodeValue(p.CodeValue) && prod.CodeValue != p.CodeValue {
		return domain.Product{}, errors.New("code value already exists")
	}
	err = r.st.UpdateOne(p)
	if err != nil {
		return domain.Product{}, errors.New("error updating product")
	}

	return p, nil
}

func (pr *repository) Delete(id int) error {
	err := pr.st.DeleteOne(id)
	if err != nil {
		return err
	}
	return nil
}
