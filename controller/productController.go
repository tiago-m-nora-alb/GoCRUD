package controller

import (
	"github.com/TiagoNora/GoCRUDV2/schemas"
	"github.com/TiagoNora/GoCRUDV2/service"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
)

type ProductController interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	FindById(c *gin.Context)
	FindAll(c *gin.Context)
}

type productController struct {
	productService service.ProductService
}

// @Summary      Create a new product
// @Description  Create a new product with name, description, price, and stock
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        product  body      schemas.ProductRequest  true  "Product data"
// @Router       /product [post]
func (p *productController) Create(c *gin.Context) {
	log.Info().Msg("Called create from Product Controller")
	productRequest := schemas.ProductRequest{}
	if err := c.ShouldBindJSON(&productRequest); err != nil {
		sendError(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	product := schemas.Product{
		Name:        productRequest.Name,
		Description: productRequest.Description,
		Price:       productRequest.Price,
		Stock:       productRequest.Stock,
	}

	product, err := p.productService.Create(product)
	if err != nil {
		sendError(c, http.StatusInternalServerError, "Error creating product")
		return
	}
	log.Info().Msg("Product created")
	sendSuccess(c, "Create Product", product)

}

// @Summary      Update a product by ID
// @Description  Update a product by ID with name, description, price, and stock
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        id       path      string  true  "Product ID"
// @Param        product  body      schemas.ProductRequest  true  "Product data"
// @Router       /product/{id} [put]
func (p *productController) Update(c *gin.Context) {
	log.Info().Msgf("Called update from Product Controller id: %s", c.Param("id"))
	id := c.Param("id")

	productRequest := schemas.ProductRequest{}
	if err := c.ShouldBindJSON(&productRequest); err != nil {
		sendError(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	productNew := schemas.Product{
		Name:        productRequest.Name,
		Description: productRequest.Description,
		Price:       productRequest.Price,
		Stock:       productRequest.Stock,
	}

	productFound, err := p.productService.FindById(id)
	if err != nil {
		sendError(c, http.StatusInternalServerError, "Error getting product")
		return
	}

	productUpdated, err := p.productService.Update(productFound, productNew)
	if err != nil {
		sendError(c, http.StatusInternalServerError, "Error updating product")
	}

	sendSuccess(c, "Update Product", productUpdated)

}

// @Summary      Delete a product by ID
// @Description  Delete a product by ID
// @Tags         Products
// @Param        id  path      string  true  "Product ID"
// @Router       /product/{id} [delete]
func (p *productController) Delete(c *gin.Context) {
	log.Info().Msgf("Called delete from Product Controller id: %s", c.Param("id"))
	idParam := c.Param("id")
	product, err := p.productService.Delete(idParam)
	if err != nil {
		sendError(c, http.StatusInternalServerError, "Error deleting product")
		return
	}
	sendSuccess(c, "delete-product", product)
}


// @Summary      Find a product by ID
// @Description  Find a product by ID
// @Tags         Products
// @Param        id  path      string  true  "Product ID"
// @Router       /product/{id} [get]
func (p *productController) FindById(c *gin.Context) {
	log.Info().Msgf("Called find by id from Product Controller id: %s", c.Param("id"))
	idParam := c.Param("id")
	product, err := p.productService.FindById(idParam)
	if err != nil {
		sendError(c, http.StatusInternalServerError, "Error getting product")
		return
	}
	sendSuccess(c, "get-product", product)
}

// @Summary      Find all products
// @Description  Find all products
// @Tags         Products
// @Router       /product [get]
func (p *productController) FindAll(c *gin.Context) {
	log.Info().Msg("Called findAll from Product Controller")
	all, err := p.productService.FindAll()
	if err != nil {
		sendError(c, http.StatusInternalServerError, "Error creating product")
		return
	}
	sendSuccess(c, "list-products", all)
}

func NewProductController() ProductController {
	return &productController{
		productService: service.NewProductService(),
	}

}
