package handler

import (
	"encoding/json"
	"go-samples/httptest/platform/newsfeed"
	"net/http"
)

func NewsfeedGet(feed newsfeed.Getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		items := feed.GetAll()
		json.NewEncoder(w).Encode(items)
	}
}

func NewsfeedPost(feed newsfeed.Adder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var item newsfeed.Item
		json.NewDecoder(r.Body).Decode(&item)
		feed.Add(item)
		json.NewEncoder(w).Encode(item)
	}
}
