package extractor

import (
	"context"
	"fmt"
	"strings"

	"github.com/ledongthuc/pdf"
)

type PDFExtractor struct{}

func NewPDFExtractor() *PDFExtractor {
	return &PDFExtractor{}
}

func (e *PDFExtractor) Extract(ctx context.Context, path string) (string, error) {

	f, reader, err := pdf.Open(path)
	if err != nil {
		return "", fmt.Errorf("open pdf: %w", err)
	}
	defer f.Close()

	var builder strings.Builder

	totalPages := reader.NumPage()

	for pageNumber := 1; pageNumber <= totalPages; pageNumber++ {
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		default:
		}

		page := reader.Page(pageNumber)
		if page.V.IsNull() {
			continue
		}

		text, err := page.GetPlainText(nil)
		if err != nil {
			return "", fmt.Errorf("extract page %d: %w", pageNumber, err)
		}
		builder.WriteString(text)
		builder.WriteString("\n")
	}

	return builder.String(), nil
}
