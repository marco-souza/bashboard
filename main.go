package main

import (
	"fmt"
	"marco-souza/binance-dashboard/cmd"
)

func main() {
	done := make(chan bool)
	// Start static server
	go cmd.StartPageServer("8001")

	go cmd.StartApp(&done) // sync / blocking

	fmt.Println("Waiting to finish")

	<-done
	fmt.Println("Finishing")
}
