package main

import (
	"fmt"
	"os"

	"vigo360.es/fuentes/internal/handler"
)

var (
	version string
)

func main() {
	fmt.Printf("Iniciando vigo360-fuentes versión %s\n", version)
	var PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler.Start(":" + PORT)
}
