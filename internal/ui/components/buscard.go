package components

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

const (
	cardPadding   = float32(16)
	cardGap       = float32(4)
	cardCornerRad = float32(16)
	cardTitleSize = float32(20)
)

type BusItem struct {
	BusNumber     int
	ArrivalTimes  []int
	OccupancyRate float32
}

type BusCard struct {
	widget.BaseWidget

	StopName string
	Items    []BusItem
}

func NewBusCard(stopName string, items []BusItem) *BusCard {
	c := &BusCard{StopName: stopName, Items: items}
	c.ExtendBaseWidget(c)
	return c
}

func (c *BusCard) CreateRenderer() fyne.WidgetRenderer {
	bg := canvas.NewRectangle(color.RGBA{R: 0x13, G: 0x20, B: 0x39, A: 0xFF})
	bg.CornerRadius = cardCornerRad

	title := canvas.NewText(c.StopName, color.White)
	title.TextSize = cardTitleSize
	title.TextStyle = fyne.TextStyle{Bold: true}

	rows := make([]*BusItemCard, len(c.Items))
	for i, item := range c.Items {
		rows[i] = NewBusItemCard(item)
	}

	return &busCardRenderer{card: c, bg: bg, title: title, rows: rows}
}

// ── Renderer ──────────────────────────────────────────────────────────────────

type busCardRenderer struct {
	card  *BusCard
	bg    *canvas.Rectangle
	title *canvas.Text
	rows  []*BusItemCard
}

func (r *busCardRenderer) Objects() []fyne.CanvasObject {
	objs := []fyne.CanvasObject{r.bg, r.title}
	for _, row := range r.rows {
		objs = append(objs, row)
	}
	return objs
}

func (r *busCardRenderer) Layout(size fyne.Size) {
	r.bg.Move(fyne.NewPos(0, 0))
	r.bg.Resize(size)

	innerW := size.Width - 2*cardPadding
	titleH := r.title.MinSize().Height
	r.title.Move(fyne.NewPos(cardPadding, cardPadding))
	r.title.Resize(fyne.NewSize(innerW, titleH))

	y := cardPadding + titleH + cardGap
	for _, row := range r.rows {
		rowH := row.MinSize().Height
		row.Move(fyne.NewPos(cardPadding, y))
		row.Resize(fyne.NewSize(innerW, rowH))
		y += rowH + cardGap
	}
}

func (r *busCardRenderer) MinSize() fyne.Size {
	titleMin := r.title.MinSize()

	var itemsW, itemsH float32
	for i, row := range r.rows {
		rowMin := row.MinSize()
		if rowMin.Width > itemsW {
			itemsW = rowMin.Width
		}
		itemsH += rowMin.Height
		if i > 0 {
			itemsH += cardGap
		}
	}

	contentW := fyne.Max(titleMin.Width, itemsW)
	contentH := titleMin.Height
	if len(r.rows) > 0 {
		contentH += cardGap + itemsH
	}

	return fyne.NewSize(contentW+2*cardPadding, contentH+2*cardPadding)
}

func (r *busCardRenderer) Refresh() {
	r.title.Text = r.card.StopName
	for i, item := range r.card.Items {
		if i < len(r.rows) {
			r.rows[i].Item = item
			r.rows[i].Refresh()
		}
	}
	canvas.Refresh(r.card)
}

func (r *busCardRenderer) Destroy() {}
