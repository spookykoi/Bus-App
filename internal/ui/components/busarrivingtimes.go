package components

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

const (
	minWidth  = float32(70)
	minHeight = float32(34)
)

type BusArrivalTimes struct {
	widget.BaseWidget

	Arrivals []int
}

func NewBusArrivalTimes(times []int) *BusArrivalTimes {
	b := &BusArrivalTimes{Arrivals: times}
	b.ExtendBaseWidget(b)
	return b
}

func (b *BusArrivalTimes) MinSize() fyne.Size {
	return fyne.NewSize(minWidth, minHeight)
}

// Helper function to set the text of the arrival time.
// Sets to `-` if Arrivals slice doesn't have that index.
func (b *BusArrivalTimes) getArrivalText(index int, prefix string) string {
	if index < len(b.Arrivals) {
		return fmt.Sprintf("%s%d min", prefix, b.Arrivals[index])
	}
	return prefix + "-"
}

// CreateRenderer implements [fyne.Widget].
func (b *BusArrivalTimes) CreateRenderer() fyne.WidgetRenderer {
	top := canvas.NewText(b.getArrivalText(0, ""), color.White)
	bottom := canvas.NewText(b.getArrivalText(1, "Next: "), color.NRGBA{R: 0x5F, G: 0x7C, B: 0xBA, A: 0xFF})

	return &busArrivalTimesRenderer{
		widget: b,
		top:    top,
		bottom: bottom,
	}
}

// Render ---------------------------------------------------------------------

type busArrivalTimesRenderer struct {
	widget *BusArrivalTimes
	top    *canvas.Text
	bottom *canvas.Text
}

// This is the actual vertical layout.
func (r *busArrivalTimesRenderer) Layout(size fyne.Size) {
	topH := r.top.MinSize().Height
	r.top.Move(fyne.NewPos(0, 0))
	r.top.Resize(fyne.NewSize(size.Width, topH))
	r.bottom.Move(fyne.NewPos(0, topH))
	r.bottom.Resize(fyne.NewSize(size.Width, size.Height-topH))
}

func (r *busArrivalTimesRenderer) MinSize() fyne.Size {
	topMin := r.top.MinSize()
	bottomMin := r.bottom.MinSize()
	return fyne.NewSize(fyne.Max(topMin.Width, bottomMin.Width), topMin.Height+bottomMin.Height)
}

func (r *busArrivalTimesRenderer) Refresh() {
	r.top.Text = r.widget.getArrivalText(0, "")
	r.bottom.Text = r.widget.getArrivalText(1, "Next: ")
	canvas.Refresh(r.widget)
}

func (r *busArrivalTimesRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.top, r.bottom}
}

func (r *busArrivalTimesRenderer) Destroy() {}
