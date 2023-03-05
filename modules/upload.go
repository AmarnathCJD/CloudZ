package modules

import (
	"encoding/base64"
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

	r.ParseMultipartForm(32 << 20)
	file, handle, err := r.FormFile("file")
	if err != nil {
		utils.ErrorWithPrefix(w, err, http.StatusBadRequest, "could not parse file")
	}
	defer file.Close()
	f, err := os.OpenFile(DOWNLOAD_PATH+handle.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
	utils.Ok(w, map[string]string{"message": "file uploaded successfully", "key": GenFileDownloadKey(handle.Filename)})
}

func GenFileDownloadKey(filename string) string {
	const prefix = "file_download_"
	default_aes_key := "12345678901234567890123456789012"
	encrypted, err := utils.EncryptAES([]byte(filename), default_aes_key)
	if err != nil {
		log.Fatal("Error encrypting filename: ", err)
	}
	return prefix + base64.RawURLEncoding.EncodeToString(encrypted)
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

	default_aes_key := "12345678901234567890123456789012"
	encrypted, _ := base64.RawURLEncoding.DecodeString(key[len(prefix):])

	decrypted, err := utils.DecryptAES(encrypted, default_aes_key)
	if err != nil {
		utils.BadRequest(w, utils.ErrInvalidKey)
		return
	}
	filename := string(decrypted)
	if r.URL.Query().Get("stream") != "" {
		http.ServeFile(w, r, DOWNLOAD_PATH+filename)
		return
	}
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
	FileSize     string `json:"filesize"`
	FileExt      string `json:"fileext"`
	DateModified string `json:"datemodified"`
	Icon         string `json:"icon"`
	IsDir        bool   `json:"isdir"`
}

func GetFiles(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	path := q.Get("path")
	files, err := os.ReadDir(DOWNLOAD_PATH + path)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}
	if len(files) == 0 {
		utils.Ok(w, []File{})
		return
	}
	var file_list []File
	for _, file := range files {
		f, err := file.Info()
		if err != nil {
			continue
		}
		ext := "Folder"
		size := f.Size()
		if !file.IsDir() {
			ext = filepath.Ext(f.Name())[1:]
		} else {
			size, _ = utils.GetFolderSize(DOWNLOAD_PATH + f.Name())
		}
		file_list = append(file_list, File{
			Filename:     file.Name(),
			FileSize:     utils.BytesToSize(size),
			FileExt:      ext,
			DateModified: f.ModTime().String(),
			Icon:         utils.GetIcon(ext, f.IsDir()),
		})
		if file.IsDir() {
			file_list[len(file_list)-1].IsDir = true
		}
	}
	utils.Ok(w, file_list)
}
