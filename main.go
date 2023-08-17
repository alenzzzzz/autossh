package main

import (
	"autossh/app"
)

var (
	Version = "unknown"
	Build   = "unknown"
)

func main() {
	app.Version = Version
	app.Build = Build
	app.Run()
}
