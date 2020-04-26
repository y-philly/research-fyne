package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/theme"
)

type clockLayout struct {
	hour, minute, second     *canvas.Line
	hourDot, secondDot, face *canvas.Circle

	canvas fyne.CanvasObject
	stop   bool
}

func (c *clockLayout) Layout(_ []fyne.CanvasObject, size fyne.Size) {
	diameter := fyne.Min(size.Width, size.Height)
	radius := diameter / 2
	// dotRadius := radius / 12
	// smallDotRadius := dotRadius / 8

	stroke := float32(diameter) / 40
	// midStroke := float32(diameter) / 90
	smallStroke := float32(diameter) / 200

	size = fyne.NewSize(diameter, diameter)
	middle := fyne.NewPos(size.Width/2, size.Height/2)
	topLeft := fyne.NewPos(middle.X-radius, middle.Y-radius)
	c.face.Move(topLeft)

	c.face.Resize(size)

	c.hour.StrokeWidth = stroke
	// c.hour.Resize(fyne.NewSize(radius, radius))
	c.hour.Resize(fyne.NewSize(1, radius))
	c.hour.Move(fyne.NewPos(middle.X-radius/2, middle.Y-radius/2))
	c.face.StrokeWidth = smallStroke
}

func (c *clockLayout) MinSize(_ []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(200, 200)
}

func (c *clockLayout) render() *fyne.Container {
	c.face = &canvas.Circle{StrokeColor: theme.TextColor(), StrokeWidth: 1}
	c.hour = &canvas.Line{StrokeColor: theme.TextColor(), StrokeWidth: 5}
	c.minute = &canvas.Line{StrokeColor: theme.TextColor(), StrokeWidth: 3}
	c.second = &canvas.Line{StrokeColor: theme.PrimaryColor(), StrokeWidth: 1}

	container := fyne.NewContainer(c.face, c.hour, c.minute, c.second)
	container.Layout = c

	c.canvas = container
	return container
}

func main() {
	app := app.New()

	w := app.NewWindow("Hello")
	clock := &clockLayout{}
	w.SetOnClosed(func() {
		clock.stop = true
	})

	content := clock.render()

	w.SetContent(content)
	w.ShowAndRun()
}
