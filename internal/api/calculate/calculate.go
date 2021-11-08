package calculate

import (
	"breeding/internal/log"
	"breeding/internal/model"
	"breeding/internal/model/axiebrieflist"
	"breeding/internal/repository"
	"breeding/internal/util"
	"container/heap"
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/schwarmco/go-cartesian-product"
	"github.com/shanemaglangit/agp"
)

func Calculate(c *gin.Context) {
	var req CalculateRequest
	var err error
	var intErr *model.Error
	var res CalculateResponse
	var p1Sale, p2Sale *float64
	var filteredTopChildren []AxieChild

	defer func() {
		util.Response(c, res, intErr, recover())
	}()

	if err = c.ShouldBindJSON(&req); err != nil {
		if err, ok := err.(validator.ValidationErrors); ok {
			intErr = util.ConvertValidationError(err)
		} else {
			intErr = &model.Error{
				HTTPCode:        http.StatusBadRequest,
				Message:         "Decode JSON Error",
				InternalMessage: err.Error(),
			}
		}
		return
	}

	repo := repository.NewRepository()
	axieParent1, axieParent2, err := getParentsDetail(c, repo, req)
	if err != nil {
		intErr = &model.Error{
			HTTPCode:        http.StatusInternalServerError,
			Message:         "Could not getParentsDetail",
			InternalMessage: err.Error(),
		}
		return
	}

	if intErr = checkBreedingCondition(axieParent1, axieParent2); intErr != nil {
		intErr.HTTPCode = http.StatusBadRequest
		return
	}

	axieParent1Genes, err := agp.ParseHexDecode(axieParent1.Genes)
	if err != nil {
		intErr = &model.Error{
			HTTPCode:        http.StatusInternalServerError,
			Message:         "Could not ParseHexDecode",
			InternalMessage: err.Error(),
		}
		return
	}
	axieParent2Genes, err := agp.ParseHexDecode(axieParent2.Genes)
	if err != nil {
		intErr = &model.Error{
			HTTPCode:        http.StatusInternalServerError,
			Message:         "Could not ParseHexDecode",
			InternalMessage: err.Error(),
		}
		return
	}

	topChildren := getTopCombinations(axieParent1.Class, axieParent2.Class, axieParent1Genes, axieParent2Genes)

	p1Sale, p2Sale, filteredTopChildren, err = getAxiePrices(c, repo, axieParent1.Class, axieParent2.Class, axieParent1Genes, axieParent2Genes, topChildren)
	if err != nil {
		intErr = &model.Error{
			HTTPCode:        http.StatusInternalServerError,
			Message:         "Could not getAxiePrices",
			InternalMessage: err.Error(),
		}
		return
	}

	res.AxieParent1 = AxieParentDetail{
		BreedCount:   axieParent1.BreedCount,
		CurrentPrice: util.ConvertETH(axieParent1.Auction.CurrentPrice),
		SalePrice:    p1Sale,
	}
	res.AxieParent2 = AxieParentDetail{
		BreedCount:   axieParent2.BreedCount,
		CurrentPrice: util.ConvertETH(axieParent2.Auction.CurrentPrice),
		SalePrice:    p2Sale,
	}
	res.AxieChildren = filteredTopChildren
}

func getParentsDetail(c context.Context, repo repository.RepositoryService, req CalculateRequest) (axieParent1, axieParent2 model.Axie, err error) {
	var res AxieDetailResponse
	q := model.Query{
		Variables: AxieDetailVariables(req),
		Query:     getParentsDetailQuery,
	}
	_, err = repo.GetAxieDetail(c, q, &res)
	if err != nil {
		return
	}
	axieParent1 = res.Data.AxieParent1
	axieParent2 = res.Data.AxieParent2
	return
}

func checkBreedingCondition(axieParent1 model.Axie, axieParent2 model.Axie) (intErr *model.Error) {
	// Check Parent exists
	if axieParent1.IsEmpty() {
		intErr = &model.Error{
			Message: "Axie 1 Not Found",
		}
		return
	}

	if axieParent2.IsEmpty() {
		intErr = &model.Error{
			Message: "Axie 2 Not Found",
		}
		return
	}

	// Check Axie Relationship
	family1 := buildFamily(axieParent1)
	family2 := buildFamily(axieParent2)
	for _, m1 := range family1 {
		for _, m2 := range family2 {
			if m1 == m2 {
				intErr = &model.Error{
					Message: "Cannot breed with parent, child, or sibling",
				}
				return
			}
		}
	}

	// Check Axie Stage
	if axieParent1.Stage != 4 {
		intErr = &model.Error{
			Message: "Axie 1 is not an adult",
		}
		return
	}
	if axieParent2.Stage != 4 {
		intErr = &model.Error{
			Message: "Axie 2 is not an adult",
		}
		return
	}
	return nil
}

func buildFamily(axie model.Axie) (family []string) {
	if axie.ID != "" {
		family = append(family, axie.ID)
	}
	if axie.MatronID != 0 {
		family = append(family, fmt.Sprint(axie.MatronID))
	}
	if axie.SireID != 0 {
		family = append(family, fmt.Sprint(axie.SireID))
	}
	return family
}

func getTopCombinations(axieParent1Class, axieParent2Class model.Class, axieParent1Genes, axieParent2Genes agp.Genes) []AxieChild {
	const limit = 4
	topChildren := make([]AxieChild, limit)
	mouth, horn, back, tail := combineGenes(axieParent1Genes, axieParent2Genes)
	class := combineClass(axieParent1Class, axieParent2Class)
	// TODO: improve performance https://www.geeksforgeeks.org/k-maximum-sum-combinations-two-arrays/

	cha := cartesian.Iter(makeClassArray(class), makeAxiePartArray(mouth), makeAxiePartArray(horn), makeAxiePartArray(back), makeAxiePartArray(tail))
	pq := make(PriorityQueue, 0)
	for product := range cha {
		classID := product[0].(model.Class)
		mouthPart := product[1].(AxiePart)
		hornPart := product[2].(AxiePart)
		backPart := product[3].(AxiePart)
		tailPart := product[4].(AxiePart)
		axieChild := AxieChild{
			Chance: class[classID] * mouthPart.Chance * hornPart.Chance * backPart.Chance * tailPart.Chance,
			Class:  classID,
			Mouth:  mouthPart,
			Horn:   hornPart,
			Back:   backPart,
			Tail:   tailPart,
		}
		heap.Push(&pq, axieChild)
	}

	// Take most likely combinations
	n := util.MinInt(limit, len(pq))
	for i := 0; i < n; i++ {
		axieChild := heap.Pop(&pq).(AxieChild)
		topChildren[i] = axieChild
	}
	log.Logger.Infof("Top Children: %v", topChildren)
	return topChildren
}
func combineClass(class1, class2 model.Class) map[model.Class]float64 {
	m := make(map[model.Class]float64)
	m[class1] = 0.5
	v := m[class2]
	v = v + 0.5
	m[class2] = v
	return m
}

func makeClassArray(m map[model.Class]float64) []interface{} {
	arr := make([]interface{}, len(m))

	i := 0
	for k := range m {
		arr[i] = k
		i++
	}
	return arr
}
func combineGenes(genes1, genes2 agp.Genes) (mouth, horn, back, tail map[AxiePart]float64) {
	mouth = make(map[AxiePart]float64)
	horn = make(map[AxiePart]float64)
	back = make(map[AxiePart]float64)
	tail = make(map[AxiePart]float64)
	addMapPart(mouth, genes1.Mouth)
	addMapPart(mouth, genes2.Mouth)
	addMapPart(horn, genes1.Horn)
	addMapPart(horn, genes2.Horn)
	addMapPart(back, genes1.Back)
	addMapPart(back, genes2.Back)
	addMapPart(tail, genes1.Tail)
	addMapPart(tail, genes2.Tail)
	return
}

func addMapPart(m map[AxiePart]float64, part agp.Part) {
	const (
		D  = 0.375
		R1 = 0.09375
		R2 = 0.03125
	)
	addMapValue(m, makeAxiePart(part.D), D)
	addMapValue(m, makeAxiePart(part.R1), R1)
	addMapValue(m, makeAxiePart(part.R2), R2)
}

func makeAxiePart(partGene agp.PartGene) AxiePart {
	return AxiePart{
		PartID: mapSpecialPartID(partGene.PartId),
		Class:  model.Class(strings.Title(string(partGene.Class))),
		Name:   partGene.Name,
	}
}

func mapSpecialPartID(partID string) string {
	if v, ok := model.MapSpecialPartID[partID]; ok {
		return v
	}
	return partID
}

func addMapValue(m map[AxiePart]float64, key AxiePart, value float64) {
	v := m[key]
	v = v + value
	m[key] = v
}

func makeAxiePartArray(m map[AxiePart]float64) []interface{} {
	arr := make([]interface{}, len(m))

	i := 0
	for k, v := range m {
		k.Chance = v
		arr[i] = k
		i++
	}
	return arr
}

func getAxiePrices(c context.Context, repo repository.RepositoryService, axieParent1Class, axieParent2Class model.Class,
	axieParent1Genes, axieParent2Genes agp.Genes, topChildren []AxieChild,
) (p1Sale, p2Sale *float64, filteredTopChildren []AxieChild, err error) {
	var res AxieBriefListResponse
	q := model.Query{
		Variables: AxieBriefListVariables{
			Parent1Criteria: makeAxieParentCriteria(axieParent1Genes, axieParent1Class),
			Parent2Criteria: makeAxieParentCriteria(axieParent2Genes, axieParent2Class),
			Child1Criteria:  makeAxieChildCriteria(topChildren[0]),
			Child2Criteria:  makeAxieChildCriteria(topChildren[1]),
			Child3Criteria:  makeAxieChildCriteria(topChildren[2]),
			Child4Criteria:  makeAxieChildCriteria(topChildren[3]),
			Child1Size:      makeQuerySize(topChildren[0]),
			Child2Size:      makeQuerySize(topChildren[1]),
			Child3Size:      makeQuerySize(topChildren[2]),
			Child4Size:      makeQuerySize(topChildren[3]),
		},
		Query: getAxiesPriceQuery,
	}
	_, err = repo.GetAxieBriefList(c, q, &res)
	if err != nil {
		return
	}
	p1Sale = getAxiePrice(res.Data.AxieParent1)
	p2Sale = getAxiePrice(res.Data.AxieParent2)
	topChildren[0].Price = getAxiePrice(res.Data.AxieChild1)
	topChildren[1].Price = getAxiePrice(res.Data.AxieChild2)
	topChildren[2].Price = getAxiePrice(res.Data.AxieChild3)
	topChildren[3].Price = getAxiePrice(res.Data.AxieChild4)

	filteredTopChildren = filterTopChildren(topChildren)
	expectedPrice := calculateExpectedPrice(topChildren)
	log.Logger.Infof("Expected Price: %v", expectedPrice)
	return
}

func makeAxieParentCriteria(axieGenes agp.Genes, axieClass model.Class) axiebrieflist.Criteria {
	criteria := axiebrieflist.Criteria{
		Parts: []string{
			axieGenes.Mouth.D.PartId,
			axieGenes.Horn.D.PartId,
			axieGenes.Back.D.PartId,
			axieGenes.Tail.D.PartId,
		},
		Classes: []model.Class{
			axieClass,
		},
	}
	return criteria
}

func makeAxieChildCriteria(axieChild AxieChild) axiebrieflist.Criteria {
	criteria := axiebrieflist.Criteria{
		Parts: []string{
			axieChild.Mouth.PartID,
			axieChild.Horn.PartID,
			axieChild.Back.PartID,
			axieChild.Tail.PartID,
		},
		BreedCount: []int{0, 0},
	}
	if axieChild.Class != "" {
		criteria.Classes = append(criteria.Classes, model.Class(axieChild.Class))
	}
	return criteria
}

func makeQuerySize(axieChild AxieChild) int {
	if axieChild.Class == "" {
		return 0
	}
	return 1
}

func getAxiePrice(res axiebrieflist.Axies) *float64 {
	if len(res.Results) == 0 {
		return nil
	}
	return util.ConvertETH(res.Results[0].Auction.CurrentPrice)
}

func filterTopChildren(topChildren []AxieChild) []AxieChild {
	var filtered []AxieChild
	for _, child := range topChildren {
		if child.Class != "" {
			filtered = append(filtered, child)
		}
	}
	return filtered
}

func calculateExpectedPrice(topChildren []AxieChild) float64 {
	var totalChance float64
	var expectedPrice float64
	for _, child := range topChildren {
		if child.Price != nil {
			totalChance += child.Chance
			expectedPrice += child.Chance * *child.Price
		}
	}
	return expectedPrice / totalChance
}
