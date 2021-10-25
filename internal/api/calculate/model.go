package calculate

import (
	"breeding/internal/model"
	"breeding/internal/model/axiebrieflist"
)

type CalculateRequest struct {
	AxieParentID1 string `json:"axieParentID1" binding:"required,numeric"`
	AxieParentID2 string `json:"axieParentID2" binding:"required,numeric"`
}
type CalculateResponse struct {
	AxieParent1CurrentPrice *float64    `json:"axieParent1CurrentPrice"`
	AxieParent1SalePrice    *float64    `json:"axieParent1SalePrice"`
	AxieParent2CurrentPrice *float64    `json:"axieParent2CurrentPrice"`
	AxieParent2SalePrice    *float64    `json:"axieParent2SalePrice"`
	AxieChildren            []AxieChild `json:"axieChildren"`
}

type AxieDetailVariables struct {
	AxieParentID1 string `json:"axieParentID1,omitempty"`
	AxieParentID2 string `json:"axieParentID2,omitempty"`
}
type AxieDetailResponse struct {
	Data AxieDetailData `json:"data"`
}
type AxieDetailData struct {
	AxieParent1 model.Axie `json:"axieParent1"`
	AxieParent2 model.Axie `json:"axieParent2"`
}

type AxieBriefListVariables struct {
	Parent1Criteria axiebrieflist.Criteria `json:"parent1Criteria,omitempty"`
	Parent2Criteria axiebrieflist.Criteria `json:"parent2Criteria,omitempty"`
	Child1Criteria  axiebrieflist.Criteria `json:"child1Criteria,omitempty"`
	Child2Criteria  axiebrieflist.Criteria `json:"child2Criteria,omitempty"`
	Child3Criteria  axiebrieflist.Criteria `json:"child3Criteria,omitempty"`
	Child4Criteria  axiebrieflist.Criteria `json:"child4Criteria,omitempty"`
	Child1Size      int                    `json:"child1Size"`
	Child2Size      int                    `json:"child2Size"`
	Child3Size      int                    `json:"child3Size"`
	Child4Size      int                    `json:"child4Size"`
}
type AxieBriefListResponse struct {
	Data AxieBriefListData `json:"data"`
}
type AxieBriefListData struct {
	AxieParent1 axiebrieflist.Axies `json:"axieParent1"`
	AxieParent2 axiebrieflist.Axies `json:"axieParent2"`
	AxieChild1  axiebrieflist.Axies `json:"axieChild1"`
	AxieChild2  axiebrieflist.Axies `json:"axieChild2"`
	AxieChild3  axiebrieflist.Axies `json:"axieChild3"`
	AxieChild4  axiebrieflist.Axies `json:"axieChild4"`
}

type AxieChild struct {
	Chance float64  `json:"chance"`
	Price  *float64 `json:"price"`
	Class  string   `json:"class"`
	Mouth  string   `json:"mouth"`
	Horn   string   `json:"horn"`
	Back   string   `json:"back"`
	Tail   string   `json:"tail"`
}
