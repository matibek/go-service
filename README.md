# Scaffold for Go Service
This project provides the basic structure for [gin-gonic](https://github.com/gin-gonic/gin) go service.

# Prerequisite
Go 1.9 or later

# How to run
- `$ go install`
- `$ scaffold-go-service` (if you have `$GOPATH/bin` in your `PATH` variable)
- Visit http://localhost:8080/ or http://localhost:8080/helloworld?name=guest

# How to create service
- Fork this project and disable auto syncing
  - For future changes, setup upstream to the repo and sync - [refer](https://help.github.com/articles/syncing-a-fork/)
- Create neccssary config files (`config.[ENV].json`) for every env (i.e local, dev, prod, etc..), inside `config/`:
  - Make neccessary changes by refering `config/config.sample.json`
- Setup the service instance inside `main.go` and remove `helloworld` service.

# Setup
- Set `SERVER_ENV=env` to load the env configration file (if it is not set, `config.local.json` will be loaded)
- Set `DEBUG=false` to disable debug mode
- Set `ENABLE_NEWRELIC=true` to enable newrelic (Note: you also need to set the config)

# Doc
[![GoDoc](https://godoc.org/github.com/matibek/scaffold-go-service/core?status.svg)](https://godoc.org/github.com/matibek/scaffold-go-service/core)

# Config
- Application configrations are defined under `/config`. Depending on `SERVER_ENV`, the config file will be loadded. Note that, all env varaibles are overwritten by env config file.
- Use `config.local.json` to put all sensetive configs without pushing it to source control.
- Config uses [Viper](https://github.com/spf13/viper) internally. So, the interface same with Viper.
```go
import "github.com/matibek/scaffold-go-service/core"
...
value := core.Config.GetBool("configKey")
```
## Logger
```go
import "github.com/matibek/scaffold-go-service/core"
...
core.Logger.Infof("hello %s", "world")
// prints: INFO[2018-02-26 12:20:46] hello world
```
- set `logger.level` for log level (i.e. `info`, `warn`, etc..)
- set `logger.file` to file path, if we want to log to file

## newrelic
- set the config in `newrelic.app` and `newrelic.key`
- to enable newrelic, set `ENABLE_NEWRELIC=true` either in config or server env

## sentry
- Set `sentry.dns` config to enable sentry

# Errors
- Please avoid panics and return the error to router
- The error middleware will handle the response and logging
```go
import "http"
import "github.com/matibek/scaffold-go-service/core"

func sampleController(c *core.Context) {
	result, err := sampleTask()
	if err != nil {
    serverErr := core.Errors.NewWithMessage(err, "Something went wrong!")
    c.Errors(serverErr)  
    // Error middleware will send a response with 500 with JSON Body: {"error": "Something went wrong!"}
    // It also make an error log with stack trace
		return
	}
	c.String(http.StatusOK, result)
}

```
- The response code for error `500` by default. You can change the response code by setting `HttpStatusCode` property of the Error.

# TODO
- unit test