package controller

import (
	"net/http"
	"strconv"

	"gitlab.com/voltunes/api-master-project/auth"
	"gitlab.com/voltunes/api-master-project/helper"
	"gitlab.com/voltunes/api-master-project/model/web"
	"gitlab.com/voltunes/api-master-project/service"

	"github.com/gin-gonic/gin"
)

type ProductControllerImpl struct {
	ProductService service.ProductService
}

func NewProductController(productService service.ProductService) ProductController {
	return &ProductControllerImpl{
		ProductService: productService,
	}
}

func (controller *ProductControllerImpl) FindAll(c *gin.Context, auth *auth.AccessDetails) {
	filters := helper.FilterFromQueryString(c, "unit_id.eq")
	productResponses := controller.ProductService.FindAll(auth, &filters, c)
	webResponse := web.WebResponse{
		Success: true,
		Message: helper.MessageDataFoundOrNot(productResponses),
		Data:    productResponses,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *ProductControllerImpl) FindByID(c *gin.Context, auth *auth.AccessDetails) {
	paramID := c.Param("id")
	productID, err := strconv.Atoi(paramID)
	helper.PanicIfError(err)

	productResponse := controller.ProductService.FindByID(auth, &productID, c)
	webResponse := web.WebResponse{
		Success: true,
		Message: helper.MessageDataFoundOrNot(productResponse),
		Data:    productResponse,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *ProductControllerImpl) Create(c *gin.Context, auth *auth.AccessDetails) {
	request := web.ProductCreateRequest{}
	helper.ReadFromRequestBody(c, &request)

	productResponse := controller.ProductService.Create(auth, &request, c)
	webResponse := web.WebResponse{
		Success: true,
		Message: "Product created successfully",
		Data:    productResponse,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *ProductControllerImpl) Update(c *gin.Context, auth *auth.AccessDetails) {
	paramID := c.Param("id")
	productID, err := strconv.Atoi(paramID)
	helper.PanicIfError(err)

	request := web.ProductUpdateRequest{}
	helper.ReadFromRequestBody(c, &request)

	productResponse := controller.ProductService.Update(auth, &productID, &request, c)
	webResponse := web.WebResponse{
		Success: true,
		Message: "Product updated successfully",
		Data:    productResponse,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *ProductControllerImpl) Delete(c *gin.Context, auth *auth.AccessDetails) {
	paramID := c.Param("id")
	productID, err := strconv.Atoi(paramID)
	helper.PanicIfError(err)

	controller.ProductService.Delete(auth, &productID, c)
	webResponse := web.WebResponse{
		Success: true,
		Message: "Product deleted successfully",
	}

	c.JSON(http.StatusOK, webResponse)
}
