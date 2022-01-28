package dto

type Item struct {
	Id           string `json:"id"`
	Title        string `json:"title"`
	Link         string `json:"link"`
	PreText      string `json:"preText"`
	PostText     string `json:"postText"`
	BoldText     string `json:"boldText"`
	CompleteText string `json:"completeText"`
}

type SearchMeta struct {
	Date     string `json:"date"`
	Query    string `json:"query"`
	Paramter string `json:"parameter"`
}

type ResponseDto struct {
	Meta                  SearchMeta `json:"meta,omitempty"`
	KeyPartners           []Item     `json:"keyPartners,omitempty"`
	KeyActivities         []Item     `json:"keyActivities,omitempty"`
	KeyResources          []Item     `json:"keyResources,omitempty"`
	ValuePropositions     []Item     `json:"valuePropositions,omitempty"`
	CustomerRelationships []Item     `json:"customerRelationships,omitempty"`
	Channels              []Item     `json:"channel,omitempty"`
	CustomerSegments      []Item     `json:"customerSegments,omitempty"`
	CostStructure         []Item     `json:"costStructure,omitempty"`
	RevenueStreams        []Item     `json:"revenueStreams,omitempty"`
}
