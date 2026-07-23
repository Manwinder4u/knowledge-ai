package retrieval

// func TestRetrieve(t *testing.T) {

// 	repo := ...

// 	embedder := embedding.NewOllamaEmbedder(
// 		"http://localhost:11434",
// 		"nomic-embed-text",
// 	)

// 	service := retrieval.NewService(
// 		embedder,
// 		repo,
// 	)

// 	chunks, err := service.Retrieve(
// 		context.Background(),
// 		"What is the outstanding Bajaj loan amount?",
// 		5,
// 	)

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	for _, chunk := range chunks {
// 		t.Log(chunk.Content)
// 	}
// }
