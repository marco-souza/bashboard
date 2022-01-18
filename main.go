package main

import (
	"fmt"
	"marco-souza/binance-dashboard/cmd"
	"os"
	"os/exec"
	"time"
)

func main() {
	systemStatus := cmd.FetchSystemStatus()
	refreshes := 0

	for {
		clear()

		refreshes++
		fmt.Println("refreshed", refreshes, "times")
		fmt.Printf("System Status: %d => %s\n", systemStatus.Status, systemStatus.Msg)

		time.Sleep(time.Second)
	}
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
