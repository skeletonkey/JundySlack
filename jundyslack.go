package jundyslack

import (
	"encoding/json"
	"fmt"
	"net/http"
	"bytes"
)

// Slack is the container of all things slacking
type Slack struct {
	url string
}

// SetURL used to set the URL end point
func (s *Slack) SetURL(url string) (err error) {
	s.url = url

	return err
}

// Send is used to send a text message to the provided slack endpoint
func (s Slack) Send(msg string) (error) {
	reqBody, err := json.Marshal(map[string]string{"text": msg})
	if err != nil {
		return fmt.Errorf("unable to create request body: %s", err)
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", s.url, bytes.NewBuffer(reqBody))
	if err != nil {
		return fmt.Errorf("error attempting to create Slack request: %s", err)
	}
	req.Header.Add("Content-type", "application/json")
	_, err = client.Do(req)
	if err != nil {
		return fmt.Errorf("error attempting to connect to Slack (%s): %s", s.url, err)
	}

	return nil
}
