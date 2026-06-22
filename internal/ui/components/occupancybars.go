package components

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

const (
	occupancyWidth  = float32(37)
	occupancyHeight = float32(34)

	barsAreaWidth  = float32(22)
	barsAreaHeight = float32(16)
	barGap         = float32(2)
	barCornerRad   = float32(4)
	numOccBars     = 3
)

type OccupancyBars struct {
	widget.BaseWidget

	AvailabilityRate float32
}

func NewOccupancyBars(occupancyRate float32) *OccupancyBars {
	o := &OccupancyBars{AvailabilityRate: occupancyRate}
	o.ExtendBaseWidget(o)
	return o
}

func (o *OccupancyBars) MinSize() fyne.Size {
	return fyne.NewSize(occupancyWidth, occupancyHeight)
}

func (o *OccupancyBars) CreateRenderer() fyne.WidgetRenderer {
	col := occupancyColor(o.AvailabilityRate)

	var bars [numOccBars]*canvas.Rectangle
	for i := range bars {
		r := canvas.NewRectangle(col)
		r.CornerRadius = barCornerRad
		bars[i] = r
	}

	label := canvas.NewText(occupancyLabel(o.AvailabilityRate), col)
	label.TextSize = theme.CaptionTextSize()
	label.Alignment = fyne.TextAlignCenter

	return &occupancyBarsRenderer{widget: o, bars: bars, label: label}
}

// ── helpers ───────────────────────────────────────────────────────────────────

func occupancyColor(rate float32) color.RGBA {
	switch {
	case rate < 0.33:
		return color.RGBA{R: 0xD1, G: 0x5F, B: 0x61, A: 0xFF}
	case rate < 0.67:
		return color.RGBA{R: 0xD1, G: 0xCD, B: 0x5F, A: 0xFF}
	default:
		return color.RGBA{R: 0x5F, G: 0xD1, B: 0x85, A: 0xFF}
	}
}

func occupancyLabel(rate float32) string {
	switch {
	case rate < 0.33:
		return "Standing only"
	case rate < 0.67:
		return "Few seats"
	default:
		return "Free seats"
	}
}

// ── Renderer ──────────────────────────────────────────────────────────────────

type occupancyBarsRenderer struct {
	widget *OccupancyBars
	bars   [numOccBars]*canvas.Rectangle
	label  *canvas.Text
}

func (r *occupancyBarsRenderer) Objects() []fyne.CanvasObject {
	objs := make([]fyne.CanvasObject, numOccBars+1)
	for i, b := range r.bars {
		objs[i] = b
	}
	objs[numOccBars] = r.label
	return objs
}

func (r *occupancyBarsRenderer) Layout(size fyne.Size) {
	barW := (barsAreaWidth - barGap*float32(numOccBars-1)) / float32(numOccBars)
	offsetX := (size.Width - barsAreaWidth) / 2

	for i, bar := range r.bars {
		x := offsetX + float32(i)*(barW+barGap)
		bar.Move(fyne.NewPos(x, 0))
		bar.Resize(fyne.NewSize(barW, barsAreaHeight))
	}

	r.label.Move(fyne.NewPos(0, barsAreaHeight+barGap))
	r.label.Resize(fyne.NewSize(size.Width, size.Height-barsAreaHeight-barGap))
}

func (r *occupancyBarsRenderer) MinSize() fyne.Size {
	return r.widget.MinSize()
}

func (r *occupancyBarsRenderer) Refresh() {
	col := occupancyColor(r.widget.AvailabilityRate)
	for _, bar := range r.bars {
		bar.FillColor = col
		bar.Refresh()
	}
	r.label.Text = occupancyLabel(r.widget.AvailabilityRate)
	r.label.Color = col
	canvas.Refresh(r.widget)
}

func (r *occupancyBarsRenderer) Destroy() {}
