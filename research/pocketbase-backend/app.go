package main

import (
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"log"
)

func main() {
	app := pocketbase.New()

	app.OnBeforeBootstrap().Add(func(e *core.BootstrapEvent) error {
		log.Println(e.App)
		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal("Error: ", err)
	}
}
