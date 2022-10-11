package delivery

import (
	"git.enigmacamp.com/enigmart-api/delivery/controller"
	"git.enigmacamp.com/enigmart-api/repository"
	"git.enigmacamp.com/enigmart-api/usecase"
	"github.com/gin-gonic/gin"
)

type appServer struct {
	prodUseCase usecase.ProductUseCase
	engine      *gin.Engine
	host        string
}

func Server() *appServer {
	ginEngine := gin.Default()
	prodRepo := repository.NewProductRepository()
	prodUseCase := usecase.NewProductUseCase(prodRepo)
	host := ":8181"
	return &appServer{
		prodUseCase: prodUseCase,
		engine:      ginEngine,
		host:        host,
	}
}

func (a *appServer) initHandlers() {
	controller.NewProductController(a.engine, a.prodUseCase)
}

func (a *appServer) Run() {
	a.initHandlers()
	err := a.engine.Run(a.host)
	if err != nil {
		panic(err)
	}
}
