package controls

import (
	"ui/comp"
	"ui/opts"
	"ui"
	"views"
)

type LabelComponent struct {
	comp.Component
	measuredText string
	measuredSize int
}

func (l *LabelComponent) SetFontSize(size int) {
	l.measuredSize = 0
	l.Component.SetFontSize(size)
	l.Measure()
}

// Layout the Label by first measuring Text and configuring our min dimensions.
func (l *LabelComponent) SetText(currentText string) {
	l.Component.SetText(currentText)
	l.Measure()
}

func (l *LabelComponent) Measure() {
	face := l.FontFace()
	currentText := l.Text()
	currentSize := l.FontSize()

	shouldUpdate := face != "" &&
		(l.measuredText != currentText || l.measuredSize != currentSize)

	if shouldUpdate {
		font := l.Context().Font(face)
		if font != nil {
			l.measuredText = currentText
			l.measuredSize = l.FontSize()

			// Update the Font Atlas with the current/updated size.
			font.SetSize(float32(l.measuredSize))
			w, bounds := font.Bounds(l.measuredText)
			h := bounds[3] - bounds[1]

			l.SetTextX(float64(bounds[0]))
			l.SetTextY(float64(bounds[1]))

			l.SetMinHeight(float64(h) + l.VerticalPadding())
			l.SetMinWidth(float64(w) + l.HorizontalPadding())
		}
	}
}

func NewLabel() *LabelComponent {
	return &LabelComponent{}
}

// Label is a component with a text title that is rendered over the background.
var Label = comp.Define("Label",
	func() ui.Displayable { return NewLabel() },
	opts.LayoutType(ui.NoLayoutType),
	opts.IsFocusable(true),
	opts.IsText(true),
	opts.View(views.LabelView))