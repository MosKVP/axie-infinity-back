package calculate

var getAxiesPriceQuery = `query GetAxieBriefList(
	$parent1Criteria: AxieSearchCriteria
	$parent2Criteria: AxieSearchCriteria
	$child1Criteria: AxieSearchCriteria
	$child2Criteria: AxieSearchCriteria
	$child3Criteria: AxieSearchCriteria
	$child4Criteria: AxieSearchCriteria
	$child1Size: Int
	$child2Size: Int
	$child3Size: Int
	$child4Size: Int
  ) {
	axieParent1: axies(
	  auctionType: Sale
	  criteria: $parent1Criteria
	  sort: PriceAsc
	  size: 1
	) {
	  ...SearchResult
	}
	axieParent2: axies(
	  auctionType: Sale
	  criteria: $parent2Criteria
	  sort: PriceAsc
	  size: 1
	) {
	  ...SearchResult
	}
	axieChild1: axies(
	  auctionType: Sale
	  criteria: $child1Criteria
	  sort: PriceAsc
	  size: $child1Size
	) {
	  ...SearchResult
	}
  
	axieChild2: axies(
	  auctionType: Sale
	  criteria: $child2Criteria
	  sort: PriceAsc
	  size: $child2Size
	) {
	  ...SearchResult
	}
  
	axieChild3: axies(
	  auctionType: Sale
	  criteria: $child3Criteria
	  sort: PriceAsc
	  size: $child3Size
	) {
	  ...SearchResult
	}
  
	axieChild4: axies(
	  auctionType: Sale
	  criteria: $child4Criteria
	  sort: PriceAsc
	  size: $child4Size
	) {
	  ...SearchResult
	}
  }
  fragment SearchResult on Axies {
	total
	results {
	  ...AxiePrice
	  __typename
	}
	__typename
  }
  fragment AxiePrice on Axie {
	auction {
	  currentPrice
	  __typename
	}
	__typename
  }  
`
var getParentsDetailQuery = `query GetAxieDetail($axieParentID1: ID!, $axieParentID2: ID!) {
	axieParent1: axie(axieId: $axieParentID1) {
	  ...AxieDetail
	  __typename
	}
	axieParent2: axie(axieId: $axieParentID2) {
	  ...AxieDetail
	  __typename
	}
  }
  fragment AxieDetail on Axie {
	id
	image
	name
	genes
	class
	stage
	sireId
	matronId
	breedCount
	auction {
	  ...AxieAuction
	  __typename
	}
	battleInfo {
	  ...AxieBattleInfo
	  __typename
	}
	__typename
  }
  fragment AxieBattleInfo on AxieBattleInfo {
	banned
	banUntil
	level
	__typename
  }
  fragment AxieAuction on Auction {
	currentPrice
	currentPriceUSD
	__typename
  }  
`
