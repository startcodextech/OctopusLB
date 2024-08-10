package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("Error obteniendo la ruta del ejecutable:", err)
		return
	}
	log.Println("Ruta del ejecutable:", exePath)

	binPathCaddy := filepath.Join(dir, "bin/darwin/arm64/octopus-caddy")

	out, err := exec.Command("sh", "-c", binPathCaddy+" start").CombinedOutput()
	println(string(out))
	if err != nil {
		panic(err)
	}
}
