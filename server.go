package main

import (							//Import the packages that are relavant
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/iancoleman/orderedmap"
	"github.com/labstack/echo/v4/middleware"
)

type WeatherData struct{			//Define struct fields for the relavant data (City name, temperature, pressure, humidity)
	Name string `json:"name"`
	Main struct{
		Temp float64 `json:temp`
		Pressure float64 `json:pressure`
		Humidity float64 `json:humidity`
	}`json:"main"`
	Weather []struct {
		Description string `json:description`
		Main string `json:main`
	}`json:"weather"`

}

func NewOrderedMapFromMap(m map[string]interface{}) *orderedmap.OrderedMap {
    om := orderedmap.New()
    for key, value := range m {
        om.Set(key, value)
    }
    return om
}

func getWeatherDetails(c echo.Context) error {
	city := c.Param("city")					//Retrieve the value of the "city" query parameter from the request
	apikey := "c0f10cb90100be1f117f65319f917b0e"	//Api-key
	apiurl := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apikey)     //Api-Url

	response, err := http.Get(apiurl)				//Make HTTP GET request to apiurl
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(response.Body)		//Read the response body of the HTTP request
	if err != nil {
		return err
	}

	var weatherData WeatherData						//Unmarshal the json data (convert the json data) 
	if err := json.Unmarshal(body, &weatherData); err != nil { 
		return err
	}

	responseData := NewOrderedMapFromMap(map[string]interface{}{			//map created to structure the weather data obtained from the Api
		"Location":	weatherData.Name,
		"Temperature":		weatherData.Main.Temp,
		"Pressure":			weatherData.Main.Pressure,
		"Humidity":			weatherData.Main.Humidity,
		"WeatherType":		weatherData.Weather[0].Main,
		"WeatheerDescription":		weatherData.Weather[0].Description,
	})

	return c.JSON(http.StatusOK, responseData)		//return the response data
}

func main() {
	e := echo.New()

	// Enable CORS
	e.Use(middleware.CORS())

	e.GET("/weather/:city", getWeatherDetails)

	e.Logger.Fatal(e.Start(":8080"))
}

