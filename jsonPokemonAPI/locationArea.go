package jsonPokemonAPI

import (
	"encoding/json"
	"io"
	"net/http"
)

func LocationAreaList(pageURL string) (MapData, error){
	//show 20 maps
	url := "https://pokeapi.co/api/v2/location-area/"
	if pageURL != ""{
		url = pageURL
	}

	res, err := http.Get(url)
	if err != nil {
		return MapData{}, err
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return MapData{}, err
	}
	if err != nil {
		return MapData{}, err
	}
	mapData := MapData{}
	err = json.Unmarshal(body, &mapData)
	if err != nil{
		return MapData{}, err
	}
	return mapData, nil
}