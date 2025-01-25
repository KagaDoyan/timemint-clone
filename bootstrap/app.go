package bootstrap

import (
	firebase "firebase.google.com/go/v4"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Application struct {
	Env      *Env
	Fiber    *fiber.App
	DB       *gorm.DB
	Firebase *firebase.App
}

var GlobalEnv Env

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	GlobalEnv = *NewEnv()
	app.Fiber = NewFiber()
	app.DB = NewDatabaseConnection(app.Env)
	app.Firebase = InitializeFirebaseApp()

	return *app
}
