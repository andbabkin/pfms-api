package main

import (
	"log"
	"runtime"

	"github.com/andbabkin/pfms-api/config"
	"github.com/andbabkin/pfms-api/internal/cli"
	"github.com/andbabkin/pfms-api/internal/storage"
)

// init sets runtime settings.
func init() {
	// Verbose logging with file name and line number
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// Use all CPU cores
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	err := config.LoadEnvIntoOS(".env")
	if err != nil {
		log.Fatalln(err)
	}

	err = storage.OpenConns()
	if err != nil {
		log.Fatalln(err)
	}

	cli.Run(Version, BuildDate)
}
