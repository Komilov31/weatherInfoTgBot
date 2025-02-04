package model

var DefaultMessage = `
Вот список доступных команд: 
/weather     - для получения температуры у тебя в местности
/setlocation - для установки твоего города

Также тебе будет отправлен стикер, при благодарности боту!
`

type UserStore interface {
	GetUserByName(string) (*User, error)
	SetLocation(*User) error
}

type User struct {
	Id       int
	UserName string
	City     string
	Lat      float64
	Lon      float64
}

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

var GratitudeMessages = map[string]struct{}{
	"Ты молодец": struct{}{},
	"Хорош":      struct{}{},
	"Thank you":  struct{}{},
	"Спасибо":    struct{}{},
	"Красавчик":  struct{}{},
	"Молодец":    struct{}{},
}
