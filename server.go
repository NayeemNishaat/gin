package main

import (
	"gin/doc"
	"gin/lib"
	"gin/model"
	"gin/route"
	"os"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Album APIs
// @version 1.0
// @description Album API Doc.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @securityDefinitions.apiKey JWT
// @in header
// @name token

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @externalDocs.description  Swaggo
// @externalDocs.url https://github.com/swaggo/swag

// @BasePath /api/v1

// @schemes http https

func main() {
	lib.StoreLog()
	model.ConnectDataBase()

	server := gin.Default()

	// Chapter: Mount Swagger
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	lib.MountFuncMap(server)
	server.Static("/style", "./public/style")
	server.LoadHTMLGlob("./public/template/*.html")

	apiRouter := server.Group("/api/v1")
	viewRouter := server.Group("/")

	// Chapter: APIs
	route.AlbumRoutes(apiRouter)
	route.AuthRoutes(apiRouter)
	route.User(apiRouter)

	// Chapter: Views
	route.View(viewRouter)

	doc.SwaggerInfo.Host = "localhost:" + os.Getenv("PORT")
	server.Run("localhost:" + os.Getenv("PORT"))
	// server.Run(":" + os.Getenv("PORT"))
	// router.Run(":3000")
	// router.Run()
}

// nodemon -q -e go --signal SIGTERM --exec go run .
// PATH=$(go env GOPATH)/bin:$PATH // Add GO Bin path to PATH
// swag init -g server.go
// swag init -g server.go -o ./doc - Create doc.go
