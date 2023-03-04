package utils

var (
	// Icons
	ICONS = map[string]string{
		"xls":    "https://img.icons8.com/fluency/96/null/microsoft-excel-2019.png",
		"xlsx":   "https://img.icons8.com/fluency/96/null/microsoft-excel-2019.png",
		"go":     "https://img.icons8.com/color/144/null/golang.png",
		"py":     "https://img.icons8.com/color/144/null/python--v1.png",
		"mp4":    "https://img.icons8.com/external-bearicons-flat-bearicons/64/null/external-MP4-file-extension-bearicons-flat-bearicons.png",
		"mp3":    "https://img.icons8.com/external-bearicons-flat-bearicons/64/null/external-MP3-file-extension-bearicons-flat-bearicons.png",
		"pdf":    "https://img.icons8.com/external-bearicons-flat-bearicons/64/null/external-pdf-file-extension-bearicons-flat-bearicons.png",
		"html":   "https://img.icons8.com/parakeet/96/null/html-filetype.png",
		"mkv":    "https://img.icons8.com/color/96/null/mov.png",
		"mov":    "https://img.icons8.com/color/96/null/mov.png",
		"folder": "https://img.icons8.com/emoji/96/null/open-file-folder-emoji.png",
	}
)

func GetIcon(ext string, isdir bool) string {
	if isdir {
		return ICONS["folder"]
	}
	if icon, ok := ICONS[ext]; ok {
		return icon
	}
	return "https://img.icons8.com/ios-glyphs/60/null/file--v1.png"
}
