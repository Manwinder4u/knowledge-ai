package documents

import "time"

type Document struct {
	ID               string `json:"id"`
	UserID           string `json:"user_id"`
	Title            string `json:"title"`
	OriginalFilename string `json:"original_filename"`

	StoragePath string `json:"storage_path"`
	MimeType    string `json:"mime_type"`
	FileSize    int64  `json:"file_size"`

	Status string `json:"status"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
