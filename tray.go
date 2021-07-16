package main

import (
	"mpvctl/icon"

	"github.com/getlantern/systray"
)

func onReady() {
	systray.SetIcon(icon.Data)
	systray.SetTitle("MpvCtl")
	systray.SetTooltip("MPVCTL")
	mQuit := systray.AddMenuItem("Quit", "Quit")

	// Sets the icon of a menu item. Only available on Mac and Windows.
	go func() {
		<-mQuit.ClickedCh
		systray.Quit()
	}()
}

func onExit() {
	kill()
	// clean up here
}
