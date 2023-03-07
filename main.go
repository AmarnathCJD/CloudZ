package main

import (
	"fmt"
	"html/template"
	"main/modules"
	"main/utils"
	"net/http"
)

var bindingPort = utils.GetOutboundPort()

func setupEndpoints() {
	http.HandleFunc("/upload", modules.UploadFile)
	http.HandleFunc("/download", modules.DownloadFileRequest)
	http.HandleFunc("/download/file", modules.ServeFile)
	http.HandleFunc("/list", modules.GetFiles)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		temp := template.Must(template.ParseFiles("index.html"))
		temp.Execute(w, nil)
	})
	http.HandleFunc("/file", func(w http.ResponseWriter, r *http.Request) {
		temp := template.Must(template.ParseFiles("file.html"))
		temp.Execute(w, nil)
	})

	http.HandleFunc("/charge", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Unable to parse form", http.StatusBadRequest)
			return
		}
		card := r.FormValue("card")
		expMo := r.FormValue("expMo")
		expYr := r.FormValue("expYr")
		cvv := r.FormValue("cvv")
		rx := modules.StarReq{}
		if err := rx.StepOne(); err != nil {
			http.Error(w, "Unable to charge card", http.StatusBadRequest)
			return
		}
		if err := rx.StepTwo(); err != nil {
			http.Error(w, "Unable to charge card", http.StatusBadRequest)
			return
		}
		code, declineCode, message, err := rx.StepThree(card, expMo, expYr, cvv)
		if err != nil {
			http.Error(w, "Unable to charge card", http.StatusBadRequest)
			return
		}
		if code == "" {
			utils.Ok(w, map[string]string{
				"success": "true",
				"charged": "4460",
			})
			return
		}
		utils.Ok(w, map[string]string{
			"code":         code,
			"decline_code": declineCode,
			"message":      message,
		})
	})

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
}

func main() {
	utils.Logger.Info("Starting ZTorr server on port " + utils.GetPortFromBinding(bindingPort) + "...")
	setupEndpoints()
	utils.HandleCtrlZ()
	utils.SendPortAndIP(bindingPort)
	if err := http.ListenAndServe(bindingPort, nil); err != nil {
		utils.Logger.Fatal(fmt.Sprintf("Error while starting server: %s", err.Error()))
	}
}
