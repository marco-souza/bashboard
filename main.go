package main

import (
	"marco-souza/binance-dashboard/cmd"
)

func main() {
	// Start static server
	go cmd.StartPageServer("8001")

	cmd.StartDesktopApp() // sync / blocking
}
