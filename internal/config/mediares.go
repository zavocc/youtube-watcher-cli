package config

import (
	"fmt"
	"os"

	"google.golang.org/genai"
)

func ParseMediaResolution(resolution string) genai.PartMediaResolutionLevel {
	var mediaResLevel genai.PartMediaResolutionLevel

	switch resolution {
	case "low":
		mediaResLevel = genai.PartMediaResolutionLevelMediaResolutionLow
	case "high":
		mediaResLevel = genai.PartMediaResolutionLevelMediaResolutionHigh
	default:
		fmt.Fprintln(os.Stderr, "Invalid media resolution specified. Use --help parameter to see supported media resolutions.")
		os.Exit(1)
		mediaResLevel = genai.PartMediaResolutionLevelMediaResolutionUnspecified
	}
	return mediaResLevel
}
