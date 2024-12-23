package weather

import (
	"fyne.io/fyne/v2/widget"
	"time"
)

func SetupWeather(apiKey, location string, weatherLabel *widget.Label) *widget.Label {
	weather := FetchWeather(apiKey, location)
	weatherLabel.SetText(FormatWeather(weather))

	WriteLog("weather_logs/history.csv", weather)

	go func() {
		for range time.Tick(10 * time.Minute) {
			weather := FetchWeather(apiKey, location)
			weatherLabel.SetText(FormatWeather(weather))
		}
	}()

	return weatherLabel
}
