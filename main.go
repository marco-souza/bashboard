package main

import (
	"fmt"
	"marco-souza/binance-dashboard/cmd"
	"os"
	"os/exec"
)

func main() {

	systemStatus := cmd.FetchSystemStatus()
	fmt.Printf("System Status: %d => %s\n", systemStatus.Status, systemStatus.Msg)

	accountSnap := cmd.FetchAccountSnapshot()
	fmt.Printf("Account Snapshot: %d => %s\n", accountSnap.Code, accountSnap.Msg)
	fmt.Println(len(accountSnap.SnapshotVos))

	for _, asset := range accountSnap.SnapshotVos {
		fmt.Println(asset)
	}
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
