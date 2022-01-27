package gui

import (
	"andy/internal/http_server"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Run() {
	a := app.New()
	w := a.NewWindow("Andy GUI")
	w.Resize(fyne.NewSize(300, 300))

	hello := widget.NewLabel("Commands")
	w.SetContent(
		container.NewVBox(
			hello,
			widget.NewButton("docker_http", func() {
				http_server.RunDockerHttpServer(1991)
			}),
		),
	)

	w.ShowAndRun()
}
