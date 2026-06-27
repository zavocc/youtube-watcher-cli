package gemini

import "google.golang.org/genai"

func genResponseSchema() *genai.Schema {
	return &genai.Schema{
		Type: genai.TypeObject,
		Properties: map[string]*genai.Schema{
			"answer": {
				Type:        genai.TypeString,
				Description: "Direct answer to the user's prompt about the video.",
			},
			"evidence_timestamps": {
				Type:        genai.TypeArray,
				Description: "Supporting timestamp and passages from the video supporting the answer.",
				Items: &genai.Schema{
					Type: genai.TypeObject,
					Properties: map[string]*genai.Schema{
						"timestamp": {
							Type:        genai.TypeString,
							Description: "Video timestamp in MM:SS or HH:MM:SS format.",
						},
						"passage": {
							Type:        genai.TypeString,
							Description: "Short observation, quote, or paraphrase from that timestamp.",
						},
					},
					PropertyOrdering: []string{"timestamp", "passage"},
					Required:         []string{"timestamp", "passage"},
				},
			},
		},
		PropertyOrdering: []string{
			"answer",
			"evidence_timestamps",
		},
		Required: []string{
			"answer",
		},
	}
}
