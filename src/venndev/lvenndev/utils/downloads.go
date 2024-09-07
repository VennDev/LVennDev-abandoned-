package utils

import (
	"io"
	"net/http"
	"os"
	"strconv"

	"fyne.io/fyne/v2/widget"
)

func DownloadFile(url, dest string, progressBar *widget.ProgressBar) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()

	totalSize, err := strconv.Atoi(resp.Header.Get("Content-Length"))
	if err != nil {
		return err
	}

	var downloadedSize int64

	buffer := make([]byte, 1024)
	for {
		n, err := resp.Body.Read(buffer)
		if n > 0 {
			_, err := out.Write(buffer[:n])
			if err != nil {
				return err
			}
			downloadedSize += int64(n)
			progressBar.SetValue(float64(downloadedSize) / float64(totalSize))
		}

		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
	}

	return nil
}
