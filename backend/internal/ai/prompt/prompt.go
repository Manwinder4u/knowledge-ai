package prompt

import "strings"

func Build(question string, contexts []string) string {

	var builder strings.Builder

	builder.WriteString(`
		You are a helpful AI assistant.

		Answer ONLY using the provided context.

		If the answer cannot be found in the context, say:

		"I couldn't find that information in the uploaded documents."

		Context:

	`)

	for _, context := range contexts {

		builder.WriteString(context)
		builder.WriteString("\n\n")
	}

	builder.WriteString("Question:\n")
	builder.WriteString(question)

	builder.WriteString("\n\nAnswer:")

	return builder.String()
}
