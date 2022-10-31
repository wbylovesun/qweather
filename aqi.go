package qweather

import (
	"net/url"
)

func (q *QWeather) Aqi(location string) (*aqiItem, error) {
	values := url.Values{}
	values.Add("location", location)
	respBytes, err := q.request("/v7/air/now", values)
	if err != nil {
		return nil, err
	}
	var r aqiResponse
	err = decode(respBytes, &r)
	if err != nil {
		return nil, err
	}
	return r.Now, nil
}
