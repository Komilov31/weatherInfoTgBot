package main

type CityLocation struct {
	Name string  `json:"name"`
	Lat  float64 `json:"lat"`
	Lon  float64 `json:"lon"`
}

type MainInfo struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
}

type Weather struct {
	Main MainInfo `json:"main"`
	Name string   `json:"name"`
}

// func (wr *Weather) ParseWeather(data []byte) {
// 	err := json.Unmarshal(data, wr)

// 	if err != nil {
// 		log.Fatal("SOMETHING WENT WRONG WITH PARSING Weather json!!!")
// 	}
// }

// func (cl *CityLocation) ParseCityInfo(data []byte) {
// 	err := json.Unmarshal(data, cl)

// 	if err != nil {
// 		log.Fatal("SOMETHING WENT WRONG WITH PARSING CITY LAT AND LON INFO!!")
// 	}
// }
