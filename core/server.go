package core

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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
	Health() error
	Clean() error
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
	// Errors contains applications errors
	Errors *Error
)

type server struct {
	base       *base
	services   []Service
	httpServer *http.Server
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
	s.httpServer = &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: Driver,
	}
	// Exit handler
	s.registerExitHandler()
	// Start connection
	Logger.Info(fmt.Sprintf("Service is running on http://localhost:%d", port))
	s.httpServer.ListenAndServe()
}

// Exit will gracefully shutdown the server with a timeout of 5 seconds.
func (s *server) Exit(sig os.Signal) {
	Logger.Info("Service is exiting with signal ", sig)
	// Clean
	s.base.clean()
	for _, service := range s.services {
		err := service.Clean()
		if err != nil {
			Logger.Error("Failed to clean service: ", err)
		}
	}
	// Shutdown http server
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.httpServer.Shutdown(ctx); err != nil {
		Logger.Fatal("Server Shutdown Error:", err)
	}
	// Exit
	if sig == nil {
		os.Exit(0) // OK
	}
	os.Exit(2) // SIGINT
}

// registerExitHandler attaches a handler for various exit conditions
func (s *server) registerExitHandler() {
	// Wait for interrupt signal to gracefully shutdown the server
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		s.Exit(sig)
	}()
}

// NewServer initializes a new instance of server with all services
func NewServer(services ...Service) Server {
	base := &base{}
	server := &server{base, services, nil}
	return Server(server)
}
