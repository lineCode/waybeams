package controls_test

import (
	"assert"
	"controls"
	"ctx"
	"opts"
	"testing"
	"ui"
)

func createLabel(text string) ui.Displayable {
	return controls.Label(
		ctx.NewTestContext(),
		opts.FontFace("Roboto"),
		opts.FontSize(12),
		opts.Text(text),
	)
}

func TestLabel(t *testing.T) {
	t.Run("Label", func(t *testing.T) {
		label := createLabel("Hello World")
		assert.Equal(t, label.Text(), "Hello World")
		assert.Equal(t, label.Height(), 11, "MinHeight set")
		assert.Equal(t, label.Width(), 51, "MinWidth set")
	})

	t.Run("Metrics change when FontSize changes", func(t *testing.T) {
		label := createLabel("Hello")
		label.SetFontSize(36)
		label.Layout()
		assert.Equal(t, label.Height(), 27, "MinHeight set")
		assert.Equal(t, label.Width(), 70, "MinWidth set")
	})
}