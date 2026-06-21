package components

import (
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// Custom widget; Refer to the fyne documentation here:
// https://docs.fyne.io/extend/custom-widget/

const (
	badgeWidth  = float32(40)
	badgeHeight = float32(34)
)

// Fyne custom widgets must inherit/compose of the widget.BaseWidget.
type BusNumberBadge struct {
	widget.BaseWidget

	Number int
}

// This is the standard convention for constructors in Go, despite Go not having actual constructors
func NewBusNumberBadge(number int) *BusNumberBadge {
	b := &BusNumberBadge{Number: number}
	b.ExtendBaseWidget(b) // required: registers this as a widget
	return b
}

// MinSize tells Fyne the smallest this widget can be.
// Total width = content width + left/right padding
// Total height = text height + top/bottom padding
func (b *BusNumberBadge) MinSize() fyne.Size {
	// textHeight := theme.TextSize()
	// return fyne.NewSize(badgeWidth+badgePadLeft+badgePadRight, textHeight+badgePadTop+badgePadBot)
	return fyne.NewSize(badgeWidth, badgeHeight)
}

// CreateRenderer returns the renderer, which owns the raw canvas objects.
func (b *BusNumberBadge) CreateRenderer() fyne.WidgetRenderer {
	bg := canvas.NewRectangle(color.RGBA{R: 0xFF, G: 0x00, B: 0x5D, A: 0xFF})
	bg.CornerRadius = 4

	text := canvas.NewText(strconv.Itoa(b.Number), color.White)
	text.Alignment = fyne.TextAlignCenter
	text.TextSize = theme.TextSize()

	return &busNumberBadgeRenderer{badge: b, bg: bg, text: text}
}

// ── Renderer ────────────────────────────────────────────────────────────────

// The renderer must implement the Layout(), MinSize(), Refresh(), Objects(), and Destroy() methods.
type busNumberBadgeRenderer struct {
	// Also reference the Widget it is rendering
	badge *BusNumberBadge
	bg    *canvas.Rectangle
	text  *canvas.Text
}

// Objects returns every canvas object that makes up this widget.
// Order matters: objects drawn first appear behind later ones.
func (r *busNumberBadgeRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.bg, r.text}
}

// Layout is called whenever the widget is resized.
// This is where you position and size each canvas object.
func (r *busNumberBadgeRenderer) Layout(size fyne.Size) {
	// background fills the whole widget
	r.bg.Resize(size)
	r.bg.Move(fyne.NewPos(0, 0))

	// just move to center
	r.text.Move(fyne.NewPos(0, 0))
	r.text.Resize(fyne.NewSize(
		size.Width,
		size.Height,
	))
}

func (r *busNumberBadgeRenderer) MinSize() fyne.Size {
	return r.badge.MinSize()
}

// Refresh is called when the widget's data changes.
func (r *busNumberBadgeRenderer) Refresh() {
	r.text.Text = strconv.Itoa(r.badge.Number)
	canvas.Refresh(r.badge)
}

func (r *busNumberBadgeRenderer) Destroy() {}
