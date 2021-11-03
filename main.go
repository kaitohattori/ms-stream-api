package main

import (
	"fmt"
	"log"
	"ms-stream-api/app/controller"
	"ms-stream-api/app/util"
	"ms-stream-api/config"
	_ "ms-stream-api/docs"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	timeout "github.com/vearne/gin-timeout"
)

// @title ms stream api
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

func main() {
	util.LoggingSettings(config.Config.LogFile)
	StartServer()
}

func StartServer() {
	engine := gin.Default()

	// Timeout
	engine.Use(
		timeout.Timeout(
			timeout.WithTimeout(config.Config.APITimeout),
			timeout.WithErrorHttpCode(http.StatusRequestTimeout),
			timeout.WithCallBack(func(r *http.Request) {
				log.Println("Request Timeout : ", r.URL.String())
			})),
	)

	// Controller
	StreamController := controller.NewStreamController()

	// Router
	v1 := engine.Group("/api/v1")
	{
		auth := v1.Group("/video")
		{
			auth.GET(":id/stream", StreamController.GetM3u8File)
			auth.GET(":id/stream/:segName", StreamController.GetHlsFile)
		}
	}
	engine.GET(fmt.Sprintf("%s/*any", config.Config.ApiDocsPath), ginSwagger.WrapHandler(swaggerFiles.Handler))
	engine.Run(":8080")
}
