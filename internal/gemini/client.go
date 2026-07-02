package gemini

import (
	"context"
	"fmt"
	"strings"

	"github.com/zavocc/youtube-watcher-cli/internal/config"
	"google.golang.org/genai"
)

func GApiClient(ctx context.Context, prompt string, url string, model string, resolution string) (string, error) {
	client, err := initGeminiClient(ctx)
	if err != nil {
		return "", fmt.Errorf("initialize the Gemini API client: %w", err)
	}

	// Check if it's either 11-character YouTube video ID or a full URL
	var actualUrl string
	if len(url) == 11 {
		actualUrl = "https://www.youtube.com/watch?v=" + url
	} else if len(url) > 11 && (url[:7] == "http://" || url[:8] == "https://") {
		actualUrl = strings.Split(url, "&")[0] // Remove any additional query parameters after the video ID
	} else {
		return "", fmt.Errorf("invalid YouTube video ID or URL specified")
	}

	// Validate model and get corresponding config
	modelSelectedConfig, err := config.ValidateModels(model)
	if err != nil {
		return "", err
	}

	mediaResLevel, err := config.ParseMediaResolution(resolution)
	if err != nil {
		return "", err
	}

	// per part resolution
	// videoPart := genai.NewPartFromURI(actualUrl, "video/mp4")
	// videoPart.MediaResolution = &genai.PartMediaResolution{
	//	Level: config.ParseMediaResolution(resolution),
	// }

	contents := []*genai.Content{
		genai.NewContentFromParts([]*genai.Part{
			genai.NewPartFromURI(actualUrl, "video/mp4"),
			genai.NewPartFromText(prompt),
		}, genai.RoleUser),
	}

	var resSchema = genResponseSchema()
	result, err := client.Models.GenerateContent(
		ctx, modelSelectedConfig.ModelID,
		contents,
		&genai.GenerateContentConfig{
			SystemInstruction: genai.NewContentFromText(systemPrompt, genai.RoleUser),
			ThinkingConfig:    modelSelectedConfig.ThinkingConfig,
			MediaResolution:   genai.MediaResolution(mediaResLevel),
			ResponseMIMEType:  "application/json",
			ResponseSchema:    resSchema,
		},
	)

	if err != nil {
		return "", fmt.Errorf("perform inference: %w", err)
	}

	return result.Text(), nil
}
