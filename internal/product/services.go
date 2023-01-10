package product

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/camikanerloy/claseWeb1Parte2/internal/domain"
	"github.com/camikanerloy/claseWeb1Parte2/internal/interfaces"
)

type ProductService struct {
	interfaces.IProductService
	RepoProd interfaces.IProductRepository
}

var cantId int = 500
var (
	ErrAlreadyExist = errors.New("error: item already exist")
)

func NewProductService(repo interfaces.IProductRepository) *ProductService {
	return &ProductService{
		RepoProd: repo,
	}
}

func (provServ ProductService) GetProducts() (prod []domain.Product, err error) {
	prod, err = provServ.RepoProd.GetAll()

	if err != nil {
		return
	}
	return prod, nil
}

func (prodServ ProductService) GetByID(id int) (product domain.Product, err error) {
	if id < 0 {
		return product, errors.New("no se puede enviar un id negativo")
	}
	product, err = prodServ.RepoProd.GetProdByID(id)
	return
}

func (prodServ ProductService) GetProductsQuery(price float64) (productsQuery []domain.Product, err error) {
	if price < 0 {
		return []domain.Product{}, errors.New("el ID no puede ser igual o menor a 0")
	}

	productsQuery, err = prodServ.RepoProd.GetProductsQuery(price)

	return
}

/*
	func (prodServ ProductService)ExistCodeValue(code string) bool {
		for _, p := range Products {
			if p.CodeValue == code {
				return true
			}
		}
		return false
	}
*/
func ValidateExpiration(expiration string) (err error) {
	re := regexp.MustCompile("([0-9]{2})([/])([0-9]{2})([/])([0-9]{4})")
	//re := regexp.MustCompile("(0?[1-9]|[12][0-9]|3[01])/(0?[1-9]|1[012])/((19|20)\\d\\d)")
	if !re.MatchString(expiration) {
		//c.String(401, "Formato de fecha de expiracion incorrecta")
		return errors.New("formato de fecha de expiracion incorrecta")
	} else {
		reA := strings.Split(expiration, "/")
		dia, errDia := strconv.Atoi(reA[0])
		if errDia != nil {
			//c.String(401, "No se pudo convertir el día de la fecha de expiracion")
			return errors.New("no se pudo convertir el día de la fecha de expiracion")
		}
		if dia < 0 || dia > 31 {
			//c.String(401, "Dia en fecha de expiracion no valido")
			return errors.New("dia en fecha de expiracion no valido")
		}
		mes, errMes := strconv.Atoi(reA[1])
		if errMes != nil {
			return errors.New("no se pudo convertir el mes de la fecha de expiracion")
		}
		if mes < 0 || mes > 12 {
			return errors.New("mes en fecha de expiracion no valido")
		}
	}
	return
}

// Ejercitacion 2
func (prodServ ProductService) Create(resq domain.Request) (prod domain.Product, err error) {
	//validaciones
	if prodServ.RepoProd.ExistCodeValue(resq.CodeValue) {
		return domain.Product{}, fmt.Errorf("%w. %s", ErrAlreadyExist, "url not unique")
	}
	errExpiration := ValidateExpiration(resq.Expiration)
	if errExpiration != nil {
		return domain.Product{}, fmt.Errorf("%w. %s", errExpiration, "Date invalidate")
	}

	cantId++
	newProd := domain.Product{
		//Id:          cantId,
		Name:        resq.Name,
		Quantity:    resq.Quantity,
		CodeValue:   resq.CodeValue,
		IsPublished: resq.IsPublished,
		Expiration:  resq.Expiration,
		Price:       resq.Price,
	}

	prod, err = prodServ.RepoProd.CreateProduct(newProd)

	return
}
