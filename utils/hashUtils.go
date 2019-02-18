package utils

import (
	"bytes"
	"crypto/sha256"
	"fmt"

	. "github.com/guillermodoghel/hashtagScraper/models"
)

//SignPosts sign each post of an array with a sha256 hash
func SignPosts(posts *[]Post) {
	for index, post := range *posts {
		h := sha256.New()
		var buffer bytes.Buffer
		buffer.WriteString(post.URL)
		buffer.WriteString(post.Text)
		h.Write([]byte(buffer.String()))
		bs := h.Sum(nil)
		post.Hash = fmt.Sprintf("%x", bs)
		(*posts)[index] = post
	}
}
