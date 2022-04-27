package main

import (
	"os"

	"vigo360.es/fuentes/internal/handler"
)

func main() {
	var PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler.Start(":" + PORT)
}
