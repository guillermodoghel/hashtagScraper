package services

import (
	"fmt"
	"log"
	"time"

	. "github.com/guillermodoghel/hashtagScraper/models"
	"github.com/guillermodoghel/hashtagScraper/utils"
	"github.com/tidwall/gjson"
	"gopkg.in/ahmdrz/goinsta.v2"
)

//GetInstagram returns []Post by hashtag from Instagram
func GetInstagram(hashtag string) []Post {
	start := time.Now()
	url := fmt.Sprintf("https://www.instagram.com/explore/tags/%s/?__a=1", hashtag)

	jsonResponse := utils.GetJSON(url)

	value := gjson.GetMany(jsonResponse, "graphql.hashtag.edge_hashtag_to_media.edges.#.node.thumbnail_src",
		"graphql.hashtag.edge_hashtag_to_media.edges.#.node.edge_media_to_caption.edges.0.node.text",
		"graphql.hashtag.edge_hashtag_to_media.edges.#.node.taken_at_timestamp",
		"graphql.hashtag.edge_hashtag_to_media.edges.#.node.owner.id")

	var response []Post
	imagesArray := value[0].Array()
	textArrays := value[1].Array()
	dates := value[2].Array()
	names := value[3].Array()

	for index, subelement := range imagesArray {
		if index < len(textArrays) {
			post := Post{Author: names[index].String(), URL: subelement.String(), Text: textArrays[index].String(), Date: int(dates[index].Int()), Origin: "instagram"}
			response = append(response, post)
		}

	}
	elapsed := time.Since(start)
	log.Printf("Instagram took %s", elapsed)

	return response
}

func GetInstagramStories(hashtag string) []Post {
	start := time.Now()

	inst := goinsta.New(utils.GetProperty("instagramAccount"), utils.GetProperty("instagramPassword"))

	if err := inst.Login(); err != nil {
		return nil
	}

	inst.Login()
	h := inst.NewHashtag(hashtag)
	stories, err := h.Stories()

	if err != nil {
		return nil
	}
	var response []Post
	for _, item := range stories.Items {
		if len(item.Images.Versions) != 0 {
			post := Post{Author: item.User.Username, URL: item.Images.Versions[0].URL, Date: int(item.TakenAt), Origin: "instagram_stories"}
			response = append(response, post)
		}
	}
	elapsed := time.Since(start)
	log.Printf("Instagram stories took %s", elapsed)
	return response
}
