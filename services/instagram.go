package services

import (
	"fmt"

	"github.com/guillermodoghel/hashtagScraper/models"
	"github.com/guillermodoghel/hashtagScraper/utils"
	"github.com/tidwall/gjson"
)

//GetInstagram returns []Post by hashtag from Instagram
func GetInstagram(hashtag string) []models.Post {
	url := fmt.Sprintf("https://www.instagram.com/explore/tags/%s/?__a=1", hashtag)

	jsonResponse := utils.GetJSON(url)

	value := gjson.GetMany(jsonResponse, "graphql.hashtag.edge_hashtag_to_media.edges.#.node.thumbnail_src",
		"graphql.hashtag.edge_hashtag_to_media.edges.#.node.edge_media_to_caption.edges.0.node.text")

	var response []models.Post
	textArrays := value[1].Array()

	for index, subelement := range value[0].Array() {
		post := models.Post{URL: subelement.String(), Text: textArrays[index].String()}
		response = append(response, post)
	}
	return response
}
