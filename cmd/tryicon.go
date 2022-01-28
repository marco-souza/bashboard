package cmd

import (
	"fmt"

	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
)

func StartApp() {
	onExit := func() {
		fmt.Println("Exiting...")
	}

	systray.Run(onReady, onExit)
}

func onReady() {
	// Sys tray icon
	systray.SetTemplateIcon(icon.Data, icon.Data)
	systray.SetTitle("Awesome App")
	systray.SetTooltip("Lantern")

	// Menu items
	mShowLantern := systray.AddMenuItem("Show Lantern", "")
	mShowWikipedia := systray.AddMenuItem("Show Wikipedia", "")
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")

	// Sets the icon of a menu item. Only available on Mac.
	mQuit.SetIcon(icon.Data)

	registerClickHandlers := func() {
		for {
			select {
			// TODO: show window
			case <-mShowLantern.ClickedCh:
				fmt.Println("Lantern")
			case <-mShowWikipedia.ClickedCh:
				fmt.Println("Wiki")
			case <-mQuit.ClickedCh:
				systray.Quit()
			}
		}
	}
	go registerClickHandlers()
}
