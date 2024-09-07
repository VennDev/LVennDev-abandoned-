package main

import (
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"github.com/venndev/LVennDev/src/venndev/lvenndev/utils"
)

const (
	imageBackground = "./images/back-ground.jpg"
	icon            = "./images/icon/icon.ico"
	github          = "https://github.com/VennDev"
)

var (
	menu        *fyne.MainMenu
	hasVSCode   bool = false
	progressBar *widget.ProgressBar
)

func checkVSCode() bool {
	_, err := exec.LookPath("code")
	if err != nil {
		fmt.Println("VSCode is not installed")
		return false
	}
	return true
}

func checkFiles() bool {
	// Check if the image background file exists
	if _, err := os.Stat(imageBackground); os.IsNotExist(err) {
		dialog.ShowInformation("Error", "Image Background file does not exist at path "+imageBackground, nil)
		return false
	}

	// Check if the icon file exists
	if _, err := os.Stat(icon); os.IsNotExist(err) {
		dialog.ShowInformation("Error", "Icon file does not exist at path "+icon, nil)
		return false
	}

	return true
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Libraries VennDev")
	myWindow.Resize(fyne.NewSize(1000, 600))

	// Check if the files exist
	if !checkFiles() {
		return
	}

	// Progress Bar
	progressBar = widget.NewProgressBar()

	// Main Menu
	menu = fyne.NewMainMenu(
		fyne.NewMenu("File",
			fyne.NewMenuItem("About", func() {
				dialog.ShowInformation(
					"About",
					"Version: 1.0.0\nAuthor: VennDev\nEmail: venndev@gmail.com",
					myWindow,
				)
			}),
		),
	)

	// Check if VSCode is installed
	hasVSCode = checkVSCode()

	// Buttons
	label := widget.NewLabel("VSCode: " + strconv.FormatBool(hasVSCode))
	buttonVSCode := widget.NewButton("Click to download!", func() {
		if !hasVSCode {
			urlDownload := "https://code.visualstudio.com/sha/download?build=stable&os=win32-x64-user"
			err := utils.DownloadFile(urlDownload, "vscode_installer.exe", progressBar)
			if err != nil {
				dialog.ShowError(err, myWindow)
			}
		} else {
			dialog.ShowInformation("VSCode", "VSCode is already installed!", myWindow)
		}
	})

	buttonContainer := container.New(layout.NewGridWrapLayout(fyne.NewSize(150, 50)), buttonVSCode)

	// Hyperlink
	githubUrl, err := url.Parse(github)
	if err != nil {
		fmt.Println("Error parsing GitHub URL:", err)
		return
	}

	githubUrlHyperlink := widget.NewHyperlink("Github: VennDev", githubUrl)
	hyperlinkContainer := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), githubUrlHyperlink)

	// Background
	background := canvas.NewImageFromFile(imageBackground)
	background.FillMode = canvas.ImageFillStretch

	// Main Content
	mainContent := container.NewVBox(
		label,
		buttonContainer,
	)

	// Content
	content := container.NewBorder(nil, hyperlinkContainer, nil, nil,
		container.NewStack(background, mainContent),
	)

	// Set Content
	myWindow.SetContent(content)

	// Set Main Menu
	myWindow.SetMainMenu(menu)

	// Set Icon
	myWindow.SetIcon(canvas.NewImageFromFile(icon).Resource)

	// Show and Run
	myWindow.ShowAndRun()
}
