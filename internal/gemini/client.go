package gemini

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/zavocc/youtube-watcher-cli/internal/config"
	"google.golang.org/genai"
)

func GApiClient(prompt string, url string, model string, resolution string) string {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, "An error has occurred while initializing the Gemini API client, error log:\n", err.Error())
		os.Exit(1)
	}

	// Check if it's either 11-character YouTube video ID or a full URL
	var actualUrl string
	if len(url) == 11 {
		actualUrl = "https://www.youtube.com/watch?v=" + url
	} else if len(url) > 11 && (url[:7] == "http://" || url[:8] == "https://") {
		actualUrl = strings.Split(url, "&")[0] // Remove any additional query parameters after the video ID
	} else {
		fmt.Fprintln(os.Stderr, "Invalid YouTube video ID or URL specified. Please provide a valid 11-character video ID or a full URL.")
		os.Exit(1)
	}

	// Validate model and get corresponding config
	modelSelectedConfig := config.ValidateModels(model)

	videoPart := genai.NewPartFromURI(actualUrl, "video/mp4")
	videoPart.MediaResolution = &genai.PartMediaResolution{
		Level: config.ParseMediaResolution(resolution),
	}

	contents := []*genai.Content{
		genai.NewContentFromParts([]*genai.Part{
			videoPart,
			genai.NewPartFromText(prompt),
		}, genai.RoleUser),
	}

	result, err := client.Models.GenerateContent(
		ctx, modelSelectedConfig.ModelID,
		contents,
		&genai.GenerateContentConfig{
			SystemInstruction: genai.NewContentFromText(systemPrompt, genai.RoleUser),
			ThinkingConfig:    modelSelectedConfig.ThinkingConfig,
		},
	)

	if err != nil {
		fmt.Fprintln(os.Stderr, "An error has occurred while performing inference, error log:\n", err.Error())
		os.Exit(1)
	}

	return result.Text()
}
