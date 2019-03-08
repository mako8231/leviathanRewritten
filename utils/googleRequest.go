package utils

import (
	"strings"

	"leviathanRewritten/config"

	"github.com/PuerkitoBio/goquery"
)

type result struct {
	Title string
	Link  string
}

//GOOGLESEARCH google search base url
const GOOGLESEARCH = "https://www.google.com.br/search?q="
const LANG = "pt"

func GoogleParse(query string, cType bool) ([]result, error) {
	var securityType string

	if !cType && config.Data.SafeMode {
		securityType = "&safe=active"
	} else {
		securityType = ""
	}

	results := []result{}
	resp, err := GetResponse(GOOGLESEARCH + Encode(query) + "&hl=" + Encode(LANG) + securityType)
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
			result := result{
				title,
				link,
			}

			results = append(results, result)
		}
	}
	return results, err
}
