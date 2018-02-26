package core

import (
	"net/http"

	"github.com/gin-contrib/sentry"
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
	debug := r.base.config.GetBool("DEBUG")
	if debug {
		driver.Use(gin.Logger())
	}
	if r.base.newrelic != nil {
		driver.Use(nrgin.Middleware(r.base.newrelic))
		Logger.Info("Newrelic is enalbed!")
	}
	// Recovery Middlewares - order is in LIFO
	driver.Use(r.errorRecovery(debug)) // our error handler should be call at last
	if r.base.sentry != nil {
		driver.Use(sentry.Recovery(r.base.sentry, false))
		Logger.Info("Sentry is enalbed!")
	}

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
	for _, service := range r.services {
		err := service.Health()
		if err != nil {
			Logger.Error("Service health error: ", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"server": "healthy"})
}

// errorRecovery returns a middleware that handle errors on all controllers
func (router) errorRecovery(debug bool) gin.HandlerFunc {
	return func(c *Context) {
		defer func() {
			// Panics
			if err := recover(); err != nil {
				Logger.Errorf("Unexpected router error: %+v", err)
				c.AbortWithStatus(http.StatusInternalServerError)
			}
			// Errors
			if err := c.Errors.Last(); err != nil {
				Logger.Errorf("Router error: %+v", err.Err)
				switch err.Err.(type) {
				case *Error:
					serverErr := err.Err.(*Error)
					c.AbortWithStatusJSON(serverErr.HttpStatusCode, serverErr.JSON(debug))
				default:
					c.AbortWithStatus(http.StatusInternalServerError)
				}
			}
		}()
		c.Next()
	}
}
