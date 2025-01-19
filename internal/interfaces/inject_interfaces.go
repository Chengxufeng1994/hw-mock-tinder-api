package interfaces

import (
	"net/http"

	"github.com/gin-contrib/cors"
	ginlogger "github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/application"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/config"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/ratelimiter"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/ws"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/interfaces/api/middleware"
	apiv1 "github.com/Chengxufeng1994/hw-mock-tinder-api/internal/interfaces/api/v1"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/pkg/logging"
)

const (
	HealthAPI = "/health"
	MetricAPI = "/metrics"
)

var InterfacesSet = wire.NewSet(
	ProviderRouter,
)

func heartbeatHandler(c *gin.Context) {
	c.AbortWithStatus(http.StatusOK)
}

func metricsHandler(c *gin.Context) {
	promhttp.Handler().ServeHTTP(c.Writer, c.Request)
}

func ProviderRouter(
	logger logging.Logger,
	config *config.Config,
	application *application.ApplicationService,
	rl ratelimiter.RateLimiter,
	wsHub *ws.Hub,
) http.Handler {
	gin.SetMode(config.GinMode)

	router := gin.New()
	root := router.Group("/")

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowMethods = []string{"*"}
	corsConfig.AllowHeaders = []string{"*"}
	corsConfig.ExposeHeaders = []string{"*"}

	root.Use(ginlogger.SetLogger(
		ginlogger.WithUTC(true),
		ginlogger.WithSkipPath([]string{
			HealthAPI,
			MetricAPI,
		}),
	))
	root.Use(gin.Recovery())
	root.Use(cors.New(corsConfig))
	root.Use(middleware.RequestID())
	root.Use(middleware.ErrorHandler)

	root.GET(MetricAPI, metricsHandler)
	root.GET(HealthAPI, heartbeatHandler)
	root.HEAD(HealthAPI, heartbeatHandler)

	apiGroup := root.Group("/api")
	v1Group := apiGroup.Group("/v1")
	apiv1.SetupRouter(logger, v1Group, application, rl, wsHub)

	return router
}
