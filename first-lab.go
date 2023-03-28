package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/tidwall/gjson"
)

/**
Author: Wang Fei (Kaelan)
institute: VIK
Date: 2023-3-14
Description: This program is used to get the data from the penweather.
**/
func main() {
	//get the weather data from openweathermap.org
	city := "Budapest"
	apiKey := "3d58c5af590dc4dfaaf561febf931a3b"
	//create the query url
	queryUrl := "https://api.openweathermap.org/data/2.5/weather?q=" + city + "&appid=" + apiKey
	//make the request
	resp, err := http.Get(queryUrl)
	//check if the request failed
	if err != nil {
		fmt.Println("Request failed:", err)
		return
	}
	//close the body
	defer resp.Body.Close()
	//read the body
	body, err := ioutil.ReadAll(resp.Body)
	//check if the read failed
	if err != nil {
		fmt.Println("Read body failed:", err)
		return
	}
	//convert the body to string
	jsonStr := string(body)
	//parse the json string and get the temperature from jsonStr
	kelvinTemperature := gjson.Get(jsonStr, "main.temp")
	//convert the temperature from kelvin to celsius
	celsiusTemperature := kelvinTemperature.Float() - 273.15
	//convert the temperature from kelvin to fahrenheit
	fahrenheitTemperature := celsiusTemperature*9/5 + 32
	fmt.Println("Temperature in Budapest:", celsiusTemperature, "Celsius")
	fmt.Println("Temperature in Budapest:", kelvinTemperature, "kelvin")
	fmt.Println("Temperature in Budapest:", fahrenheitTemperature, "Fahrenheit")
}
