package routes

import (
	"github.com/camikanerloy/claseWeb1Parte2/cmd/server/handlers"
	"github.com/camikanerloy/claseWeb1Parte2/internal/domain"
	"github.com/camikanerloy/claseWeb1Parte2/internal/product"
	"github.com/gin-gonic/gin"
)

type Router struct {
	db *[]domain.Product
	en *gin.Engine
}

func NewRoute(db *[]domain.Product, en *gin.Engine) *Router {
	return &Router{
		db: db,
		en: en,
	}
}

func (r *Router) SetRoutes() {
	r.SetProduct()
}

// Product
func (r *Router) SetProduct() {
	// instances
	rp := product.NewRepository(*r.db, 500)
	sv := product.NewService(rp)
	h := handlers.NewProductHandler(sv)

	r.en.GET("/ping", h.GetPong())
	prods := r.en.Group("/products")
	{
		prods.GET("/", h.GetProducts())

		prods.GET("/:id", h.GetProductById())

		prods.GET("/search", h.GetProductQuery())

		//Ejercitacion 2
		prods.POST("/", h.CreateProduct())
		prods.PUT("/:id", h.Put())
		prods.PATCH("/:id", h.Patch())
	}
}
