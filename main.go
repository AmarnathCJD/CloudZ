package main

import (
	"html/template"
	"main/modules"
	"main/utils"
	"net/http"
)

func setupEndpoints() {
	http.HandleFunc("/upload", modules.UploadFile)
	http.HandleFunc("/download", modules.DownloadFileRequest)
	http.HandleFunc("/download/file", modules.ServeFile)
	http.HandleFunc("/list", modules.GetFiles)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		temp := template.Must(template.ParseFiles("index.html"))
		temp.Execute(w, nil)
	})

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
}

func main() {
	setupEndpoints()
	utils.HandleCtrlZ()
	utils.SendPortAndIP(8080)
	http.ListenAndServe(":8080", nil)
}
