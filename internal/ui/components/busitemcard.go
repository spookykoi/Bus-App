package components

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

const (
	itemCardPadding   = float32(10)
	itemCardGap       = float32(8)
	itemCardCornerRad = float32(8)
)

type BusItemCard struct {
	widget.BaseWidget

	Item BusItem
}

func NewBusItemCard(item BusItem) *BusItemCard {
	c := &BusItemCard{Item: item}
	c.ExtendBaseWidget(c)
	return c
}

func (c *BusItemCard) CreateRenderer() fyne.WidgetRenderer {
	bg := canvas.NewRectangle(color.RGBA{R: 0x20, G: 0x36, B: 0x61, A: 0xFF})
	bg.CornerRadius = itemCardCornerRad

	badge := NewBusNumberBadge(c.Item.BusNumber)
	arrivals := NewBusArrivalTimes(c.Item.ArrivalTimes)
	occ := NewOccupancyBars(c.Item.OccupancyRate)

	return &busItemCardRenderer{
		card:     c,
		bg:       bg,
		badge:    badge,
		arrivals: arrivals,
		occ:      occ,
	}
}

// ── Renderer ──────────────────────────────────────────────────────────────────

type busItemCardRenderer struct {
	card     *BusItemCard
	bg       *canvas.Rectangle
	badge    *BusNumberBadge
	arrivals *BusArrivalTimes
	occ      *OccupancyBars
}

func (r *busItemCardRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.bg, r.badge, r.arrivals, r.occ}
}

func (r *busItemCardRenderer) Layout(size fyne.Size) {
	r.bg.Move(fyne.NewPos(0, 0))
	r.bg.Resize(size)

	badgeMin := r.badge.MinSize()
	occMin := r.occ.MinSize()
	innerW := size.Width - 2*itemCardPadding
	arrivalsW := innerW - badgeMin.Width - occMin.Width - 2*itemCardGap
	rowH := size.Height - 2*itemCardPadding

	r.badge.Move(fyne.NewPos(itemCardPadding, itemCardPadding))
	r.badge.Resize(badgeMin)

	r.arrivals.Move(fyne.NewPos(itemCardPadding+badgeMin.Width+itemCardGap, itemCardPadding))
	r.arrivals.Resize(fyne.NewSize(arrivalsW, rowH))

	r.occ.Move(fyne.NewPos(size.Width-itemCardPadding-occMin.Width, itemCardPadding))
	r.occ.Resize(occMin)
}

func (r *busItemCardRenderer) MinSize() fyne.Size {
	badgeMin := r.badge.MinSize()
	arrMin := r.arrivals.MinSize()
	occMin := r.occ.MinSize()

	rowH := fyne.Max(fyne.Max(badgeMin.Height, arrMin.Height), occMin.Height)
	contentW := badgeMin.Width + itemCardGap + arrMin.Width + itemCardGap + occMin.Width

	return fyne.NewSize(contentW+2*itemCardPadding, rowH+2*itemCardPadding)
}

func (r *busItemCardRenderer) Refresh() {
	r.badge.Number = r.card.Item.BusNumber
	r.badge.Refresh()
	r.arrivals.Arrivals = r.card.Item.ArrivalTimes
	r.arrivals.Refresh()
	r.occ.OccupancyRate = r.card.Item.OccupancyRate
	r.occ.Refresh()
	canvas.Refresh(r.card)
}

func (r *busItemCardRenderer) Destroy() {}
