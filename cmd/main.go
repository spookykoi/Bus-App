package main

import (
	"bus-app/internal/ui/components"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello World")
	w.SetContent(container.NewPadded(
		components.NewBusArrivalTimes([]int{3, 14}),
	))
	w.ShowAndRun()
}
