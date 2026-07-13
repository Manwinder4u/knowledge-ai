package chunker

type Chunker interface {
	Chunk(text string) []string
}
