package models

// Model of the json response from the giphy /search endpoint
type GiphyJson struct {
	Type string `json:"type"`
	Data []struct {
		Images struct {
			Fixed_height struct {
				Url string `json:"url"`
			} `json:"fixed_height"`
		} `json:"images"`
	} `json:"data"`
}
