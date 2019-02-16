package services

import (
	"fmt"

	. "github.com/guillermodoghel/example/models"
	"github.com/guillermodoghel/example/utils"
	"github.com/tidwall/gjson"
)

func GetInstagram(hashtag string) []Post {
	url := fmt.Sprintf("https://www.instagram.com/explore/tags/%s/?__a=1", hashtag)

	result := utils.GetJSON(url)
	json := string(result[:])
	value := gjson.GetMany(json, "graphql.hashtag.edge_hashtag_to_media.edges.#.node.thumbnail_src",
		"graphql.hashtag.edge_hashtag_to_media.edges.#.node.edge_media_to_caption.edges.0.node.text")

	var response []Post
	textArrays := value[1].Array()
	for index, subelement := range value[0].Array() {
		println(subelement.String())

		post := Post{URL: subelement.String(), Text: textArrays[index].String()}
		response = append(response, post)

	}

	return response
}