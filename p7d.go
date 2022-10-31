package qweather

import (
	"net/url"
)

func (q *QWeather) P7d(location string) ([]*pDailyItem, error) {
	values := url.Values{}
	values.Add("location", location)
	respBytes, err := q.request("/v7/weather/3d", values)
	if err != nil {
		return nil, err
	}
	var r pDailyResponse
	err = decode(respBytes, &r)
	if err != nil {
		return nil, err
	}
	return r.Daily, nil
}
