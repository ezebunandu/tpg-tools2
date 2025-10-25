package main

import (
	"fmt"
	"os"

	"github.com/ezebunandu/weather"
)

const BaseURL = "https://api.openweathermap.org"

const Usage = `Usage: weather LOCATION

Example: weather Calgary,CA`

func main(){
    key := os.Getenv("OPENWEATHERMAP_API_KEY")
    if key == "" {
        fmt.Fprintln(os.Stderr, "Please set the environment variable OPENWEATHERMAP_API_KEY.")
        os.Exit(1)
    }
    if len(os.Args) < 2 {
        fmt.Println(Usage)
        os.Exit(0)
    }
    location := os.Args[1]
    c := weather.NewClient(key)
    conditions, err := c.GetWeather(location)
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
    fmt.Println(conditions)
}