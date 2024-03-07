package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello World")

	// Crie widgets adicionais
	label := widget.NewLabel("Hello World!")
	button := widget.NewButton("Click me!", func() {
		label.SetText("Button clicked!")
	})

	// Organize os widgets em um layout
	content := container.NewVBox(
		label,
		button,
	)

	// Configure a janela com o novo layout
	w.SetContent(content)
	w.ShowAndRun()
}
