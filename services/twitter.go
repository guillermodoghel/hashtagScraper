package services

import (
	"fmt"
	"log"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/guillermodoghel/hashtagScraper/models"
	"github.com/guillermodoghel/hashtagScraper/utils"
)

//GetTwitter returns []Post by hashtag from twitter
func GetTwitter(hashtag string) []models.Post {
	start := time.Now()

	config := oauth1.NewConfig(utils.GetProperty("twitterConsumerKey"), utils.GetProperty("twitterConsumerSecret"))
	token := oauth1.NewToken(utils.GetProperty("twitterAccessToken"), utils.GetProperty("twitterAccessSecret"))
	httpClient := config.Client(oauth1.NoContext, token)

	client := twitter.NewClient(httpClient)
	search, _, _ := client.Search.Tweets(&twitter.SearchTweetParams{
		Query: hashtag,
		Count: 100,
	})
	var response []models.Post
	for _, twit := range search.Statuses {
		t, err := time.Parse(time.RubyDate, twit.CreatedAt)
		if err != nil {
			fmt.Println(err)
		}

		if len(twit.Entities.Media) >= 1 {
			post := models.Post{URL: twit.Entities.Media[0].MediaURLHttps, Text: twit.Text, Date: int(t.Unix()), Origin: "twitter"}
			response = append(response, post)
		} else {
			post := models.Post{Author: twit.User.ScreenName, URL: "", Text: twit.Text, Date: int(t.Unix()), Origin: "twitter"}
			response = append(response, post)
		}
	}

	elapsed := time.Since(start)
	log.Printf("Twitter took %s", elapsed)

	return response
}
