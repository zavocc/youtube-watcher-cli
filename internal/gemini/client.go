package gemini

import (
	"context"
	"fmt"
	"os"

	"google.golang.org/genai"
)

func GApiClient(prompt string, id string) string {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, "An error has occurred while initializing the Gemini API client, error log:\n", err.Error())
		os.Exit(1)
	}

	contents := []*genai.Content{
		genai.NewContentFromParts([]*genai.Part{
			genai.NewPartFromText(prompt),
			genai.NewPartFromURI("https://www.youtube.com/watch?v="+id, "video/mp4"),
		}, genai.RoleUser),
	}

	result, err := client.Models.GenerateContent(
		ctx, defaultModel,
		contents,
		&genai.GenerateContentConfig{
			SystemInstruction: genai.NewContentFromText(systemPrompt, genai.RoleUser),
			ThinkingConfig: &genai.ThinkingConfig{
				ThinkingBudget:  &thinkingBudget,
				IncludeThoughts: false,
			},
		},
	)

	if err != nil {
		fmt.Fprintln(os.Stderr, "An error has occurred while performing inference, error log:\n", err.Error())
		os.Exit(1)
	}

	return result.Text()
}
