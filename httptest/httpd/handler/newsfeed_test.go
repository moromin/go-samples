package handler_test

import (
	"go-samples/httptest/httpd/handler"
	"go-samples/httptest/platform/mock_http"
	"go-samples/httptest/platform/newsfeed"
	"net/http"
	"testing"
)

func TestNewsfeedGet(t *testing.T) {
	feed := newsfeed.New()
	feed.Add(newsfeed.Item{"hello", "world"})

	handler := handler.NewsfeedGet(feed)

	w := &mock_http.ResponseWriter{}
	r := &http.Request{}

	handler(w, r)
	result := w.GetBodyJSONArray()

	if len(result) != 1 {
		t.Errorf("Expected 1 item, got %d", len(result))
	}

	if result[0]["title"] != "hello" {
		t.Errorf("Expected title 'hello', got %s", result[0]["title"])
	}
}

func TestNewsfeedPost(t *testing.T) {
	feed := newsfeed.New()

	headers := http.Header{}
	headers.Add("content-type", "application/json")

	w := &mock_http.ResponseWriter{}
	r := &http.Request{
		Header: headers,
	}

	r.Body = mock_http.RequestBody(map[string]interface{}{
		"title": "hello",
		"post":  "world",
	})

	handler := handler.NewsfeedPost(feed)
	handler(w, r)

	result := w.GetBodyJSON()

	if result["title"] != "hello" {
		t.Errorf("Expected title 'hello', got %s", result["title"])
	}

	if len(feed.GetAll()) != 1 {
		t.Errorf("Expected 1 item, got %d", len(feed.GetAll()))
	}

}
