package clients

import (
	"io"
	"math/rand/v2"
	"net/http"
	"os"

	"github.com/Valentin-Foucher/doctor-meme/pkg/utils"
	"github.com/tidwall/gjson"
)

type HttpYoutubeClient struct {
	client        *HttpClient
	configuration utils.YoutubeConfig
}

func CreateHttpYoutubeClient(configuration *utils.Configuration) *HttpYoutubeClient {
	defaultQueryParameters := make(map[string][]string)
	defaultQueryParameters["key"] = []string{os.Getenv("DOCTOR_MEME_GOOGLE_API_KEY")}
	return &HttpYoutubeClient{
		&HttpClient{
			configuration.HttpConfig.Youtube.BaseUrl,
			nil,
			defaultQueryParameters,
			configuration.HttpConfig.MaxRetries,
			&http.Client{},
		},
		configuration.YoutubeConfig,
	}
}

func (youtubeClient *HttpYoutubeClient) CountItemsInPlaylists() (map[string]int64, error) {
	queryParameters := make(map[string][]string)
	queryParameters["id"] = youtubeClient.configuration.PlaylistIDs
	queryParameters["part"] = []string{"contentDetails"}
	queryParameters["maxResults"] = []string{"50"}

	response, err := youtubeClient.client.get("playlists", []int{200}, nil, queryParameters)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	response.Body.Close()

	itemsCount := utils.MapSlice(
		gjson.Get(string(body), "items.#.contentDetails.itemCount").Array(),
		func(result gjson.Result) int64 {
			return result.Int()
		},
	)
	return utils.MapFromSlices(youtubeClient.configuration.PlaylistIDs, itemsCount), nil
}

func (youtubeClient *HttpYoutubeClient) GetRandomVideoUrl(playlistItemCounts map[string]int64) (*string, error) {
	totalVideos := utils.Sum(utils.MapValues(playlistItemCounts))
	videoIDs := make([]string, 0)
	videoIDsCh := make(chan (string))
	stopCh := make(chan (bool))

	remainingGoroutines := len(playlistItemCounts)

	for _, playlistId := range utils.MapKeys(playlistItemCounts) {
		go func(id string) error {
			var nextPageToken string
			queryParameters := make(map[string][]string)
			queryParameters["part"] = []string{"contentDetails"}
			queryParameters["maxResults"] = []string{"50"}

			for {
				queryParameters["pageToken"] = []string{nextPageToken}
				queryParameters["playlistId"] = []string{id}

				response, err := youtubeClient.client.get("playlistItems", []int{200}, nil, queryParameters)
				if err != nil {
					return err
				}

				body, err := io.ReadAll(response.Body)
				if err != nil {
					return err
				}
				for _, videoIdResult := range gjson.Get(string(body), "items.#.contentDetails.videoId").Array() {
					videoIDsCh <- videoIdResult.String()
				}
				response.Body.Close()

				nextPageToken = gjson.Get(string(body), "nextPageToken").String()
				if nextPageToken == "" {
					stopCh <- true
					break
				}
			}
			return nil

		}(playlistId)
	}

	for {
		select {
		case <-stopCh:
			remainingGoroutines--
		case videoId := <-videoIDsCh:
			videoIDs = append(videoIDs, videoId)
		}
		if remainingGoroutines == 0 {
			break
		}
	}

	result := new(string)
	*result = utils.BuildYoutubeUrl(videoIDs[rand.IntN(int(totalVideos))])
	return result, nil
}
