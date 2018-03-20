package display

type Surface interface {
	MoveTo(x float64, y float64)
	Arc(xc float64, yc float64, radius float64, angle1 float64, angle2 float64)
	DrawRectangle(x float64, y float64, width float64, height float64)
	Fill()
	FillPreserve()
	SetLineWidth(width float64)
	SetRgba(r, g, b, a float64)
	Stroke()

	// Displayable render support
	Push(d Displayable) error
	GetRoot() Displayable

	/*
		NewPath()
		NewSubPath()
		LineTo(x float64, y float64)
		CurveTo(x1 float64, y1 float64, x2 float64, y2 float64, x3 float64, y3 float64)
		Arc(xc float64, yc float64, radius float64, angle1 float64, angle2 float64)
		ArcNegative(xc float64, yc float64, radius float64, angle1 float64, angle2 float64)
		RelMoveTo(dx float64, dy float64)
		RelLineTo(dx float64, dy float64)
		RelCurveTo(dx1 float64, dy1 float64, dx2 float64, dy2 float64, dx3 float64, dy3 float64)
		DrawRectangle(x float64, y float64, width float64, height float64)
		ClosePath()
		PathExtents(x1 *float64, y1 *float64, x2 *float64, y2 *float64)

		// FillPreserve()
		// InStroke(x float64, y float64) bool
		// InFill(x float64, y float64) bool
		// InClip(x float64, y float64) bool
		// StrokeExtents(x1 *float64, y1 *float64, x2 *float64, y2 *float64)
		// FillExtents(x1 *float64, y1 *float64, x2 *float64, y2 *float64)
		// ResetClip()
		// Clip()
		// ClipPreserve()
		// ClipExtents(x1 *float64, y1 *float64, x2 *float64, y2 *float64)

		// SelectFontFace(family string, slant FontSlant, weight FontWeight)
		// SetFontOptions(options *FontOptions)
		// SetFontFace(fontFace *FontFace)
	*/
}