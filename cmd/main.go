package main

import (
	"fmt"
	"github.com/startcodextech/octopuslb/internal/logs"
	"os/exec"
	"strings"
)

func init() {
	logs.Init()
}

func main() {
	cmd := exec.Command("networksetup", "-listnetworkserviceorder")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Procesar la salida
	lines := strings.Split(string(output), "\n")
	mapping := make(map[string]string)

	for _, line := range lines {
		if strings.Contains(line, "Hardware Port") {
			parts := strings.Split(line, ", Device: ")
			if len(parts) == 2 {
				device := strings.TrimSpace(parts[1])
				device = strings.TrimRight(device, ")") // Eliminar paréntesis de cierre si existen
				service := strings.Split(parts[0], ": ")[1]
				if device != "" { // Solo agregar si 'device' no está vacío
					mapping[device] = service
				}
			}
		}
	}

	// Mostrar el mapeo
	for device, service := range mapping {
		fmt.Printf("Interface: %s, Service: %s\n", device, service)
	}
}
