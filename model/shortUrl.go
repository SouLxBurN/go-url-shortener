package model

// ShortURL struct representation
type ShortURL struct {
	URL  string `json:"url" bson:"url"`
	Hash string `json:"hash" bson:"hash"`
}
