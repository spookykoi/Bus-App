package components

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type BusItem struct {
	BusNumber     int
	ArrivalTimes []int
	Occupancy     string
}

func NewBusCard() fyne.CanvasObject {
	exampleItem := BusItem{123, []int{3, 14}, "Free seats"}

	content := container.NewHBox(
		container.NewCenter(NewBusNumberBadge(exampleItem.BusNumber)),
		container.NewCenter(container.NewVBox(
			widget.NewLabel(strconv.Itoa(exampleItem.ArrivalTimes[0])),
			widget.NewLabel(strconv.Itoa(exampleItem.ArrivalTimes[1])),
		)),
		container.NewCenter(widget.NewLabel(exampleItem.Occupancy)),
	)
	return widget.NewCard("", "", content)
}
