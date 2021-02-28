package http

import (
	"fmt"
	"my_grpc_do_zero/model"
	"net/http"
	"my_grpc_do_zero/application/grpc"
	"github.com/labstack/echo/v4"
)

type webServer struct{

}

// NewWebServer retorna um webServer
func NewWebServer() *webServer {
	return &webServer{}
}

func (w webServer) Serve() {
	e := echo.New()
	e.GET("/products", w.getAll)
	e.POST("/products", w.createProduct)
	e.Start(":8080")
}

func (w webServer) getAll(e echo.Context) error {
	return e.JSON(http.StatusOK, grpc.ProductList)
}

func (w webServer) createProduct(e echo.Context) error {
	product := model.NewProduct()

	if err := e.Bind(product); err != nil {
		return err
	}

	fmt.Println(product)

	grpc.ProductList.Add(product)
	return e.JSON(http.StatusOK, product)
}
