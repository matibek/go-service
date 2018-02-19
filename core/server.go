package core

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Server provides methods for controlling application lifecycle
type Server interface {
	Init()
	Start()
	Exit(sig os.Signal)
}

// Service provides methods for interacting with a microservice instance
type Service interface {
	RegisterRoute(serviceDriver *Engine)
}

// Context is a router context
type Context = gin.Context

// Engine is a http service drive
type Engine = gin.Engine

var (
	// Logger is the default application logger
	Logger = getLogger()
	// Config is the application configration
	Config *viper.Viper
	// Driver is applications http driver
	Driver *Engine
)

type server struct {
	base     *base
	services []Service
}

// Init initializes the dependencies
func (s *server) Init() {
	err := s.base.init(s.services)
	if err != nil {
		panic(err)
	}
	Config = s.base.config
	Driver = s.base.driver
}

// Start will start the server with service
func (s *server) Start() {
	if Driver == nil {
		// Initialize server
		s.Init()
	}
	Config.SetDefault("server.port", 8080)
	port := Config.GetInt("server.port")
	Logger.Info(fmt.Sprintf("Service is running on http://localhost:%d", port))
	Driver.Run(fmt.Sprintf(":%d", port))
}

// Exit will gracefully shuts down the server
func (s *server) Exit(sig os.Signal) {
	Logger.Info("Service is exiting with signal ", sig)
	s.base.clean()
	if sig == nil {
		os.Exit(0) // OK
	}
	os.Exit(2) // SIGINT
}

// registerExitHandler attaches a handler for various exit conditions
func registerExitHandler(server *server) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		server.Exit(sig)
	}()
}

// NewServer initializes a new instance of server with all services
func NewServer(services ...Service) Server {
	base := &base{}
	server := &server{base, services}
	registerExitHandler(server)
	return Server(server)
}
