package apiman

import (
	"context"
	"net/http"

	"github.com/or0986113303/remotedebug/service/apiman/routers"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type APIRouter interface {
	StartServer()
	StopServer()
	RegisterFunc()
	RegisterRouter()
}

type Worker struct {
	ctx          context.Context
	cancel       context.CancelFunc
	server       *http.Server
	log          *logrus.Entry
	routerEngine *gin.Engine
	v1table      routers.APITable
}

// New will new a API manager
func New() (worker *Worker) {
	gin.SetMode(gin.ReleaseMode)
	ctx, cancel := context.WithCancel(context.TODO())
	worker = &Worker{
		ctx:          ctx,
		cancel:       cancel,
		server:       nil,
		log:          logrus.WithFields(logrus.Fields{"origin": "apiman"}),
		routerEngine: gin.New(),
		v1table:      routers.New(),
	}
	worker.RegisterRouter()
	worker.routerEngine.Use(gin.Logger())
	worker.routerEngine.Use(gin.Recovery())
	worker.routerEngine.NoRoute(func(c *gin.Context) {
		// fmt.Println("page not found")
		worker.log.Panic("page not found")
	})
	return
}

// StartServer ...
func (w *Worker) StartServer() error {
	w.server = &http.Server{
		Addr:    "0.0.0.0:5000",
		Handler: w.routerEngine,
	}
	go func() {
		defer w.cancel()
		for w.ctx.Err() == nil {
			w.log.Info("ready: ", w.server.Addr)
			if err := w.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				w.log.Fatalf("listen: %s\n", err)
				return
			}
		}
	}()
	return nil
}

// StopServer ...
func (w *Worker) StopServer() {
	w.cancel()
}

// RegisterFunc ...
func (w *Worker) RegisterFunc() {

}

// RegisterRouter ...
func (w *Worker) RegisterRouter() {
	PublicGroup := w.routerEngine.Group("")
	w.v1table.RegisterRouter(PublicGroup)
}

// GetRouter
func (w *Worker) GetRouter() *gin.Engine {
	return w.routerEngine
}
