package main

import (
	"github.com/startcodextech/octopuslb/internal/api"
	"github.com/startcodextech/octopuslb/internal/http"
	"github.com/startcodextech/octopuslb/internal/logs"
)

func init() {
	logs.Init()
}

func main() {
	app, err := api.Init()
	if err != nil {
		panic(err)
	}

	server := http.NewServer(app)

	server.Start()

}
