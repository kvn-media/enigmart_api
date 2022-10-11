package controller

import (
	"net/http"

	"git.enigmacamp.com/enigmart-api/models"
	"git.enigmacamp.com/enigmart-api/usecase"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	router      *gin.Engine
	prodUseCase usecase.ProductUseCase
}

func (pc *ProductController) CreateNewProduct(ctx *gin.Context) {
	var newProduct *models.Product
	err := ctx.BindJSON(&newProduct)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		product := pc.prodUseCase.CreateNewProduct(newProduct)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK",
			"data":    product,
		})
	}
}

func (pc *ProductController) GetAllProduct(ctx *gin.Context) {
	products := pc.prodUseCase.GetAllProduct()
	ctx.JSON(http.StatusOK, gin.H {
		"message": "OK",
		"data": products,
	})
}

// func (pc *ProductController) GetProduct(ctx *gin.Context) {
// 	product := pc.prodUseCase.GetProduct()
// 	ctx.JSON(http.StatusOK, gin.H {
// 		"message": "OK",
// 		"data": product,
// 	})
// }

func NewProductController(router *gin.Engine, prodUseCase usecase.ProductUseCase) *ProductController {
	newProdController := ProductController{
		router:      router,
		prodUseCase: prodUseCase,
	}
	router.POST("/product", newProdController.CreateNewProduct)
	router.GET("/product", newProdController.GetAllProduct)
	// router.GET("/product/:id", newProdController.GetProduct)
	// router.DELETE("/product/:id", newProdController.DeleteProduct)
	// router.PUT("/product", newProdController.UpdateProduct)
	return &newProdController
}
