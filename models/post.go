package models

// Post represents a social media post
type Post struct {
	Author string
	URL    string
	Text   string
	Date   int
	Origin string
	Hash   string
}
