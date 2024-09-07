package main

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Liraries VennDev")
	myWindow.Resize(fyne.NewSize(1000, 500))

	imagePath := "./images/back-ground.jpg"
	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		fmt.Printf("Error: Image file does not exist at path %s\n", imagePath)
		return
	}

	label := widget.NewLabel("This is an image below:")
	label.Resize(fyne.NewSize(100, 50))

	button := widget.NewButton("Click Me", func() {
		// Handle button click
	})
	button.Resize(fyne.NewSize(100, 50))

	buttonContainer := container.NewHBox(
		container.NewCenter(button),
	)

	content := container.NewStack(
		canvas.NewImageFromFile(imagePath),
		container.NewVBox(
			label,
			buttonContainer,
		),
	)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
