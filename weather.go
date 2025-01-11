package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"

	"github.com/Komilov31/weatherInfoBot/model"
)

func GetWeatherStringByCity(city string) string {
	coordinates, err := GetCityCoordinates(city)
	if err != nil {
		return err.Error()
	}

	weather, err := GetWeatherByCoordinates(coordinates.Lon, coordinates.Lat)
	if err != nil {
		return err.Error()
	}

	tempCelcius := int(math.Round(weather.Main.Temp - 273.15))
	return WeatherToString(city, tempCelcius)
}

func GetCityCoordinates(city string) (model.CityLocation, error) {
	urlForCityInfo := fmt.Sprintf("https://api.openweathermap.org/geo/1.0/direct?q=%s&limit=1&appid=%s", city, os.Getenv("weather_api_token"))
	response, err := http.Get(urlForCityInfo)
	if err != nil {
		return model.CityLocation{}, errors.New("Леее, че нормально город не можешь писать??")
	}

	if response.StatusCode != http.StatusOK {
		return model.CityLocation{}, errors.New("Леее, че нормально город не можешь писать??")
	}

	var cl []model.CityLocation
	err = json.NewDecoder(response.Body).Decode(&cl)
	if err != nil {
		return model.CityLocation{}, errors.New("Леее, че нормально город не можешь писать??")
	}

	if len(cl) == 0 {
		return model.CityLocation{}, errors.New("Леее, че нормально город не можешь писать??")
	}

	return cl[0], nil
}

func GetWeatherByCoordinates(lon, lat float64) (model.Weather, error) {
	urlForTempInfo := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%.7f&lon=%.7f&appid=%s", lat, lon, os.Getenv("weather_api_token"))
	response, err := http.Get(urlForTempInfo)
	if err != nil {
		log.Println("error http.Get", err)
		return model.Weather{}, errors.New("не смог узнать погоду, прости")
	}

	if response.StatusCode != http.StatusOK {
		log.Println("error status code not 200")
		return model.Weather{}, errors.New("не смог узнать погоду, прости")
	}

	read := response.Body
	defer read.Close()
	data, err := io.ReadAll(read)
	if err != nil {
		log.Println("error io.ReadAll", err)
		return model.Weather{}, errors.New("не смог узнать погоду, прости")
	}

	var weather model.Weather
	err = json.Unmarshal(data, &weather)
	if err != nil {
		log.Println("error json.Unmarshal", err)
		return model.Weather{}, errors.New("не смог узнать погоду, прости")
	}

	return weather, nil
}

func WeatherToString(city string, tempCelcius int) string {
	answer := ""
	if tempCelcius < 5 {
		answer = fmt.Sprintf("Температура в %s сейчас %d градусов. Холодно, оденься теплее", city, tempCelcius)
	} else if tempCelcius < 12 {
		answer = fmt.Sprintf("Температура в %s сейчас %d градусов. Прохладно сегодня", city, tempCelcius)
	} else {
		answer = fmt.Sprintf("Температура в %s сейчас %d градусов. Погода как раз, чтобы прогуляться", city, tempCelcius)
	}

	return answer
}
