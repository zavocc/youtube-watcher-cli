package dataapi

import (
	"google.golang.org/api/youtube/v3"
)

func Video(service *youtube.Service, videoID string) (*youtube.VideoListResponse, error) {
	call := service.Videos.List([]string{"id", "snippet", "contentDetails"}).
		Id(videoID).
		MaxResults(1)

	return call.Do()
}
