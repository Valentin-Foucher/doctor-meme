package main

import (
	"fmt"
	"os"

	"github.com/Valentin-Foucher/doctor-meme/pkg/clients"
	"github.com/Valentin-Foucher/doctor-meme/pkg/utils"
)

func main() {
	config := &utils.Configuration{}
	utils.ReadConfiguration(config)

	youtubeClient := clients.CreateHttpYoutubeClient(config)
	itemsCount, err := youtubeClient.CountItemsInPlaylists()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	video, err := youtubeClient.GetRandomVideoUrl(itemsCount)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = playVideo(*video)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
