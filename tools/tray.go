package tools

import (
	"github.com/getlantern/systray"
)

func onReady() {
	systray.SetIcon(icon)
	systray.SetTitle("Immersive")

	systray.SetTooltip("Immersive")
	menuQuit := systray.AddMenuItem("Quit", "Quit the whole app")

	go func() {
		<-menuQuit.ClickedCh
		systray.Quit()
	}()
}

func onExit() {
	Error.Fatal("EXITING MANUALLY")
}

func TrayStart() {
	systray.Run(onReady, onExit)
}
