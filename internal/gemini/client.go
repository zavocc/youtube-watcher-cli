package gemini

import (
	"context"
	"log"

	"google.golang.org/genai"
)

func GApiClient(prompt string, id string) string {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal(err)
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
				ThinkingLevel:   genai.ThinkingLevelMinimal,
				IncludeThoughts: false,
			},
		},
	)

	if err != nil {
		log.Fatalln("An error has occurred while performing inference, error log:\n", err.Error())
	}

	return result.Text()
}
