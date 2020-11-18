package spoke

import (
	"encoding/json"
	"time"
)

type RequestsArray struct {
	Results []Request
	Total   int `json:"total"`
	Start   int `json:"start"`
	Limit   int `json:"limit"`
}

type Request struct {
	Subject              string        `json:"subject"`
	Requester            string        `json:"requester"`
	Owner                string        `json:"owner"`
	Status               string        `json:"status"`
	PrivacyLevel         string        `json:"privacyLevel"`
	Team                 string        `json:"team"`
	Org                  string        `json:"org"`
	Permalink            string        `json:"permalink"`
	CreatedAt            time.Time     `json:"createdAt"`
	ID                   string        `json:"id"`
	IsAutoResolved       bool          `json:"isAutoResolved"`
	IsFiled              bool          `json:"isFiled"`
	Email                string        `json:"email"`
	UpdatedAt            time.Time     `json:"updatedAt"`
	TeamResponseTimeInMs int           `json:"teamResponseTimeInMs"`
	Tags                 []interface{} `json:"tags"`
	RequestTypeInfo      struct {
		AnsweredFields []interface{} `json:"answeredFields"`
	} `json:"requestTypeInfo"`
}

func (widget *Widget) listRequests() (*RequestsArray, error) {
	RequestStruct := &RequestsArray{}
	resource, err := widget.api("GET", "/requests", "")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(resource.Raw), RequestStruct)
	if err != nil {
		return nil, err
	}

	return RequestStruct, err
}
