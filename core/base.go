package core

import (
	"github.com/getsentry/raven-go"
	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent"
	"github.com/spf13/viper"
)

// base builds all the application dependencies
type base struct {
	config   *viper.Viper
	newrelic newrelic.Application
	sentry   *raven.Client
	driver   *Engine
	router   *router
}

// initializes all the dependencies
func (b *base) init(services []Service) (err error) {
	if err = b.initConfig(); err != nil {
		return
	}
	if err = b.initNewrelic(); err != nil {
		return
	}
	if err = b.initSentry(); err != nil {
		return
	}
	b.initLogger()
	b.initDriver()
	b.initRouter(services)
	return
}

// initializes configration
func (b *base) initConfig() (err error) {
	config := viper.New()
	defaultEnv := "local"
	config.SetDefault("SERVER_ENV", defaultEnv)
	config.SetDefault("DEBUG", true)
	config.SetDefault("ENABLE_NEWRELIC", false)
	config.AutomaticEnv()
	config.AddConfigPath("./config/")
	config.SetConfigName("config." + config.GetString("SERVER_ENV"))
	err = config.ReadInConfig()
	if err != nil && config.GetString("SERVER_ENV") == defaultEnv {
		err = nil // Incase of local config missing, continue
	}
	b.config = config
	return
}

// initializes Newrelic
func (b *base) initNewrelic() (err error) {
	if !b.config.GetBool("ENABLE_NEWRELIC") || !b.config.IsSet("newrelic") {
		return
	}
	config := b.config.GetStringMapString("newrelic")
	newrelicConfig := newrelic.NewConfig(config["app"], config["key"])
	b.newrelic, err = newrelic.NewApplication(newrelicConfig)
	return
}

// initializes Sentry
func (b *base) initSentry() (err error) {
	if !b.config.IsSet("sentry") {
		return
	}
	err = raven.SetDSN(b.config.GetString("sentry.dns"))
	b.sentry = raven.DefaultClient
	return
}

// initializes logger
func (b *base) initLogger() {
	if !b.config.IsSet("logger") {
		return
	}
	config := b.config.GetStringMapString("logger")
	initLogger(config)
}

// initializes the service driver. The default driver is gin gonic.
func (b *base) initDriver() {
	if !b.config.GetBool("DEBUG") {
		gin.SetMode(gin.ReleaseMode) // Switch to "release" mode in production.
	}
	b.driver = gin.New()
}

// initializes the service router with its middlewares
func (b *base) initRouter(services []Service) {
	router := &router{b, services}
	router.init()
	b.router = router
}

func (b *base) clean() {
	// TODO: put cleaning tasks here
}
