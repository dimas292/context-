package route

import (
	"defer/controller"
	"net/http"
)

func SetUp(){
	handle := controller.NewWeather()
	http.HandleFunc("/weather", handle.GetCurrentWeather)
}