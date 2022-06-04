package model

type (
	Response struct {
		Message   string      `json:"message"`
		Status    string      `json:"status"`
		Timestamp string      `json:"timestamp"`
		Data      interface{} `json:"data,omitempty"`
	}

	Empty struct{}
)
