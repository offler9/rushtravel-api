package main

import (
    "os"
    "fmt"
    "net/http"
    "log"
    "encoding/json"
    "io/ioutil"
)

const helloMessage = "Hello to the weather program. Please enter the name of the city and the weather will show."
const googleApiUri = "https://maps.googleapis.com/maps/api/geocode/json?key=MYKEY&address="



type googleApiResponse struct {
    Results Results `json:"results"`
}

type Results []Geometry

type Geometry struct {
    Geometry Location `json:"geometry"`
}

type Location struct {
    Location Coordinates `json:"location"`
}

type Coordinates struct {
    Latitude float64 `json:"lat"`
    Longitude float64 `json:"lng"`
}


func main() {
    fmt.Println(helloMessage)

    args := os.Args
    getCityCoordinates(args[0])
}

func getCityCoordinates(city string) {
    fmt.Println("Fetching langitude and longitude of the city ...")
    resp, err := http.Get(googleApiUri + city)

    if err != nil {
        log.Fatal("Fetching google api uri data error: ", err)
    }

    bytes, err := ioutil.ReadAll(resp.Body)
    defer resp.Body.Close()
    if err != nil {
        log.Fatal("Reading google api data error: ", err)
    }

    var data googleApiResponse
    json.Unmarshal(bytes, &data)
    fmt.Println(data.Results[0].Geometry.Location.Latitude)

    fmt.Println("Fetching langitude and longitude ended successful ...")
}