package main

import (
	"fmt"
	"hello-world/service"
)

func main() {
	path := "./conf.yaml"
	app, err := service.NewApplication(path)
	if err != nil {
		fmt.Println("Couldn't create a new application to start: ", err)
		if app == nil {
			return
		}
	}
	err = app.Run()
	if err != nil {
		app.Logger.Errorf("An error occurred while running the app: %v", err)
	}
}
