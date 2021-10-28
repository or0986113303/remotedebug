package routers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type APITable interface {
	RegisterRouter(Router *gin.RouterGroup) (R gin.IRoutes)
}

type API struct {
	ctx    context.Context
	cancel context.CancelFunc
	log    *logrus.Entry
}

func New() (api *API) {
	gin.SetMode(gin.ReleaseMode)
	ctx, cancel := context.WithCancel(context.TODO())
	api = &API{
		ctx:    ctx,
		cancel: cancel,
		log:    logrus.WithFields(logrus.Fields{"origin": "routers"}),
	}

	return
}

func (api *API) RegisterRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	// v1 group route
	v1 := Router.Group("/api/v1/")
	{
		api.TestRouter(v1)
		api.HostInfoRouter(v1)
		api.UsrInfoRouter(v1)
	}

	return v1
}
