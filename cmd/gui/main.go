package main

import (
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	SystemInfo "github.com/hutchisont/system_monitor/internal/system_info"
)

func main() {
	a := app.New()
	w := a.NewWindow("System Monitor")
	s := SystemInfo.SystemInfo{}

	infoWidget := widget.NewLabel(s.String())
	w.SetContent(container.NewVBox(infoWidget))

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	go func() {
		for {
			select {
			case <-ticker.C:
				s.UpdateAllReadings()
				infoWidget.SetText(s.String())
			}
		}
	}()

	w.ShowAndRun()
}
