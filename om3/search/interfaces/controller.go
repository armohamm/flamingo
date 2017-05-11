package interfaces

import (
	"flamingo/framework/web"
	"flamingo/framework/web/responder"
	"flamingo/om3/search/domain"
)

type (
	// ViewController demonstrates a search view controller
	ViewController struct {
		*responder.ErrorAware  `inject:""`
		*responder.RenderAware `inject:""`
		domain.SearchService   `inject:""`
	}

	// ViewData is used for search rendering
	ViewData struct {
		SearchResult map[string]interface{}
	}
)

// Get Response for search
func (vc *ViewController) Get(c web.Context) web.Response {
	query := c.MustQuery1("q")
	searchResult, err := vc.SearchService.Search(c, query)

	// catch error
	if err != nil {
		return vc.Error(c, err)
	}

	// render page
	return vc.Render(c, "pages/search/view", ViewData{SearchResult: map[string]interface{}{
		"type": "product", //@todo: add subroutes for other types
		"query": query,
		"results": map[string]interface{}{
			"product":  searchResult.Results.Product,
			"brand":    searchResult.Results.Brand,
			"location": searchResult.Results.Location,
			"retailer": searchResult.Results.Retailer,
		},
	}})
}
