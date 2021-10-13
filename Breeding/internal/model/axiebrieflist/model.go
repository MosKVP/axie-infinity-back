package axiebrieflist

import "breeding/internal/model"

type Response struct {
	Data Data `json:"data,omitempty"`
}
type Data struct {
	Axies Axies `json:"axies,omitempty"`
}
type Axies struct {
	Total    int          `json:"total,omitempty"`
	Results  []model.Axie `json:"results,omitempty"`
	Typename string       `json:"__typename,omitempty"`
}

type Variables struct {
	From        int               `json:"from,omitempty"`
	Size        int               `json:"size,omitempty"`
	Sort        model.Sort        `json:"sort,omitempty"`
	AuctionType model.AuctionType `json:"auctionType,omitempty"`
	Owner       *string           `json:"owner,omitempty"`
	Criteria    Criteria          `json:"criteria,omitempty"`
}
type Criteria struct {
	Region     string        `json:"region,omitempty"`
	Parts      []string      `json:"parts,omitempty"`
	BodyShapes []string      `json:"bodyShapes,omitempty"`
	Classes    []model.Class `json:"classes,omitempty"`
	Stages     []int         `json:"stages,omitempty"`
	NumMystic  []int         `json:"numMystic,omitempty"`
	Pureness   []int         `json:"pureness,omitempty"`
	Title      []string      `json:"title,omitempty"`
	Breedable  *bool         `json:"breedable,omitempty"`
	BreedCount []int         `json:"breedCount,omitempty"`
	Hp         []int         `json:"hp,omitempty"`
	Skill      []int         `json:"skill,omitempty"`
	Speed      []int         `json:"speed,omitempty"`
	Morale     []int         `json:"morale,omitempty"`
}
