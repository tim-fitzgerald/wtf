package spoke

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

type Resource struct {
	Response interface{}
	Raw      string
}

func (widget *Widget) api(method string, path string, params string) (*Resource, error) {
	client := &http.Client{}

	baseURL := "https://api.askspoke.com/api/v1/"
	URL := baseURL + path

	req, err := http.NewRequest(method, URL, bytes.NewBufferString(params))
	req.Header.Add("Api-Key", widget.settings.apiKey)

	q := req.URL.Query()     // Get a copy of the query values.
	q.Add("filter", "inbox") // Add a new value to the set.
	q.Add("status", "OPEN")
	req.URL.RawQuery = q.Encode() // Encode and assign back to the original query.

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &Resource{Response: &resp, Raw: string(data)}, nil
}
