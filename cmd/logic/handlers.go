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
	weather := GetWeatherStringByCity(user.City)

	return weather
}

func (h *Handler) HandleSetLocationCommand(userName, city string) string {
	err := h.store.SetLocation(userName, city)
	if err != nil {
		log.Fatal("Something went wrong while inserting to DB!")
		return "Ваш город не был сохранен. Попробуйте еще раз!"
	}
	return "Ваш город был успешно сохранен"
}
