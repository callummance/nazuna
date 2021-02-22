package restclient

import (
	"net/url"

	"github.com/callummance/nazuna/messages"
)

type pagination struct {
	Before string `json:"before,omitempty"`
	After  string `json:"after,omitempty"`
}

func (p pagination) insertToValues(initialValues *url.Values) {
	if p.After != "" {
		(*initialValues)["after"] = []string{p.After}
	} else if p.Before != "" {
		(*initialValues)["before"] = []string{p.Before}
	}
}

type paginationCursor struct {
	Cursor string `json:"cursor"`
}

type subscriptionsPage struct {
	Total      int                     `json:"total"`
	Data       []messages.Subscription `json:"data"`
	Limit      int                     `json:"limit"`
	Pagination paginationCursor        `json:"pagination"`
}
