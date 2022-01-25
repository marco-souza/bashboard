package cmd

import (
	"fmt"

	"github.com/webview/webview"
)

func StartDesktopApp() {
	debug := true
	w := webview.New(debug)
	defer w.Destroy()
	fmt.Println("Starting desktop app")

	// setup window
	w.SetTitle("Bashboard")
	w.SetSize(800, 600, webview.HintNone)
	w.Navigate("https://podcodar.com")

	w.Run()
}
