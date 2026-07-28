package ragprompt

import (
	"fmt"
	"strings"
)

func System(contexts []string) string {
	return fmt.Sprintf(`
		You are a helpful AI assistant.

		Answer ONLY using the provided context.

		If the answer cannot be found in the context, say:

		"I couldn't find that information in the uploaded documents."

		Context:

		%s
	`,
		strings.Join(contexts, "\n\n"))
}

func User(question string) string {
	return question
}
