package search

import (
	"strings"

	"leviathanRewritten/config"
	"leviathanRewritten/utils"

	"github.com/PuerkitoBio/goquery"
)

//GOOGLESEARCH google search base url
const GOOGLESEARCH = "https://www.google.com.br/search?q="
const LANG = "pt"

func SearchGoogle(query string, cType bool) ([]Result, error) {
	var securityType string

	if !cType && config.Data.SafeMode {
		securityType = "&safe=active"
	} else {
		securityType = ""
	}

	results := []Result{}
	resp, err := utils.GetResponse(GOOGLESEARCH + utils.Encode(query) + "&hl=" + utils.Encode(LANG) + securityType)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return nil, err
	}

	sel := doc.Find("div.g")
	for i := range sel.Nodes {
		item := sel.Eq(i)
		linkTag := item.Find("a")
		link, _ := linkTag.Attr("href")
		link = strings.Replace(link, "/url?q=", "", -1)
		titleTag := item.Find("h3.r")
		title := titleTag.Text()

		if link != "" && link != "#" && !strings.HasPrefix(link, "/") {
			result := Result{
				title,
				link,
			}

			results = append(results, result)
		}
	}
	return results, err
}
