package llm

type Role string

const (
	SystemRole    Role = "system"
	UserRole      Role = "user"
	AssistantRole Role = "assistant"
)

type Message struct {
	Role    Role
	Content string
}
