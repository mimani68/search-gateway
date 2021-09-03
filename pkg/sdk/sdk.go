package sdk

import "market.ir/dto"

func Search(keyword string) []dto.ArticleContent {
	return []dto.ArticleContent{
		{
			Title: "How sell something",
		},
	}
}
