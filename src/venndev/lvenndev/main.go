package main

import (
	"net/url"
	"os"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/venndev/LVennDev/src/venndev/lvenndev/theme"
	"github.com/venndev/LVennDev/src/venndev/lvenndev/utils"
)

const (
	title           = "Libraries VennDev"
	imageBackground = "./images/back-ground.jpg"
	icon            = "./images/icon/icon.ico"
	github          = "https://github.com/VennDev"
	version         = "1.0.0"
	author          = "VennDev"
	email           = "venndev@gmail.com"
)

var (
	menu        *fyne.MainMenu
	hasVSCode   bool = false
	progressBar *widget.ProgressBar
	buttonScale = fyne.NewSize(80, 30)
)

func about(window fyne.Window) {
	dialog.ShowInformation(
		"About",
		"Version: "+version+"\nAuthor: "+author+"\nEmail: "+email+"\nGithub: "+github,
		window,
	)
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
	myApp.Settings().SetTheme(theme.VTheme{})
	myWindow := myApp.NewWindow(title)
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
				about(myWindow)
			}),
		),
	)

	// Check if VSCode is installed
	hasVSCode = utils.CheckVSCode()

	// Buttons
	label := widget.NewLabel("VSCode: " + strconv.FormatBool(hasVSCode))
	label.Alignment = fyne.TextAlignCenter
	label.TextStyle = fyne.TextStyle{Bold: true}
	buttonVSCode := widget.NewButton("Download", func() {
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
	buttonVSCode.Importance = widget.HighImportance
	buttonContainer := container.New(layout.NewGridWrapLayout(buttonScale), buttonVSCode)

	// Hyperlink
	githubUrl, err := url.Parse(github)
	if err != nil {
		dialog.ShowInformation("Error", "Error parsing GitHub URL:"+err.Error(), myWindow)
		return
	}

	githubUrlHyperlink := widget.NewHyperlink("Github: VennDev", githubUrl)
	hyperlinkContainer := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), githubUrlHyperlink)

	// Background
	background := canvas.NewImageFromFile(imageBackground)
	background.FillMode = canvas.ImageFillStretch
	background.SetMinSize(fyne.NewSize(1000, 600))

	// Main Content
	mainContentRight := container.NewHBox(
		label,
		buttonContainer,
	)
	mainContentLeft := container.NewHBox(
		label,
		buttonContainer,
	)

	// Content
	content := container.New(
		layout.NewStackLayout(),
		background,
		container.NewBorder(
			nil,
			hyperlinkContainer,
			mainContentLeft,
			mainContentRight,
		),
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
