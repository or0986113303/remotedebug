package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type test struct {
	Message string `json:"message"`
}

func (api *API) TestRouter(Router *gin.RouterGroup) (R gin.IRoutes) {

	TestRouter := Router.Group("test")
	{
		TestRouter.GET("message", api.Test)
	}
	return TestRouter

}

func (api *API) Test(context *gin.Context) {
	api.log.Info(gin.H{
		"message": "Hello World",
	})
	context.JSON(http.StatusOK, test{Message: "Hello world"})
}
