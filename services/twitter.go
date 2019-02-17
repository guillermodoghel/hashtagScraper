package services

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	. "github.com/guillermodoghel/hashtagScraper/models"
	"github.com/guillermodoghel/hashtagScraper/utils"
)

//GetTwitter returns []Post by hashtag from twitter
func GetTwitter(hashtag string) []Post {

	config := oauth1.NewConfig(utils.GetProperty("twitterConsumerKey"), utils.GetProperty("twitterConsumerSecret"))
	token := oauth1.NewToken(utils.GetProperty("twitterAccessToken"), utils.GetProperty("twitterAccessSecret"))
	httpClient := config.Client(oauth1.NoContext, token)

	client := twitter.NewClient(httpClient)
	search, _, _ := client.Search.Tweets(&twitter.SearchTweetParams{
		Query: hashtag,
		Count: 500,
	})
	var response []Post
	for _, twit := range search.Statuses {

		if len(twit.Entities.Media) >= 1 {
			post := Post{URL: twit.Entities.Media[0].MediaURLHttps, Text: twit.Text}
			response = append(response, post)
		}
	}

	return response
}
