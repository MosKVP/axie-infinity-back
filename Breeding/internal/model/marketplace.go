package model

type AuctionType string

const (
	AuctionTypeAll        AuctionType = "All"
	AuctionTypeSale       AuctionType = "Sale"
	AuctionTypeNotForSale AuctionType = "NotForSale"
)

type Sort string

const (
	SortPriceAsc  Sort = "PriceAsc"
	SortPriceDesc Sort = "PriceDesc"
	SortIDAsc     Sort = "IdAsc"
	SortIDDesc    Sort = "IdDesc"
	SortLatest    Sort = "Latest"
)

type Auction struct {
	StartingPrice     string `json:"startingPrice,omitempty"`
	EndingPrice       string `json:"endingPrice,omitempty"`
	StartingTimestamp string `json:"startingTimestamp,omitempty"`
	EndingTimestamp   string `json:"endingTimestamp,omitempty"`
	Duration          string `json:"duration,omitempty"`
	TimeLeft          string `json:"timeLeft,omitempty"`
	CurrentPrice      string `json:"currentPrice,omitempty"`
	CurrentPriceUSD   string `json:"currentPriceUSD,omitempty"`
	SuggestedPrice    string `json:"suggestedPrice,omitempty"`
	Seller            string `json:"seller,omitempty"`
	ListingIndex      int    `json:"listingIndex,omitempty"`
	State             string `json:"state,omitempty"`
	Typename          string `json:"__typename,omitempty"`
}
