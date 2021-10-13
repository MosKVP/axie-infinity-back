package model

type Query struct {
	OperationName string      `json:"operationName,omitempty"`
	Variables     interface{} `json:"variables,omitempty"`
	Query         string      `json:"query,omitempty"`
}

