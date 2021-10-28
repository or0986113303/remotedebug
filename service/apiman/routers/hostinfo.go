package routers

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/or0986113303/remotedebug/pkg/device"
	"github.com/or0986113303/remotedebug/pkg/system"
)

func (api *API) HostInfoRouter(Router *gin.RouterGroup) (R gin.IRoutes) {

	HostRouter := Router.Group("host")
	{
		HostRouter.GET("disk", api.GetDiskStatus)
		HostRouter.GET("macaddress", api.GetMacAddress)
		HostRouter.GET("hostname", api.GetHostname)
		HostRouter.GET("cpuuseage", api.GetCPUUseage)
	}
	return HostRouter

}

// @Tags host
// @Summary get disk status
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} device.DiskStatus "capacity of disk"
// @Router /auth/api/v1/host/disk [get]
func (api *API) GetDiskStatus(context *gin.Context) {
	tmp, _ := device.New().GetDiskStatus("/")
	api.log.Info(tmp)
	context.JSON(http.StatusOK, tmp)
}

func (api *API) GetMacAddress(context *gin.Context) {
	tmp, _ := device.New().GetMacAddress()
	api.log.Info(tmp)
	context.JSON(http.StatusOK, tmp)
}

func (api *API) GetHostname(context *gin.Context) {
	tmp, _ := device.New().GetHostname()
	api.log.Info(tmp)
	context.JSON(http.StatusOK, tmp)
}

func (api *API) GetCPUUseage(context *gin.Context) {
	tmp, _ := system.New().Sysinfo()
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		api.log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		api.log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		api.log.Fatal(err)
	}
	api.log.Printf("%s", body)
	api.log.Info(tmp)
	context.JSON(http.StatusOK, tmp)
}
