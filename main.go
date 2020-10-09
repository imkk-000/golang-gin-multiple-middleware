package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func authorizedLoginMiddleware(c *gin.Context) {
	log.Printf("authorization middleware login: %s\n", c.Request.URL)
}

func loginEndpoint(c *gin.Context) {
	queryParams := c.Request.URL.Query()
	log.Printf("doing login: %+v\n", queryParams)
}

func authorizedProjectMiddleware(c *gin.Context) {
	log.Printf("authorization middleware project: %s\n", c.Request.URL)
}

func projectEndpoint(c *gin.Context) {
	queryParams := c.Request.URL.Query()
	log.Printf("doing project: %+v\n", queryParams)
}

func authenticateMiddleware(c *gin.Context) {
	log.Printf("authentication: %s\n", c.Request.URL)
}

func main() {
	r := gin.Default()
	r.Use(authenticateMiddleware)
	loginRouter := r.Group("/")
	createLoginRoutes(loginRouter)
	projectRouter := r.Group("/")
	createProjectRoutes(projectRouter)
	log.Fatal(r.Run(":8080"))
}

func createLoginRoutes(r *gin.RouterGroup) {
	r.
		Use(authorizedLoginMiddleware)
	{
		r.GET("/login", loginEndpoint)
	}
}

func createProjectRoutes(r *gin.RouterGroup) {
	r.Use(authorizedProjectMiddleware)
	{
		r.GET("/project", projectEndpoint)
	}
}
