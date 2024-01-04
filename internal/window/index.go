package window

import (
	"andy/internal/docker_server"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Run() {
	a := app.New()
	w := a.NewWindow("Andy GUI")
	icon, _ := fyne.LoadResourceFromURLString("https://svgsilh.com/svg/1300187.svg")

	a.SetIcon(icon)
	w.Resize(fyne.NewSize(300, 150))

	w.SetContent(
		container.NewVBox(
			widget.NewButton("docker_http", func() {
				docker_server.RunDockerHttpServer(1991)
			}),
		),
	)

	w.ShowAndRun()
}
