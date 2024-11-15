package main

import (
	"flag"

	"github.com/gofiber/fiber/v2"
	"github.com/khaledibrahim1015/hotel-reservation/api"
)

func main() {

	// command line
	var listenAddr *string
	listenAddr = flag.String("listenAddr", ":5000", "the listen address of api server ")
	flag.Parse()

	app := fiber.New()

	apiv1 := app.Group("/api/v1")
	apiv1.Get("/user", api.HandleGetUsers)
	apiv1.Get("/user/:id", api.HandleGetUser)

	app.Listen(*listenAddr)
}
