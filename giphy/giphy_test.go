package giphy

import (
	"testing"
)

func TestCallSearchApi(t *testing.T) {
	t.Log("calling the giphy api")
	gif, err := callSearchApi("testing")
	if err != nil {
		t.Error("An error occured" + err.Error())
	}

	if gif == "" {
		t.Error("No gif was returned when it should have")
	}
}

func TestGetKeyword(t *testing.T) {
	t.Log("does the proper keyword get pulled out of the string")
	testString := "gif me test"
	keyword, err := getKeyword(testString)
	if err != nil {
		t.Error("An error occured" + err.Error())
	}

	if keyword != "test" {
		t.Error("the proper keyword was not pulled out of the string")
	}
}

func TestGetGif(t *testing.T) {
	t.Log("does the exposed method pass")
	testString := "gif me string"
	gif, err := GetGif(testString)
	if err != nil {
		t.Error("An error occured" + err.Error())
	}

	if gif == "" {
		t.Error("An empy string was returned instead of a gif")
	}
}
