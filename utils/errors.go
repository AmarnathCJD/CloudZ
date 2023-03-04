package utils

import (
	"github.com/pkg/errors"
)

var (
	ErrNotMultipartForm = errors.New("not a multipart form, please use multipart/form-data")
	ErrMissingFilename  = errors.New("missing filename, please provide the 'filename' query parameter")
	ErrNotPost          = errors.New("method not allowed, please use POST method")

	ErrMissingKey   = errors.New("missing key, please provide the 'key' query parameter")
	ErrInvalidKey   = errors.New("invalid key, please provide a valid key")
	ErrFileNotFound = errors.New("file not found, has been deleted or expired")
)
