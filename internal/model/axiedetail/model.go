package axiedetail

import "breeding/internal/model"

type Response struct {
	Data Data `json:"data,omitempty"`
}
type Data struct {
	Axie model.Axie `json:"axie,omitempty"`
}

type Variables struct {
	AxieID string `json:"axieID,omitempty"`
}
