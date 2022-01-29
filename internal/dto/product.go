package dto

type ProductCategory struct {
	Id         string `json:"id"`
	Title      string `json:"title"`
	TitleFarsi string `json:"titleFa"`
}

type Metrics struct {
	MarketCap         string `json:"marketCap"`
	InterestForInvest string `json:"intrestForInvest"`
	WorldBudget       string `json:"worldBudget"`
}

type ProductMetadataResponse struct {
	Category []ProductCategory `json:"category"`
}

type ProductResponse struct {
	Category []ProductCategory `json:"category"`
	Metrics  Metrics           `json:"metrics"`
}
