package weather_test

import (
	"os"
	"testing"

	"github.com/ezebunandu/weather"

    "github.com/google/go-cmp/cmp"
)
func TestParseResponse__CorrectlyParsesJSONData(t *testing.T){
    t.Parallel()
    data, err := os.ReadFile("testdata/weather.json")
    if err != nil {
        t.Fatal(err)
    }
    want := weather.Conditions{
        Summary: "Clouds",
    }
    got, err := weather.ParseResponse(data)
    if err != nil {
        t.Fatal(err)
    }
    if !cmp.Equal(want, got){
        t.Error(cmp.Diff(want, got))
    }
}

func TestParseResponse__ReturnsErrorGivenEmptyData(t *testing.T){
    t.Parallel()
    _, err := weather.ParseResponse([]byte{})
    if err == nil {
        t.Fatal("want error parsing empty response, got nil")
    }
}

func TestParseResponse__ReturnsErrorGivenInvalidJSON(t *testing.T){
    t.Parallel()
    data, err := os.ReadFile("testdata/weather_invalid.json")
    if err != nil {
        t.Fatal(err)
    }
    _, err = weather.ParseResponse(data)
    if err == nil {
        t.Fatal("want error parsing invalid response, got nil")
    }
}

func TestFormatURL__ReturnsCorrectURLForGivenInput(t *testing.T){
    t.Parallel()
    baseURL := weather.BaseURL
    location := "Calgary,CA"
    key := "dummyAPIKey"
    want := "https://api.openweathermap.org/data/2.5/weather?q=Calgary,CA&appid=dummyAPIKey"
    got := weather.FormatURL(baseURL, location, key)
    if !cmp.Equal(want, got){
        t.Error(cmp.Diff(want, got))
    }
}