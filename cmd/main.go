package main

import (
	"github.com/startcodextech/managerlb/internal/api"
	"github.com/startcodextech/managerlb/internal/logs"
)

func init() {
	logs.Init()
}

func main() {
	_, err := api.Init()
	if err != nil {
		panic(err)
	}

}
