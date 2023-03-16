// Package http provides a function for making HTTP requests and converting data to byte arrays.
package http

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

// Request sends an HTTP request with the specified method, URL, headers, and body.
func Request(method string, url string, headers map[string]string, body []byte) (*http.Response, error) {
	// Create a new request object with the specified URL and method
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	// Add any additional headers to the request
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Create a new HTTP client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// Return the response object
	return resp, nil
}

// ConvertToBytes converts data to a byte array in the specified format (either "json" or "xml").
func ConvertToBytes(data interface{}, format string) ([]byte, error) {
	switch format {
	case "json":
		jsonData, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		return jsonData, nil
	case "xml":
		xmlData, err := xml.Marshal(data)
		if err != nil {
			return nil, err
		}
		return xmlData, nil
	default:
		return nil, fmt.Errorf("unsupported format: %s", format)
	}
}
