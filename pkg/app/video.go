package main

import (
	"fmt"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func playVideo(videoUrl string) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Exited manually")
		}
	}()

	page := rod.New().ControlURL(
		launcher.
			New().
			Headless(false).
			UserDataDir("path").
			MustLaunch(),
	).MustConnect().MustPage(videoUrl)
	defer page.Close()

	page.MustElement("#movie_player").WaitLoad()
	page.MustElement(`#movie_player`).MustClick()
	page.MustEval("() => document.getElementById(\"movie_player\").click()")
	page.MustWaitStable()

	page.MustElement(".ended-mode").WaitLoad()

	return nil
}
