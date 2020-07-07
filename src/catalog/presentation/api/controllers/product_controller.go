package controllers

import (
	"fmt"
	"net/http"

	"github.com/NicolasDeveloper/store/src/catalog/core/dtos"

	"github.com/NicolasDeveloper/store/src/catalog/core/usecases"
	drivenadapters "github.com/NicolasDeveloper/store/src/catalog/infrastructure/driven-adapters"

	"github.com/NicolasDeveloper/store/src/catalog/infrastructure/dbcontext"
	"github.com/NicolasDeveloper/store/src/catalog/presentation/api/common"
)

//ProductController controller
type ProductController struct {
	ctx dbcontext.DbContext
	common.Controller
}

//NewProductController contructor
func NewProductController(ctx dbcontext.DbContext) ProductController {
	return ProductController{
		ctx: ctx,
	}
}

//Index get products in database
func (c *ProductController) Index(w http.ResponseWriter, r *http.Request) {

}

//Post insert product
func (c *ProductController) Post(w http.ResponseWriter, r *http.Request) {
	dto := &dtos.ProductDTO{}
	c.GetContent(dto, r)
	fmt.Println(dto)

	repository, _ := drivenadapters.NewProductsRepositoryAdapter(c.ctx)
	usecase := usecases.NewRegisterUseCase(repository)
	usecase.Execute(*dto)

	c.SendJSON(
		w,
		dto,
		http.StatusOK,
	)
}
