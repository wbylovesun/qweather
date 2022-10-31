package qweather

import (
	"net/url"
)

func (q *QWeather) P24h(location string) ([]*pHourlyItem, error) {
	values := url.Values{}
	values.Add("location", location)
	respBytes, err := q.request("/v7/weather/24h", values)
	if err != nil {
		return nil, err
	}
	var r pHourlyResponse
	err = decode(respBytes, &r)
	if err != nil {
		return nil, err
	}
	return r.Hourly, nil
}
