package main

import (
	"fmt"

	"github.com/vector-ops/go-starter/configs"
	"github.com/vector-ops/go-starter/internal/bootstrap"
)

func main() {
	configs.LoadEnv()
	app := bootstrap.NewApp()
	app.Run()
	fmt.Println(configs.GetEnv("POSTGRES_DB", ""))
}
