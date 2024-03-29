package repository

import (
	"axie-infinity-back/internal/client"
	"axie-infinity-back/internal/model"
	"context"
	"net/http"
)

const (
	AxieGraphQL string = "https://graphql-gateway.axieinfinity.com/graphql"
)

type Repository struct {
}

func NewRepository() RepositoryService {
	return &Repository{}
}

type RepositoryService interface {
	GetAxieBriefList(ctx context.Context, req model.Query, res interface{}) (*http.Response, error)
	GetAxieDetail(ctx context.Context, req model.Query, res interface{}) (*http.Response, error)
}

func (re *Repository) GetAxieBriefList(ctx context.Context, req model.Query, res interface{}) (*http.Response, error) {
	header := map[string]string{
		"Content-Type":  "application/json",
		"cache-control": "no-cache",
	}
	req.OperationName = "GetAxieBriefList"
	httpRes, err := client.Send(AxieGraphQL, http.MethodPost, header, req, res)
	if err != nil {
		return httpRes, err
	}
	return httpRes, nil
}

func (re *Repository) GetAxieDetail(ctx context.Context, req model.Query, res interface{}) (*http.Response, error) {
	header := map[string]string{
		"Content-Type":  "application/json",
		"cache-control": "no-cache",
	}
	req.OperationName = "GetAxieDetail"
	httpRes, err := client.Send(AxieGraphQL, http.MethodPost, header, req, res)
	if err != nil {
		return httpRes, err
	}
	return httpRes, nil
}
