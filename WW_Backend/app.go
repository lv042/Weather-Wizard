package main

import (
	"github.com/pocketbase/pocketbase"
	"log"
)

func main() {

	//create a new config object
	config := Config{
		DefaultDebug:   true,
		DefaultDataDir: "./pb_data",

		HideStartBanner:  false,
		DataMaxOpenConns: 10,
		DataMaxIdleConns: 10,
		LogsMaxOpenConns: 10,
		LogsMaxIdleConns: 10,
	}

	app := pocketbase.NewWithConfig((*pocketbase.Config)(&config))
	if err := app.Start(); err != nil {
		log.Fatal(err)
	}

	//create a database

}

type Config struct {
	// optional default values for the console flags
	DefaultDebug         bool
	DefaultDataDir       string
	DefaultEncryptionEnv string

	// hide the default console server info on app startup
	HideStartBanner bool

	// optional DB configurations
	DataMaxOpenConns int // default to core.DefaultDataMaxOpenConns
	DataMaxIdleConns int // default to core.DefaultDataMaxIdleConns
	LogsMaxOpenConns int // default to core.DefaultLogsMaxOpenConns
	LogsMaxIdleConns int // default to core.DefaultLogsMaxIdleConns
}
