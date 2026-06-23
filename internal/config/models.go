package config

import (
	"fmt"
	"os"

	"google.golang.org/genai"
)

type configTemplate struct {
	ModelID        string
	ThinkingConfig *genai.ThinkingConfig
}

// Thinking budgets
var thinkingBudget = int32(1000)

func ValidateModels(model string) configTemplate {
	switch model {
	case "gemini-2.5-flash":
		return configTemplate{
			ModelID: "gemini-2.5-flash",
			ThinkingConfig: &genai.ThinkingConfig{
				ThinkingBudget:  &thinkingBudget,
				IncludeThoughts: false,
			},
		}
	case "gemini-3-flash-preview":
		return configTemplate{
			ModelID: "gemini-3-flash-preview",
			ThinkingConfig: &genai.ThinkingConfig{
				ThinkingLevel:   genai.ThinkingLevelMinimal,
				IncludeThoughts: false,
			},
		}
	case "gemini-3.1-flash-lite":
		return configTemplate{
			ModelID: "gemini-3.1-flash-lite",
			ThinkingConfig: &genai.ThinkingConfig{
				ThinkingLevel:   genai.ThinkingLevelLow,
				IncludeThoughts: false,
			},
		}
	default:
		fmt.Fprintln(os.Stderr, "Invalid model specified. Use --help parameter to see supported models.")
		os.Exit(1)
	}

	return configTemplate{}
}
