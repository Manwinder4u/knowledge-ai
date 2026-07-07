package documents

type CreateDocumentRequest struct {
	Title            string `json:"title" validate:"required,min=3,max=255"`
	OriginalFilename string `json:"original_filename" validate:"required"`
}

type DocumentResponse struct {
	ID               string `json:"id"`
	Title            string `json:"title"`
	OriginalFilename string `json:"original_filename"`
	Status           string `json:"status"`
}

func ToResponse(doc *Document) DocumentResponse {
	return DocumentResponse{
		ID:               doc.ID,
		Title:            doc.Title,
		OriginalFilename: doc.OriginalFilename,
		Status:           doc.Status,
	}
}
