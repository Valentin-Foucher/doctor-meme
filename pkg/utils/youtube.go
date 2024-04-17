package utils

import "fmt"

func BuildYoutubeUrl(videoId string) string {
	return fmt.Sprintf("https://www.youtube.com/watch?v=%s", videoId)
}
