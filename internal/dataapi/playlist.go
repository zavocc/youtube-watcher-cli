package dataapi

import (
	"google.golang.org/api/youtube/v3"
)

func Playlist(service *youtube.Service, playlistID string, maxResults int64, nextPageToken string) (*youtube.PlaylistItemListResponse, error) {
	// Cap the maxResults to 50, if exceeds
	if maxResults > 50 {
		maxResults = 50
	}

	call := service.PlaylistItems.List([]string{"id", "snippet", "contentDetails"}).
		PlaylistId(playlistID).
		MaxResults(maxResults)

	if nextPageToken != "" {
		call = call.PageToken(nextPageToken)
	}

	return call.Do()
}
