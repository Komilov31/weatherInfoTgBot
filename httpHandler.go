package main

import (
	"encoding/json"
	"fmt"
	"io"
	// "log"
	"net/http"
)

var cl []CityLocation
var weather Weather
var data = make([]byte, 0, 1000000)

func GetTempByCity(city string) string {
	urlForCityInfo := fmt.Sprintf("https://api.openweathermap.org/geo/1.0/direct?q=%s&limit=1&appid=2362176bfeea8c0f2a129c30a714883b", city)

	response, err := http.Get(urlForCityInfo)

	if err != nil {
		// log.Fatal("Something with requesting for CityInfo went wrong!!")
		return "Леее, че нормально город не можешь писать??"
	}

	if response.StatusCode == http.StatusOK {
		read := response.Body
		defer read.Close()
		data, err = io.ReadAll(read)

		if err != nil {
			// log.Fatal("Something went wrong with reading response")
			return "Леее, че нормально город не можешь писать??"
		}
	}

	err = json.Unmarshal(data, &cl)
	if err != nil {
		// log.Fatal("Something went wrong while unmarshaling CityInfo")
		return "Леее, че нормально город не можешь писать??"
	}

	if len(cl) == 0 {
		return "Леее, че нормально город не можешь писать??"
	}
	urlForTempInfo := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%.7f&lon=%.7f&appid=2362176bfeea8c0f2a129c30a714883b", cl[0].Lat, cl[0].Lon)
	response, err = http.Get(urlForTempInfo)

	if err != nil {
		// log.Fatal("Something with requesting for Temperature went wrong!!")
		return "Леее, че нормально город не можешь писать??"
	}

	if response.StatusCode == http.StatusOK {
		read := response.Body
		defer read.Close()
		data, err = io.ReadAll(read)

		if err != nil {
			// log.Fatal("Something went wrong with reading response")
			return "Леее, че нормально город не можешь писать??"
		}
	}

	err = json.Unmarshal(data, &weather)
	if err != nil {
		// log.Fatal("Something went wrong while unmarshaling WeatherInfo")
		return "Леее, че нормально город не можешь писать??"
	}

	result := weather.Main.Temp - 273.15
	answer := ""

	if result < 5 {
		answer = fmt.Sprintf("Температура в %s сейчас %.2f градусов. Холодно, оденься теплее", city, result)
	} else if result < 12 {
		answer = fmt.Sprintf("Температура в %s сейчас %.2f градусов. Прохладно сегодня", city, result)
	} else {
		answer = fmt.Sprintf("Температура в %s сейчас %.2f градусов. Погода как раз, чтобы прогуляться", city, result)
	}

	return answer
}
