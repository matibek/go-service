# Go Service
This project provides the basic structure for [gin-gonic](https://github.com/gin-gonic/gin) based go service.

# Prerequisite
Go 1.9 or later

# How to run
- `$ dep ensure`
- `$ go install`
- `$ go-service` (if you have `$GOPATH/bin` in your `PATH` variable)
- Visit http://localhost:8080/ or http://localhost:8080/helloworld?name=guest

# How to create service
- Create a `config/` dir inside your repo and add neccssary config files (`config.[ENV].json`) for every env (i.e local, dev, prod, etc..):
  - Make neccessary changes by refering `config/config.sample.json`
- Attach the service instance to the server. Similar to `main.go`.
```go
package main

import (
  "github.com/matibek/go-service/core"
  "github.com/username/your-project/yourservice"
)

func main() {
	yourservice := yourservice.NewService()
	server := core.NewServer(yourservice)
	server.Start()
}
```

# Setup
- Set `SERVER_ENV=env` to load the env configration file (if it is not set, `config.local.json` will be loaded)
- Set `DEBUG=false` to disable debug mode
- Set `ENABLE_NEWRELIC=true` to enable newrelic (Note: you also need to set the config)

# Doc
[![GoDoc](https://godoc.org/github.com/matibek/go-service/core?status.svg)](https://godoc.org/github.com/matibek/go-service/core)

# Config
- Application configrations are defined under `/config`. Depending on `SERVER_ENV`, the config file will be loaded. Note: all server env config varaibles are overwritten by varaibles on config file.
- Use `config.local.json` to put all sensetive configs without pushing it to source control.
- Config uses [Viper](https://github.com/spf13/viper) internally. So, the interface is the same with Viper.
```go
import "github.com/matibek/go-service/core"
...
value := core.Config.GetBool("configKey")
```
## Logger
```go
import "github.com/matibek/go-service/core"
...
core.Logger.Infof("hello %s", "world")
// prints: INFO[2018-02-26 12:20:46] hello world
```
- Set the config variable `logger.level` for to adjust the log level (i.e. `info`, `warn`, etc..)
- Set the config variable `logger.file` to file path when you want to log to file.

## newrelic
- Set the config variable `newrelic.app` and `newrelic.key`
- To enable newrelic, set `ENABLE_NEWRELIC=true` either using config file or server env

## sentry
- Set the config variable `sentry.dns` to enable sentry

# Errors
- Please avoid using `panics`. Always return Error to router context.
- The error middleware will handle all router errors and by sending error response and logging
> visit http://localhost:8080/helloworld?name=error to see a sample error
```go
import "http"
import "github.com/matibek/go-service/core"

func sampleController(c *core.Context) {
	result, err := sampleTask()
	if err != nil {
    serverErr := core.Errors.NewWithMessage(err, "Something went wrong!")
    c.Errors(serverErr)  
    // Error middleware will send a response with 500 
    // with JSON Body: {"error": "Something went wrong!"}
    // It also make an error log with the stack trace
    return
	}
	c.String(http.StatusOK, result)
}

```
- By default, the response code for error is `500`. You can change the response code by setting `HttpStatusCode` property of the Error.

# TODO
- unit test