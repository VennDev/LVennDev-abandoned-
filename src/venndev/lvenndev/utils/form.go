package utils

import (
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func CreateButton(
	name string,
	url string,
	hasDownload bool,
	downloadPath string,
	progressBar *widget.ProgressBar,
	window fyne.Window,
	wg *sync.WaitGroup,
) *widget.Button {
	button := widget.NewButtonWithIcon(name, theme.DownloadIcon(), func() {
		if !hasDownload {
			wg.Add(1)
			go DownloadFileAndRun(url, downloadPath, progressBar, window, wg)
		} else {
			dialog.ShowInformation(name, name+" is already installed!", window)
		}
	})
	return button
}
