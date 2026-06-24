package dataapi

import (
	"fmt"

	"google.golang.org/api/youtube/v3"
)

func Search(service *youtube.Service, query string, filter string, maxResults int64, nextPageToken string) (*youtube.SearchListResponse, error) {
	// Cap the maxResults to 50, if exceeds
	if maxResults > 50 {
		maxResults = 50
	}

	call := service.Search.List([]string{"id", "snippet"}).
		Q(query).
		MaxResults(maxResults).
		SafeSearch("none")

	switch filter {
	case "video", "playlist":
		call = call.Type(filter)
	case "mixed":
		// Omitting Type allows mixed search results.
	default:
		return nil, fmt.Errorf("invalid search filter %q: expected video, playlist, or mixed", filter)
	}

	if nextPageToken != "" {
		call = call.PageToken(nextPageToken)
	}

	return call.Do()
}
