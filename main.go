package main

import (
	"fmt"
	"net/http"
	"os"

	"vigo360.es/fuentes/internal"
	"vigo360.es/fuentes/internal/database"
	"vigo360.es/fuentes/internal/fuente"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	fmt.Printf("Iniciando vigo360-fuentes\n")
	var PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	var fr = fuente.NewMysqlRepository(database.GetDB())
	var s = internal.NewServer(fr)
	fmt.Printf("Escuchando en puerto %s\n", PORT)
	err := http.ListenAndServe(":"+PORT, s)
	return err
}
