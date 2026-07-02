package gemini

import (
	"context"
	"fmt"
	"os"

	"google.golang.org/genai"
)

// initializes gemini client
func initGeminiClient(ctx context.Context) (*genai.Client, error) {
	// check if GOOGLE_GENAI_USE_ENTERPRISE is set to true
	// If using enterprise, we check the existence of GOOGLE_APPLICATION_CREDENTIALS, GOOGLE_CLOUD_PROJECT
	_, useEnterprise := os.LookupEnv("GOOGLE_GENAI_USE_ENTERPRISE")
	if useEnterprise {
		_, credsExists := os.LookupEnv("GOOGLE_APPLICATION_CREDENTIALS")
		_, projectExists := os.LookupEnv("GOOGLE_CLOUD_PROJECT")

		if !projectExists || !credsExists {
			return nil, fmt.Errorf("GCP environment variables are not set, please set GOOGLE_APPLICATION_CREDENTIALS and GOOGLE_CLOUD_PROJECT.")
		}

		// unset GEMINI_API_KEY if it exists, since we are using enterprise credentials
		_, geminiKeyExists := os.LookupEnv("GEMINI_API_KEY")
		if geminiKeyExists {
			os.Unsetenv("GEMINI_API_KEY")
		}

		// if GOOGLE_CLOUD_LOCATION is set, set it otherwise default to global.
		// We take GOOGLE_CLOUD_LOCATION as priority over GOOGLE_CLOUD_REGION if both are set
		_, locationExists := os.LookupEnv("GOOGLE_CLOUD_LOCATION")
		if !locationExists {
			os.Setenv("GOOGLE_CLOUD_LOCATION", "global")
		}

		// Set region to global
		return genai.NewClient(ctx, &genai.ClientConfig{
			Project:  os.Getenv("GOOGLE_CLOUD_PROJECT"),
			Location: os.Getenv("GOOGLE_CLOUD_LOCATION"),
			Backend:  genai.BackendEnterprise,
		})
	} else {
		_, exists := os.LookupEnv("GEMINI_API_KEY")

		if !exists {
			return nil, fmt.Errorf("GEMINI_API_KEY environment variable not set, please set it.")
		}

		return genai.NewClient(ctx, &genai.ClientConfig{
			APIKey: os.Getenv("GEMINI_API_KEY"),
		})
	}

}
