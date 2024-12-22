package clock

import (
	"fyne.io/fyne/v2/widget"
	"time"
)

func SetupClock() *widget.Label {
	clock := widget.NewLabel("")
	updateTime(clock)

	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()

	return clock
}

func updateTime(clock *widget.Label) {
	formatted := time.Now().Format("Time: 03:04:05")
	clock.SetText(formatted)
}
