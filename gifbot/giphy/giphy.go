// Model of the json response from the giphy /search endpoint
package giphy

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strings"

	"github.com/user/gifbot/app/models"
)

func getKeyword(content string) (string, error) {
	var err error
	var keyword string
	content = strings.ToLower(content)
	splitArr := strings.Split(content, "gif me ")
	if len(splitArr) > 1 {
		keyword = splitArr[1]
	} else {
		err = errors.New("No keyword found in message")
	}
	return keyword, err
}

func GetGif(discordMessage string) (string, error) {
	var err error
	keyword, err := getKeyword(discordMessage)
	if err != nil {
		return "", err
	}
	gif, err := callSearchApi(keyword)
	return gif, err
}

func callSearchApi(keyword string) (string, error) {
	var err error
	noGifsFound := "No gifs found when searching for " + keyword
	var giphyJson models.GiphyJson
	keywordQuery := url.QueryEscape(keyword)
	// grabes the maximun number of images gify allows
	resp, err := http.Get("http://api.giphy.com/v1/gifs/search?q=" + keywordQuery + "&api_key=dc6zaTOxFJmzC&limit=100")
	if err != nil {
		err = errors.New("An error occured trying to contact giphy")
		return "", err
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bodyBytes, &giphyJson)
	if err != nil {
		err = errors.New("An error occured trying to parse giphy's json")
		return "", err
	}

	if len(giphyJson.Data) == 0 {
		return noGifsFound, nil
	}
	// pick a random images from the images array
	gif := giphyJson.Data[rand.Intn(len(giphyJson.Data))].Images.Fixed_height.Url
	if gif == "" {
		return noGifsFound, nil
	}
	return gif, err
}
