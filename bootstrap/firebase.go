package bootstrap

import (
	"context"
	"go-fiber/core/logs"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func InitializeFirebaseApp() (app *firebase.App) {

	var err error
	ctx := context.Background()
	opt := option.WithCredentialsFile(GlobalEnv.App.FirebasePath)

	// Initialize the Firebase Admin SDK
	app, err = firebase.NewApp(ctx, nil, opt)

	if err != nil {
		panic("can not connect to firebase")
	}

	logs.Info("firebase connection success")

	return app
}
