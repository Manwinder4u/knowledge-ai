package documents

import (
	"errors"
	"mime/multipart"
)

const (
	PDFMimeType = "application/pdf"
)

var (
	ErrFileRequired    = errors.New("file is required")
	ErrInvalidFileType = errors.New("only PDF files are allowed")
	ErrFileTooLarge    = errors.New("file exceeds maximum allowed size")
)

func ValidateUpload(header *multipart.FileHeader, maxUploadSize int64) error {

	if header == nil {
		return ErrFileRequired
	}

	if header.Size > maxUploadSize {
		return ErrFileTooLarge
	}

	if header.Header.Get("Content-Type") != PDFMimeType {
		return ErrInvalidFileType
	}

	return nil
}
