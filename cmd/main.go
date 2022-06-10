package main

import (
	"gitlab.com/a-system/golang-common/internal/application"
)

func main() {
	app := application.New("./config.yml")
	err := app.Bootstrap()

	if err != nil {
		application.Log.Fatal(err)
	}

	err = app.Start()

	if err != nil {
		application.Log.Fatal(err)
	}

	err = app.GracefulShutdown()

	if err != nil {
		application.Log.Fatal(err)
	}
}
