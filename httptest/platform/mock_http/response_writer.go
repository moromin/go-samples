package mock_http

import (
	"encoding/json"
	"net/http"
)

type ResponseWriter struct {
	StatusCode int
	Written    []byte
}

func (m *ResponseWriter) Header() http.Header {
	return http.Header{}
}

func (m *ResponseWriter) Write(b []byte) (int, error) {
	m.Written = b
	return 0, nil
}

func (m *ResponseWriter) WriteHeader(statusCode int) {
	m.StatusCode = statusCode
}

func (m *ResponseWriter) GetBodyJSON() map[string]interface{} {
	var body map[string]interface{}
	json.Unmarshal(m.Written, &body)
	return body
}

func (m *ResponseWriter) GetBodyJSONArray() []map[string]interface{} {
	var body []map[string]interface{}
	json.Unmarshal(m.Written, &body)
	return body
}

func (m *ResponseWriter) GetBodyString() string {
	return string(m.Written)
}
