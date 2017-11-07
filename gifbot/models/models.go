package models

// Model of the json response from the giphy /search endpoint
type GiphyResp struct {
	Type string `json:"type"`
	Data []struct {
		Images struct {
			Original struct {
				Url string `json:"url"`
			} `json:"original"`
		} `json:"images"`
	} `json:"data"`
}
