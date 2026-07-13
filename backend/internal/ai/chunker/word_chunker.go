package chunker

import "strings"

type WordChunker struct {
	chunkSize int
	overlap   int
}

func NewWordChunker(chunkSize int, overlap int) *WordChunker {
	return &WordChunker{
		chunkSize: chunkSize,
		overlap:   overlap,
	}
}

func (c *WordChunker) Chunk(text string) []string {

	words := strings.Fields(text)

	if len(words) == 0 {
		return nil
	}

	var chunks []string

	step := c.chunkSize - c.overlap

	if step <= 0 {
		step = c.chunkSize
	}

	for start := 0; start < len(words); start += step {
		end := start + c.chunkSize
		if end > len(words) {
			end = len(words)
		}

		chunk := strings.Join(words[start:end], " ")
		chunks = append(chunks, chunk)

		if end == len(words) {
			break
		}
	}

	return chunks
}
