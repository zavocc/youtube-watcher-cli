package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/zavocc/youtube-watcher-cli/internal/shared"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func newYouTubeService(ctx context.Context) (*youtube.Service, error) {
	if err := shared.LoadEnvironment(); err != nil {
		return nil, fmt.Errorf("load environment: %w", err)
	}

	apiKey, exists := os.LookupEnv("YOUTUBE_DATA_API_KEY")
	if !exists || apiKey == "" {
		return nil, errors.New("YOUTUBE_DATA_API_KEY is not set")
	}

	service, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, fmt.Errorf("create YouTube service: %w", err)
	}

	return service, nil
}
