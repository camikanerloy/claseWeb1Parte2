package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/camikanerloy/claseWeb1Parte2/cmd/server/handlers"
	"github.com/camikanerloy/claseWeb1Parte2/internal/domain"
	"github.com/camikanerloy/claseWeb1Parte2/internal/product"
	"github.com/camikanerloy/claseWeb1Parte2/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type response struct {
	Data interface{} `json:"data"`
}

func createServer(token string) *gin.Engine {

	if token != "" {
		err := os.Setenv("TOKEN", token)
		if err != nil {
			panic(err)
		}
	}

	db := store.NewStore("./products_copy.json")
	repo := product.NewRepository(501, db)
	service := product.NewService(repo)
	productHandler := handlers.NewProductHandler(service)
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	pr := r.Group("/products")
	{
		pr.GET("", productHandler.GetAll())
		pr.GET(":id", productHandler.GetByID())
		pr.GET("/search", productHandler.Search())
		pr.POST("", productHandler.Post())
		pr.DELETE(":id", productHandler.Delete())
		pr.PATCH(":id", productHandler.Patch())
		pr.PUT(":id", productHandler.Put())
	}
	return r
}

func createRequestTest(method string, url string, body string, token string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	if token != "" {
		req.Header.Add("TOKEN", token)
	}
	return req, httptest.NewRecorder()
}

func loadProducts(path string) ([]domain.Product, error) {
	var products []domain.Product
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(file), &products)
	if err != nil {
		return nil, err
	}
	return products, nil
}

/*
	func writeProducts(path string, list []domain.Product) error {
		bytes, err := json.Marshal(list)
		if err != nil {
			return err
		}
		err = os.WriteFile(path, bytes, 0644)
		if err != nil {
			return err
		}
		return err
	}
*/
func Test_GetAll_OK(t *testing.T) {
	var expectd = response{Data: []domain.Product{}}

	r := createServer("my-secret-token")
	req, rr := createRequestTest(http.MethodGet, "/products", "", "my-secret-token")

	p, err := loadProducts("./products_copy.json")
	if err != nil {
		panic(err)
	}
	expectd.Data = p
	actual := map[string][]domain.Product{}

	r.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
	err = json.Unmarshal(rr.Body.Bytes(), &actual)
	assert.Nil(t, err)
	assert.Equal(t, expectd.Data, actual["data"])
}
