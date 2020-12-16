package search

// Result é um resultado de pesquisa
type Result struct {
	Title string
	Link  string
}

// Search faz uma pesquisa nos seguintes serviços: Google, Bing (nessa ordem)
func Search(query string, nsfw bool) ([]Result, string, error) {
	var results []Result
	var provider string

	results, err := SearchGoogle(query, nsfw)
	provider = "google"

	// se não houver resultados, pesquisar no Bing
	if err != nil || len(results) < 1 {
		results, err = SearchBing(query, nsfw)
		provider = "bing"
	}

	return results, provider, err
}
