package responses

import "github.com/Limpid-LLC/store-suppliers-helper/entities"

type ArticleSearchResponse struct {
	Details []entities.Detail `json:"details"`
	Count   int               `json:"count"`
}
