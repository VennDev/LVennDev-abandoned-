package main

import (
	"net/url"
	"os"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/venndev/LVennDev/src/venndev/lvenndev/compons"
	"github.com/venndev/LVennDev/src/venndev/lvenndev/theme"
	"github.com/venndev/LVennDev/src/venndev/lvenndev/utils"
)

const (
	title           = "Libraries VennDev"
	imageBackground = "./images/back-ground.jpg"
	icon            = "./images/icon/icon.jpg"
	github          = "https://github.com/VennDev"
	version         = "1.0.0"
	author          = "VennDev"
	email           = "venndev@gmail.com"
	downloadPath    = "./downloads"
)

var (
	menu        *fyne.MainMenu
	hasVSCode   bool = false
	hasGoogle   bool = false
	hasComposer bool = false
	progressBar *widget.ProgressBar
	buttonScale = fyne.NewSize(80, 30)
)

func about(window fyne.Window) {
	dialog.ShowInformation(
		"About",
		`
Version: `+version+`
Author: `+author+`
Email: `+email+`
Github: `+github+`
		`,
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

func getToggles() (
	toggleVSCode *compons.CustomCheck,
	toggleGoogle *compons.CustomCheck,
	toggleComposer *compons.CustomCheck,
) {
	// Check if VSCode is installed
	hasVSCode = utils.CheckVSCode()
	toggleVSCode = compons.NewCustomCheck("VSCode", nil)
	toggleVSCode.SetChecked(hasVSCode)
	toggleVSCode.OnChanged = func(checked bool) {
		if checked {
			toggleVSCode.Enable()
		}
	}

	// Check if Google Chrome is installed
	hasGoogle = utils.CheckGoogleHasDownloaded()
	toggleGoogle = compons.NewCustomCheck("Google Chrome", nil)
	toggleGoogle.SetChecked(hasGoogle)
	toggleGoogle.OnChanged = func(checked bool) {
		if checked {
			toggleGoogle.Enable()
		}
	}

	// Check if Composer is installed
	hasComposer = utils.CheckComposer()
	toggleComposer = compons.NewCustomCheck("Composer", nil)
	toggleComposer.SetChecked(hasComposer)
	toggleComposer.OnChanged = func(checked bool) {
		if checked {
			toggleComposer.Enable()
		}
	}

	return toggleVSCode, toggleGoogle, toggleComposer
}

func getButtons(window fyne.Window, wg *sync.WaitGroup) (ButtonRight *fyne.Container, ButtonLeft *fyne.Container) {
	// Toggles
	toggleVSCode, toggleGoogle, toggleComposer := getToggles()

	buttonVSCode := utils.CreateButton(
		"", utils.VscodeUrl, hasVSCode, downloadPath+"/vscode_installer.exe", progressBar, window, wg,
	)

	buttonGoogle := utils.CreateButton(
		"", utils.ChromeUrl, hasGoogle, downloadPath+"/chrome_installer.exe", progressBar, window, wg,
	)

	buttonComposer := utils.CreateButton(
		"", utils.ComposerUrl, hasComposer, downloadPath+"/composer_installer.exe", progressBar, window, wg,
	)

	buttonContainerRight := container.New(layout.NewGridWrapLayout(buttonScale))
	buttonContainerLeft := container.New(layout.NewGridLayoutWithColumns(1),
		container.NewHBox(
			toggleVSCode,
			layout.NewSpacer(),
			buttonVSCode,
		),
		container.NewHBox(
			toggleGoogle,
			layout.NewSpacer(),
			buttonGoogle,
		),
		container.NewHBox(
			toggleComposer,
			layout.NewSpacer(),
			buttonComposer,
		),
	)

	return buttonContainerRight, buttonContainerLeft
}

func main() {
	wg := sync.WaitGroup{}
	myApp := app.New()
	myApp.Settings().SetTheme(theme.VTheme{})
	myWindow := myApp.NewWindow(title)
	myWindow.Resize(fyne.NewSize(1000, 600))

	// Check if the files exist
	if !checkFiles() {
		return
	}

	// Check path download
	if _, err := os.Stat(downloadPath); os.IsNotExist(err) {
		os.Mkdir(downloadPath, os.ModePerm)
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

	// Search Bar
	searchBar := widget.NewEntry()
	searchBar.SetPlaceHolder("Search...")
	searchBarContainer := container.NewGridWrap(fyne.NewSize(200, 40), searchBar)

	// Hyperlink
	githubUrl, err := url.Parse(github)
	if err != nil {
		dialog.ShowInformation("Error", "Error parsing GitHub URL:"+err.Error(), myWindow)
		return
	}
	githubUrlHyperlink := widget.NewHyperlink("Github: VennDev", githubUrl)
	hyperlinkContainer := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), githubUrlHyperlink)

	// Buttons
	buttonContainerRight, buttonContainerLeft := getButtons(myWindow, &wg)

	// Background
	background := canvas.NewImageFromFile(imageBackground)
	background.FillMode = canvas.ImageFillStretch
	background.SetMinSize(fyne.NewSize(1000, 600))

	// Main Content
	mainContentRight := container.NewBorder(
		nil,
		nil,
		container.NewVBox(
			buttonContainerRight,
		),
		nil,
	)
	mainContentLeft := container.NewBorder(
		nil,
		nil,
		container.NewVBox(
			buttonContainerLeft,
		),
		nil,
	)

	// Main Content
	mainContent := container.NewHSplit(mainContentLeft, mainContentRight)
	mainContent.SetOffset(0.2)

	// Content
	content := container.NewVScroll(
		container.NewBorder(
			searchBarContainer,
			container.NewVBox(progressBar, hyperlinkContainer),
			nil,
			nil,
			background,
			mainContent,
		),
	)

	// Set Content
	myWindow.SetContent(content)

	// Set Main Menu
	myWindow.SetMainMenu(menu)

	// Set Icon
	if iconResource, err := utils.LoadResourceFromPath(icon); err == nil {
		myWindow.SetIcon(iconResource)
	}

	// Show and Run
	myWindow.ShowAndRun()
}
