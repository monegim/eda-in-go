package main

import (
	"eda-in-go/internal/config"
	"fmt"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
func run() error {
	var cfg config.AppConfig

	cfg, err := config.
}