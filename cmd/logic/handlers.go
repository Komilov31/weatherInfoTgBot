package logic

import (
	"log"

	"github.com/Komilov31/weatherInfoBot/model"
)

type Handler struct {
	store model.UserStore
}

func NewHandler(store model.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) HandleWeatherCommand(userName string) string {

	user, err := h.store.GetUserByName(userName)
	if err != nil {
		return "Ты пока не указал свой город!" + model.DefaultMessage
	}
	weather := getWeatherByCoordinates(user)

	return weather
}

func (h *Handler) HandleSetLocationCommand(userName, city string) string {

	cityLocation, err := GetCityCoordinates(city)
	if err != nil {
		log.Fatal("Something went wrong while taking coordinates")
		return "Леее, че нормально не можешь город писать? Попробуй еще раз!"
	}

	user := model.User{
		UserName: userName,
		City:     city,
		Lat:      cityLocation.Lat,
		Lon:      cityLocation.Lon,
	}

	err = h.store.SetLocation(&user)
	if err != nil {
		log.Fatal("Something went wrong while inserting to DB!")
		return "Ваш город не был сохранен. Попробуйте еще раз!"
	}
	return "Ваш город был успешно сохранен"
}

func getWeatherByCoordinates(user *model.User) string {
	weather, err := GetWeatherByCoordinates(user.Lon, user.Lat)
	if err != nil {
		return "Что то пошло не так. Не смог узнать погоду, прости!"
	}

	return WeatherToString(user.City, int(weather.Main.Temp-273.15))
}
