package modules

import (
	"fmt"
	"io"
	"log"
	"main/utils"
	"net/http"
	"os"
	"path/filepath"
)

const DOWNLOAD_PATH = "download/"

func init() {
	if _, err := os.Stat(DOWNLOAD_PATH); os.IsNotExist(err) {
		if err := os.Mkdir(DOWNLOAD_PATH, 0755); err != nil {
			log.Fatal("Error creating download directory: ", err)
		}
	}
}

func UploadFile(w http.ResponseWriter, r *http.Request) {
	// utils.SetAccessControl(w)
	if !utils.IsPost(r) {
		utils.BadRequest(w, utils.ErrNotPost)
		return
	}
	if !utils.IsMultiPartForm(r) {
		utils.BadRequest(w, utils.ErrNotMultipartForm)
		return
	}
	fmt.Println("uploading file")
	r.ParseMultipartForm(32 << 20)
	file, _, err := r.FormFile("file")
	fmt.Println("file uploaded")
	if err != nil {
		utils.ErrorWithPrefix(w, err, http.StatusBadRequest, "could not parse file")
	}
	defer file.Close()
	f, err := os.OpenFile(DOWNLOAD_PATH+"haha", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
	utils.Ok(w, map[string]string{"message": "file uploaded successfully"})
}

func GenFileDownloadKey(filename string) string {
	const prefix = "file_download_"
	default_aes_key := []byte("default_aes_key")
	encrypted, err := utils.AesEncrypt([]byte(filename), default_aes_key)
	if err != nil {
		log.Fatal("Error encrypting filename: ", err)
	}
	return prefix + string(encrypted)
}

func DownloadFileRequest(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Query().Get("filename")
	if filename == "" {
		utils.BadRequest(w, utils.ErrMissingFilename)
		return
	}
	key := GenFileDownloadKey(filename)
	utils.Ok(w, map[string]string{"key": key})
}

func ServeFile(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		utils.BadRequest(w, utils.ErrMissingKey)
		return
	}
	const prefix = "file_download_"
	default_aes_key := []byte("default_aes_key")
	encrypted := []byte(key[len(prefix):])
	decrypted, err := utils.AesDecrypt(encrypted, default_aes_key)
	if err != nil {
		utils.BadRequest(w, utils.ErrInvalidKey)
		return
	}
	filename := string(decrypted)
	f, err := os.Open(DOWNLOAD_PATH + filename)
	if err != nil {
		utils.NotFound(w, utils.ErrFileNotFound)
		return
	}
	defer f.Close()
	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	io.Copy(w, f)
}

type File struct {
	Filename     string `json:"filename"`
	FileSize     int64  `json:"filesize"`
	FileExt      string `json:"fileext"`
	DateModified string `json:"datemodified"`
}

func GetFiles(w http.ResponseWriter, r *http.Request) {
	files, err := os.ReadDir(DOWNLOAD_PATH)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}
	var file_list []File
	for _, file := range files {
		f, err := file.Info()
		if err != nil {
			continue
		}
		file_list = append(file_list, File{
			Filename:     file.Name(),
			FileSize:     f.Size(),
			FileExt:      filepath.Ext(f.Name())[1:],
			DateModified: f.ModTime().String(),
		})
	}
	utils.Ok(w, file_list)
}