package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type WikiResponse struct {
	Query struct {
		Pages map[string]struct {
			Extract   string `json:"extract"`
			Title     string `json:"title"`
			PageProps struct {
				ShortDescription string `json:"wikibase-shortdesc"`
			} `json:"pageprops"`
		} `json:"pages"`
	} `json:"query"`
}

type SummaryData struct {
	title   string
	extract string
}

func getWikipediaSummary(title string, flags Flags) (SummaryData, error) {
	apiUrl := "https://en.wikipedia.org/w/api.php"

	params := url.Values{}
	params.Add("action", "query")
	params.Add("prop", "extracts|pageprops")
	params.Add("explaintext", "true")
	params.Add("titles", title)
	params.Add("format", "json")

	if flags.intro {
		params.Add("exintro", "true")
	}

	resp, err := http.Get(fmt.Sprintf("%s?%s", apiUrl, params.Encode()))
	if err != nil {
		return SummaryData{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return SummaryData{}, err
	}

	var wikiResponse WikiResponse
	err = json.Unmarshal(dat, &wikiResponse)
	if err != nil {
		return SummaryData{}, err
	}

	for _, page := range wikiResponse.Query.Pages {
		return SummaryData{extract: page.Extract, title: page.Title}, nil
	}

	return SummaryData{title: title, extract: "No summary available for this query"}, nil
}
