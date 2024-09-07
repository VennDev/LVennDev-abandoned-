package utils

import (
	"io"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/hashicorp/go-getter"
)

type ProgressTracker struct {
	ProgressBar *widget.ProgressBar
	totalSize   int64
	downloaded  int64
}

func (pt *ProgressTracker) TrackProgress(src string, currentSize, totalSize int64, stream io.ReadCloser) io.ReadCloser {
	pt.totalSize = totalSize
	return &progressReader{ReadCloser: stream, tracker: pt}
}

type progressReader struct {
	io.ReadCloser
	tracker *ProgressTracker
}

func (pr *progressReader) Read(p []byte) (int, error) {
	n, err := pr.ReadCloser.Read(p)
	if n > 0 {
		pr.tracker.downloaded += int64(n)
		pr.tracker.ProgressBar.SetValue(float64(pr.tracker.downloaded) / float64(pr.tracker.totalSize))
	}
	return n, err
}

func DownloadFile(url string, dst string, progressBar *widget.ProgressBar, myWindow fyne.Window, wg *sync.WaitGroup) {
	tracker := &ProgressTracker{ProgressBar: progressBar}
	client := &getter.Client{
		Src:              url,
		Dst:              dst,
		Mode:             getter.ClientModeFile,
		ProgressListener: tracker,
	}
	err := client.Get()
	if err != nil {
		dialog.ShowError(err, myWindow)
	}
	wg.Done()
}

func DownloadFileAndRun(url string, dst string, progressBar *widget.ProgressBar, myWindow fyne.Window, wg *sync.WaitGroup) {
	DownloadFile(url, dst, progressBar, myWindow, wg)
	wg.Wait()
	RunCommand(dst)
}
