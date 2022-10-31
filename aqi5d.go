package qweather

import (
	"net/url"
)

func (q *QWeather) Aqi5d(location string) ([]*aqiDailyItem, error) {
	values := url.Values{}
	values.Add("location", location)
	respBytes, err := q.request("/v7/air/5d", values)
	if err != nil {
		return nil, err
	}
	var r aqiDailyResponse
	err = decode(respBytes, &r)
	if err != nil {
		return nil, err
	}
	return r.Daily, nil
}
