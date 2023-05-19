package main

import (							//Import the packages that are relavant
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/labstack/echo/v4"
)

type WeatherData struct{			//Define struct fields for the relavant data (City name, temperature, pressure, humidity)
	Name string `json:"name"`
	Main struct{
		Temp float64 `json:temp`
		Pressure float64 `json:pressure`
		Humidity float64 `json:humidity`
	}`json:"main"`

}

func getWeatherDetails(c echo.Context) error {
	city := c.QueryParam("city")																		//Retrieve the value of the "city" query parameter from the request
	apikey := "c0f10cb90100be1f117f65319f917b0e"	//Api-key
	apiurl := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apikey)						//Api-Url

	response, err := http.Get(apiurl)						//Make HTTP GET request to apiurl
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(response.Body)			//Read the response body of the HTTP request
	if err != nil {
		return err
	}

	var weatherData WeatherData																	//Unmarshal thejson data (convert the json data) 
	if err := json.Unmarshal(body, &weatherData); err != nil { 
		return err
	}

	responseData := map[string]interface{}{				//map created to structure the weather data obtained from the Api
		"Location (city)":	weatherData.Name,
		"Temperature":		weatherData.Main.Temp,
		"Pressure":			weatherData.Main.Pressure,
		"Humidity":			weatherData.Main.Humidity,
	}

	return c.JSON(http.StatusOK, responseData)		//return the response data
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello Guys, Welcome to the WEATHER APP")
	})

	e.GET("/weather/", getWeatherDetails)

	e.Logger.Fatal(e.Start(":1323"))
}
