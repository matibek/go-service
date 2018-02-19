package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent/_integrations/nrgin/v1"
)

// router is responsible for all rounting in the service
type router struct {
	base     *base
	services []Service
}

// Initializes driver with middlewares
func (r *router) init() {
	// Middlewares
	driver := r.base.driver
	if r.base.newrelic != nil {
		driver.Use(nrgin.Middleware(r.base.newrelic))
	}
	if r.base.config.GetBool("DEBUG") {
		driver.Use(gin.Logger())
	}
	// router error handler
	driver.Use(r.errorMiddleware)

	// Core routes
	r.registerRoute()

	// Service routes
	for _, service := range r.services {
		service.RegisterRoute(driver)
	}
}

// RegisterRoute create all service routings
func (r *router) registerRoute() {
	driver := r.base.driver
	// global routes
	driver.GET("/health", r.healthController)
	driver.GET("/", r.homeController)
}

func (r *router) homeController(c *Context) {
	appName := r.base.config.GetString("name")
	version := r.base.config.GetString("version")
	c.JSON(http.StatusOK, gin.H{"app": appName, "version": version})
}

func (r *router) healthController(c *Context) {
	c.JSON(http.StatusOK, gin.H{"status": "healthy"})
	// TODO: add test for database here
}

// errorMiddleware handle errors on all controllers
func (r *router) errorMiddleware(c *Context) {
	defer func() {
		if err := recover(); err != nil {
			Logger.Error("Unexpected router error: ", err)
			if serverErr, ok := err.(*ServerError); ok {
				c.AbortWithStatus(serverErr.HttpStatusCode)
			} else {
				c.AbortWithStatus(500)
			}
		}
	}()
	c.Next()
}
