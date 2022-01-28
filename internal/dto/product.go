package dto

type ProductCategory struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

type Metrics struct {
	MarketCap         string `json:"marketCap"`
	InterestForInvest string `json:"intrestForInvest"`
}

type ProductResponse struct {
	Category []ProductCategory `json:"category"`
	Metrics  Metrics           `json:"metrics"`
}
