package compons

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type CustomCheck struct {
	widget.Check
	disabledColor color.Color
}

func NewCustomCheck(label string, changed func(bool)) *CustomCheck {
	check := &CustomCheck{
		Check:         *widget.NewCheck(label, changed),
		disabledColor: color.RGBA{R: 247, G: 229, B: 173, A: 255},
	}
	check.ExtendBaseWidget(check)
	return check
}

func (c *CustomCheck) CreateRenderer() fyne.WidgetRenderer {
	renderer := c.Check.CreateRenderer()
	originalBackground := color.Transparent

	return &customCheckRenderer{
		WidgetRenderer:     renderer,
		check:              c,
		originalBackground: originalBackground,
	}
}

type customCheckRenderer struct {
	fyne.WidgetRenderer
	check              *CustomCheck
	originalBackground color.Color
}

func (r *customCheckRenderer) BackgroundColor() color.Color {
	if !r.check.Disabled() {
		return r.originalBackground
	}
	return r.check.disabledColor
}
