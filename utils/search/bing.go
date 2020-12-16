package search

import (
	"leviathanRewritten/utils"

	"github.com/PuerkitoBio/goquery"
)

func SearchBing(query string, nsfw bool) ([]Result, error) {
	results := []Result{}

	resp, err := utils.GetResponse("https://www.bing.com/search?q=" + utils.Encode(query))
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return nil, err
	}

	sel := doc.Find(".b_algo h2 a")
	for i := range sel.Nodes {
		anchorTag := sel.Eq(i)

		title := anchorTag.Text()
		link, _ := anchorTag.Attr("href")

		results = append(results, Result{
			Title: title,
			Link:  link,
		})
	}

	return results, err
}
