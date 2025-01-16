package model

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

var GratitudeMessages map[string]struct{} = map[string]struct{}{
	"Ты молодец": struct{}{},
	"Хорош":      struct{}{},
	"Thank you":  struct{}{},
	"Спасибо":    struct{}{},
	"Красавчик":  struct{}{},
}
