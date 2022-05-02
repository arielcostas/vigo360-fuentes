package main

import (
	"fmt"
	"net/http"
	"os"

	"vigo360.es/fuentes/internal"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	fmt.Printf("Iniciando vigo360-fuentes\n")

	var s = internal.Server{}
	s.Routes()
	err := http.ListenAndServe(":8080", &s)
	return err
}
