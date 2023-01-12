package routes

import (
	"github.com/camikanerloy/claseWeb1Parte2/cmd/server/handlers"
	"github.com/camikanerloy/claseWeb1Parte2/internal/product"
	"github.com/camikanerloy/claseWeb1Parte2/pkg/store"
	"github.com/gin-gonic/gin"
)

type Router struct {
	en *gin.Engine
}

func NewRoute(en *gin.Engine) *Router {
	return &Router{
		en: en,
	}
}

func (r *Router) SetRoutes() {
	r.SetProduct()
}

// Product
func (r *Router) SetProduct() {
	// instances
	js := store.NewStore("/Users/CKANER/bootcamp/claseWeb1Parte2/products.json")
	rp := product.NewRepository(500, js)
	sv := product.NewService(rp)
	h := handlers.NewProductHandler(sv)

	r.en.GET("/ping", h.GetPong())
	prods := r.en.Group("/products")
	prods.Use(handlers.TokenAuthMiddleware())
	{
		prods.GET("/", h.GetAll())

		prods.GET("/:id", h.GetByID())

		prods.GET("/search", h.Search())

		//Ejercitacion 2
		prods.POST("/", h.Post())
		prods.PUT("/:id", h.Put())
		prods.PATCH("/:id", h.Patch())
		prods.DELETE("/:id", h.Delete())
	}
}
