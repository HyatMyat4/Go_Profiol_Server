package main

import (
	"fmt"
	routes "go-profiol-server/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	fmt.Println("here port : ", port)
	if port == "" {
		port = "4000"
	}

	app := gin.Default()

	app.Use(gin.Logger())

	app.LoadHTMLGlob("./views/*")

	routes.Get_allskill_route(app)

	app.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", map[string]string{"title": "home page"})
	})
	app.GET("/api-v1", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"success": "Access granted for api-v1"})
	})

	if err := app.Run(":" + port); err != nil {
		log.Fatal(err)
	}

}
