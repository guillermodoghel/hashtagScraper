package services

import (
	"fmt"

	. "github.com/guillermodoghel/hashtagScraper/models"
	fb "github.com/huandu/facebook"
)

var globalApp = fb.New("383526045806634", "406f029e718b0be18b4cd3b9c381a684")

func GetFacebook(hashtag string) []Post {

	// Use the new session to send an API request with access token.
	res, _ := fb.Get("/search", fb.Params{
		"access_token": "383526045806634|6hA4cACKhb1GcsLtMZKA1NDJczc",
		"type":         "post",
		"q":            hashtag,
	})

	fmt.Println(res)

	return []Post{}
}
