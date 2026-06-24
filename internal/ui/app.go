package ui

import (
	"bus-app/internal/ui/components"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func Run() {
	a := app.New()
	w := a.NewWindow("Bus App")
	w.SetContent(buildContent())
	w.ShowAndRun()
}

func buildContent() fyne.CanvasObject {
	return container.NewPadded(
		components.NewBusCard("Example bus stop name", []components.BusItem{
			{BusNumber: 123, ArrivalTimes: []int{3, 5, 14}, OccupancyRate: 0.33},
			{BusNumber: 12, ArrivalTimes: []int{2, 8}, OccupancyRate: 0.8},
			{BusNumber: 31, ArrivalTimes: []int{5}, OccupancyRate: 0.2},
			{BusNumber: 985, ArrivalTimes: []int{}, OccupancyRate: 0},
		}),
	)
}
