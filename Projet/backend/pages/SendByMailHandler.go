package pages

import (
	f "Proj48h/functions"
	"encoding/json"
	"fmt"
	"net/http"
)

// EmailRequest is a struct to decode the JSON request
type EmailRequest struct {
	Email string `json:"email"`
}

func SendByMail(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var emailRequest EmailRequest
		// Décoder le corps de la requête JSON
		err := json.NewDecoder(r.Body).Decode(&emailRequest)
		if err != nil {
			http.Error(w, "Error while decoding JSON", http.StatusBadRequest)
			f.ErrorPrintln("Error while decoding JSON when trying to send a file by mail")
			return
		}

		// Getting the file ID from the cookie
		fileId := f.GetCookie(w, r, "FileID").Value
		if fileId == "" {
			http.Error(w, "No file ID provided", http.StatusBadRequest)
			f.ErrorPrintln("No file ID provided when trying to download a file")
			f.SendMail(emailRequest.Email, "An error occurred", "An error occurred while trying to send the file by mail")
			return
		}

		f.SendMailWithAttachments(emailRequest.Email, "Your file", "Here is your file", fmt.Sprintf("%s/%s.pdf", f.TmpDirPath, fileId))
		f.InfoPrintf("File %s.pdf was sent by mail to %s", fileId, emailRequest.Email)
	}
}
