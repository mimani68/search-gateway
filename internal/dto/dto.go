package dto

type Item struct {
	Id           string `json:"id"`
	Order        int    `json:"order"`
	PreText      string `json:"preText"`
	PostText     string `json:"postText"`
	BoldText     string `json:"boldText"`
	CompleteText string `json:"completeText"`
	Link         string `json:"link"`
}

type SearchRequest struct {
	Date     string `json:"date"`
	Query    string `json:"query"`
	Paramter string `json:"parameter"`
}

type ResponseDto struct {
	Meta                  SearchRequest `json:"meta"`
	KeyPartners           []Item        `json:"keyPartners"`
	KeyActivities         []Item        `json:"keyActivities"`
	KeyResources          []Item        `json:"keyResources"`
	ValuePropositions     []Item        `json:"valuePropositions"`
	CustomerRelationships []Item        `json:"customerRelationships"`
	Channels              []Item        `json:"channels"`
	CustomerSegments      []Item        `json:"customerSegments"`
	CostStructure         []Item        `json:"costStructure"`
	RevenueStreams        []Item        `json:"revenueStreams"`
}
