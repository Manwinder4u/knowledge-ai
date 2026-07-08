package documents

type DocumentResponse struct {
	ID               string `json:"id"`
	Title            string `json:"title"`
	OriginalFilename string `json:"original_filename"`

	StoragePath string `json:"storage_path"`
	MimeType    string `json:"mime_type"`
	FileSize    int64  `json:"file_size"`

	Status string `json:"status"`
}

func ToResponse(doc *Document) DocumentResponse {
	return DocumentResponse{
		ID:               doc.ID,
		Title:            doc.Title,
		OriginalFilename: doc.OriginalFilename,
		Status:           doc.Status,
	}
}
