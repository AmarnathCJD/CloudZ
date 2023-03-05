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
