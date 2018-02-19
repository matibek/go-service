# Service Scaffolding for Go Project
This project provides the basic structure for [gin-gonic](https://github.com/gin-gonic/gin) go service.

# Prerequisite
Go 1.9 or later

# How to run
- `$ go install`
- `$ service-scaffolding-go` (if you have `$GOPATH/bin` in your `PATH` variable)
- Visit http://localhost:8080/ or http://localhost:8080/helloworld?name=guest

# How to create service
- Fork this project with disabled syncing
  - For future changes, setup upstream to the fork - [refer](https://help.github.com/articles/syncing-a-fork/)
- Create neccssary config files (`config.[ENV].json`) for every env (i.e local, dev, prod, etc..), inside `config/`:
  - Make neccessary changes by refering `config/config.sample.json`
- Setup the service instance inside `main.go`

# Setup
- Set `SERVER_ENV` on environment to load the env configration file (if it is not set, `config.local.json` will be loaded)
- Set `DEBUG=false` to disable debug mode
- Set `ENABLE_NEWRELIC=true` to enable newrelic (Note: we need to setup the configration of newrelic too)

# Doc
[![GoDoc](https://godoc.org/github.com/matibek/service-scaffolding-go/core?status.svg)](https://godoc.org/github.com/matibek/service-scaffolding-go/core)

# TODO
- unit test