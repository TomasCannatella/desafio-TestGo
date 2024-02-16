package application

import (
	"app/internal"
	"app/internal/handler"
	"app/internal/repository"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// ConfigApplicationDefault is the configuration of the application.
type ConfigApplicationDefault struct {
	// Addr is the address of the application.
	Addr string
}

// NewApplicationDefault returns a new ApplicationDefault.
func NewApplicationDefault(cfg *ConfigApplicationDefault) *ApplicationDefault {
	// default values
	defaultRt := chi.NewRouter()
	defaultCfg := &ConfigApplicationDefault{
		Addr: ":8080",
	}
	if cfg != nil {
		if cfg.Addr != "" {
			defaultCfg.Addr = cfg.Addr
		}
	}

	return &ApplicationDefault{
		rt:   defaultRt,
		addr: defaultCfg.Addr,
	}
}

// ApplicationDefault is an struct that implements the Application interface.
type ApplicationDefault struct {
	// rt is the router of the application.
	rt *chi.Mux
	// addr is the address of the application.
	addr string
}

// TearDown tears down the application.
// - should be used as a defer function
func (a *ApplicationDefault) TearDown() (err error) {
	return
}

// SetUp sets up the application.
func (a *ApplicationDefault) SetUp() (err error) {
	// dependencies
	// - repository
	db := map[int]internal.Product{
		1: {Id: 1, ProductAttributes: internal.ProductAttributes{Description: "product 1", Price: 10, SellerId: 1}},
		2: {Id: 2, ProductAttributes: internal.ProductAttributes{Description: "product 2", Price: 20, SellerId: 2}},
		3: {Id: 3, ProductAttributes: internal.ProductAttributes{Description: "product 3", Price: 30, SellerId: 3}},
		4: {Id: 4, ProductAttributes: internal.ProductAttributes{Description: "product 4", Price: 40, SellerId: 4}},
		5: {Id: 5, ProductAttributes: internal.ProductAttributes{Description: "product 5", Price: 50, SellerId: 5}},
	}
	rpProduct := repository.NewProductsMap(db)
	// - handler
	hdProduct := handler.NewProductsDefault(rpProduct)

	// router
	// - middleware
	a.rt.Use(middleware.Logger)
	a.rt.Use(middleware.Recoverer)
	// - endpoints
	a.rt.Route("/product", func(r chi.Router) {
		// - GET /product
		r.Get("/", hdProduct.Get())
	})

	return
}

// Run runs the application.
func (a *ApplicationDefault) Run() (err error) {
	err = http.ListenAndServe(a.addr, a.rt)
	return
}
