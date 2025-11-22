package main

import (
	"log"
	"milestone3/be/config"
	"os"

	"github.com/labstack/echo"
)

func main() {
	config.ConnectionDb()
	e := echo.New()

	address := os.Getenv("PORT")
	if err := e.Start(":" + address); err != nil {
		log.Printf("faile to start server %s", err)
	}
}