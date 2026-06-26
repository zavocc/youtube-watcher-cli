package dataapi

import (
	"fmt"

	"google.golang.org/api/youtube/v3"
)

func Channel(service *youtube.Service, query string, channelQueryType string, maxResults int64, nextPageToken string) (*youtube.PlaylistItemListResponse, error) {
	call := service.Channels.List([]string{"id", "contentDetails"}).
		MaxResults(1)

	// check if channelQueryType is "username", "handle" or "id" and set the appropriate parameter, by default, it will be "handle"
	switch channelQueryType {
	case "id":
		call = call.Id(query)
	case "username":
		call = call.ForUsername(query)
	case "handle":
		// check if it starts with @, if not, add it
		if query[0] != '@' {
			query = "@" + query
		}

		call = call.ForHandle(query)
	default:
		return nil, fmt.Errorf("invalid channel query type %q: expected id, username or handle", channelQueryType)
	}

	channelResponse, err := call.Do()
	if err != nil {
		return nil, err
	}

	if len(channelResponse.Items) == 0 {
		return nil, fmt.Errorf("no channel found for %s query %q", channelQueryType, query)
	}

	uploadsPlaylistID := channelResponse.Items[0].ContentDetails.RelatedPlaylists.Uploads
	if uploadsPlaylistID == "" {
		return nil, fmt.Errorf("channel %q has no uploads playlist", channelResponse.Items[0].Id)
	}

	return Playlist(service, uploadsPlaylistID, maxResults, nextPageToken)
}
