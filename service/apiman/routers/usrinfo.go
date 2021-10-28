package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/or0986113303/remotedebug/pkg/usr"
)

func (api *API) UsrInfoRouter(Router *gin.RouterGroup) (R gin.IRoutes) {

	HostRouter := Router.Group("usr")
	{
		HostRouter.GET(":id", api.GetUserInfo)
	}
	return HostRouter

}

func (api *API) GetUserInfo(context *gin.Context) {
	context.SetCookie("TomsFavouriteDrink", "Beer", 0, "/", "here.com", false, false)
	id := context.Param("id")
	if id == "16701" {
		tmp := usr.New().Usrinfo(id)
		api.log.Info(tmp)
		context.JSON(http.StatusOK, tmp)
		return
	}
	context.AbortWithStatus(http.StatusNotFound)
}
