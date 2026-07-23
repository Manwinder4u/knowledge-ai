package chat

type AskResponse struct {
	Answer string `json:"answer"`
}

type AskRequest struct {
	Question string `json:"question"`
}
