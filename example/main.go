package main

import (
	"fmt"
	"github.com/wbylovesun/qweather"
)

var location = "101280603"
var lat = 22.54101
var lng = 114.05095
var q, _ = qweather.NewQWeather(qweather.FreeMode, "b2ce44e38fe74e1f9447a91d90ef7ced")

func main() {
	realtime(location)
	p24h(location)
	rain(lat, lng)
	aqi(location)
	aqi5d(location)
	indice1d(location)
	indice3d(location)
	p3d(location)
	p7d(location)
	warning(location)
}

func realtime(location string) {
	value, err := q.Realtime(location)
	if err != nil {
		fmt.Println("qweather, realtime: ", err)
	} else {
		fmt.Println(value)
	}
}

func p24h(location string) {
	value, err := q.P24h(location)
	if err != nil {
		fmt.Println("qweather, p24h: ", err)
	} else {
		for _, v := range value {
			fmt.Println(v)
		}
	}
}

func rain(lat, lng float64) {
	value, err := q.Rain(lat, lng)
	if err != nil {
		fmt.Println("qweather, rain: ", err)
	} else {
		fmt.Println(value.Summary)
		for _, v := range value.Minutely {
			fmt.Println(v)
		}
	}
}

func aqi(location string) {
	value, err := q.Aqi(location)
	if err != nil {
		fmt.Println("qweather, aqi: ", err)
	} else {
		fmt.Println(value)
	}
}

func aqi5d(location string) {
	value, err := q.Aqi5d(location)
	if err != nil {
		fmt.Println("qweather, aqi5d: ", err)
	} else {
		for _, v := range value {
			fmt.Println(v)
		}
	}
}

func indice1d(location string) {
	value, err := q.Indice1d(location, nil)
	if err != nil {
		fmt.Println("qweather, indice1d: ", err)
	} else {
		for _, v := range value {
			fmt.Println(v)
		}
	}
}

func indice3d(location string) {
	value, err := q.Indice3d(location, nil)
	if err != nil {
		fmt.Println("qweather, indice3d: ", err)
	} else {
		for _, v := range value {
			fmt.Println(v)
		}
	}
}

func p3d(location string) {
	value, err := q.P3d(location)
	if err != nil {
		fmt.Println("qweather, p3d: ", err)
	} else {
		for _, v := range value {
			fmt.Println(v)
		}
	}
}

func p7d(location string) {
	value, err := q.P7d(location)
	if err != nil {
		fmt.Println("qweather, p7d: ", err)
	} else {
		for _, v := range value {
			fmt.Println(v)
		}
	}
}

func warning(location string) {
	value, err := q.Warning(location)
	if err != nil {
		fmt.Println("qweather, warning: ", err)
	} else {
		for _, v := range value {
			fmt.Println(v)
		}
	}
}
