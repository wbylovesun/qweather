package qweather

import (
	"net/url"
)

func (q *QWeather) Realtime(location string) (*realtimeNow, error) {
	values := url.Values{}
	values.Add("location", location)
	respBytes, err := q.request("/v7/weather/now", values)
	if err != nil {
		return nil, err
	}
	var r realtimeResponse
	err = decode(respBytes, &r)
	if err != nil {
		return nil, err
	}
	return &r.Now, nil
}
