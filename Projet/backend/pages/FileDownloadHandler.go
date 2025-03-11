package pages

import (
	f "Proj48h/functions"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func DownloadFile(w http.ResponseWriter, r *http.Request) {
	// Getting the file ID from the cookie
	fileId := f.GetCookie(w, r, "FileID").Value
	if fileId == "" {
		http.Error(w, "No file ID provided", http.StatusBadRequest)
		f.ErrorPrintln("No file ID provided when trying to download a file")
		return
	}

	file, err := os.Open(fmt.Sprintf("%s/%s.pdf", f.TmpDirPath, fileId))
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		f.ErrorPrintf("File not found when trying to download a file -> %s.pdf\n", fileId)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			f.ErrorPrintln("Error while closing file when trying to download a file")
		}
	}(file)

	// Defining the http headers to let the browser know that the file is a pdf file and that it should be downloaded
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s.pdf", filepath.Base(fileId+".pdf")))

	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "Error while writing file to response", http.StatusInternalServerError)
		f.ErrorPrintln("Error while writing file to response when trying to download a file")
		return
	}
	f.InfoPrintf("File downloaded to %s", fileId)
}
