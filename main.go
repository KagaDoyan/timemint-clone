package main

import (
	"fmt"
	"log"

	"go-fiber/api/route"
	"go-fiber/bootstrap"
)

func main() {
	app := bootstrap.App()
	globalEnv := app.Env
	fiber := app.Fiber
	db := app.DB
	Fb := app.Firebase
	route.Setup(fiber, Fb, db)
	log.Fatal(fiber.Listen(fmt.Sprintf(":%v", globalEnv.App.Port)))
}
