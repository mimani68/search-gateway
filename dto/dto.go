package dto

type ArticleContent struct {
	Id      string
	Title   string
	Summery string
	Body    string
	Link    string
}

type SearchRequest struct {
	Date     string
	Query    string
	Paramter string
}

type ResponceDto struct {
	Request SearchRequest
	Data    []ArticleContent
}
